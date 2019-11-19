package option

// RunCmdConfig is config for run command
type RunCmdConfig struct {
	File  string
}

// NewRunCmdConfigFromViper generate config for run command from viper
func NewRunCmdConfigFromViper() (*RunCmdConfig, error) {
	rawConfig, err := newCmdRawConfig()
	return newRunCmdConfigFromRawConfig(rawConfig), err
}

func newRunCmdConfigFromRawConfig(rawConfig *CmdRawConfig) *RunCmdConfig {
	return &RunCmdConfig{
		File:  rawConfig.File,
	}
}

func (c *RunCmdConfig) validate() error {
	return nil
}
