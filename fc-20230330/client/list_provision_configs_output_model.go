// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvisionConfigsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListProvisionConfigsOutput
	GetNextToken() *string
	SetProvisionConfigs(v []*ProvisionConfig) *ListProvisionConfigsOutput
	GetProvisionConfigs() []*ProvisionConfig
}

type ListProvisionConfigsOutput struct {
	// example:
	//
	// next_token
	NextToken        *string            `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ProvisionConfigs []*ProvisionConfig `json:"provisionConfigs" xml:"provisionConfigs" type:"Repeated"`
}

func (s ListProvisionConfigsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionConfigsOutput) GoString() string {
	return s.String()
}

func (s *ListProvisionConfigsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProvisionConfigsOutput) GetProvisionConfigs() []*ProvisionConfig {
	return s.ProvisionConfigs
}

func (s *ListProvisionConfigsOutput) SetNextToken(v string) *ListProvisionConfigsOutput {
	s.NextToken = &v
	return s
}

func (s *ListProvisionConfigsOutput) SetProvisionConfigs(v []*ProvisionConfig) *ListProvisionConfigsOutput {
	s.ProvisionConfigs = v
	return s
}

func (s *ListProvisionConfigsOutput) Validate() error {
	if s.ProvisionConfigs != nil {
		for _, item := range s.ProvisionConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
