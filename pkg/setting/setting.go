package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	// 配置文件名称，没有扩展后缀等
	vp.SetConfigName("config")
	// 配置文件路径
	vp.AddConfigPath("configs/")
	// 配置文件类型
	vp.SetConfigType("yaml")
	// 读取配置文件
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
