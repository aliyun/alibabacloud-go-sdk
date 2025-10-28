// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncInvokeConfigOutput interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*AsyncConfig) *ListAsyncInvokeConfigOutput
	GetConfigs() []*AsyncConfig
	SetNextToken(v string) *ListAsyncInvokeConfigOutput
	GetNextToken() *string
}

type ListAsyncInvokeConfigOutput struct {
	Configs []*AsyncConfig `json:"configs" xml:"configs" type:"Repeated"`
	// example:
	//
	// 8bj81uI8n****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAsyncInvokeConfigOutput) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncInvokeConfigOutput) GoString() string {
	return s.String()
}

func (s *ListAsyncInvokeConfigOutput) GetConfigs() []*AsyncConfig {
	return s.Configs
}

func (s *ListAsyncInvokeConfigOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAsyncInvokeConfigOutput) SetConfigs(v []*AsyncConfig) *ListAsyncInvokeConfigOutput {
	s.Configs = v
	return s
}

func (s *ListAsyncInvokeConfigOutput) SetNextToken(v string) *ListAsyncInvokeConfigOutput {
	s.NextToken = &v
	return s
}

func (s *ListAsyncInvokeConfigOutput) Validate() error {
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
