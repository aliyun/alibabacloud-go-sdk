// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConcurrencyConfigsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ConcurrencyConfig) *ListConcurrencyConfigsOutput
	GetConfigs() []*ConcurrencyConfig
	SetNextToken(v string) *ListConcurrencyConfigsOutput
	GetNextToken() *string
}

type ListConcurrencyConfigsOutput struct {
	Configs []*ConcurrencyConfig `json:"configs" xml:"configs" type:"Repeated"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListConcurrencyConfigsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListConcurrencyConfigsOutput) GoString() string {
	return s.String()
}

func (s *ListConcurrencyConfigsOutput) GetConfigs() []*ConcurrencyConfig {
	return s.Configs
}

func (s *ListConcurrencyConfigsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConcurrencyConfigsOutput) SetConfigs(v []*ConcurrencyConfig) *ListConcurrencyConfigsOutput {
	s.Configs = v
	return s
}

func (s *ListConcurrencyConfigsOutput) SetNextToken(v string) *ListConcurrencyConfigsOutput {
	s.NextToken = &v
	return s
}

func (s *ListConcurrencyConfigsOutput) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
