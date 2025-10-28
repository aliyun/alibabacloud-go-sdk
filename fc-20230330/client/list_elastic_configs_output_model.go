// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListElasticConfigsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetElasticConfigs(v []*ElasticConfigStatus) *ListElasticConfigsOutput
	GetElasticConfigs() []*ElasticConfigStatus
	SetNextToken(v string) *ListElasticConfigsOutput
	GetNextToken() *string
}

type ListElasticConfigsOutput struct {
	ElasticConfigs []*ElasticConfigStatus `json:"elasticConfigs" xml:"elasticConfigs" type:"Repeated"`
	NextToken      *string                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListElasticConfigsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListElasticConfigsOutput) GoString() string {
	return s.String()
}

func (s *ListElasticConfigsOutput) GetElasticConfigs() []*ElasticConfigStatus {
	return s.ElasticConfigs
}

func (s *ListElasticConfigsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListElasticConfigsOutput) SetElasticConfigs(v []*ElasticConfigStatus) *ListElasticConfigsOutput {
	s.ElasticConfigs = v
	return s
}

func (s *ListElasticConfigsOutput) SetNextToken(v string) *ListElasticConfigsOutput {
	s.NextToken = &v
	return s
}

func (s *ListElasticConfigsOutput) Validate() error {
	if s.ElasticConfigs != nil {
		for _, item := range s.ElasticConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
