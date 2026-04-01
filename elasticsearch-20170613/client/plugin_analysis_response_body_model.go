// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPluginAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PluginAnalysisResponseBody
	GetRequestId() *string
	SetResult(v []*PluginAnalysisResponseBodyResult) *PluginAnalysisResponseBody
	GetResult() []*PluginAnalysisResponseBodyResult
}

type PluginAnalysisResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*PluginAnalysisResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s PluginAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PluginAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *PluginAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PluginAnalysisResponseBody) GetResult() []*PluginAnalysisResponseBodyResult {
	return s.Result
}

func (s *PluginAnalysisResponseBody) SetRequestId(v string) *PluginAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *PluginAnalysisResponseBody) SetResult(v []*PluginAnalysisResponseBodyResult) *PluginAnalysisResponseBody {
	s.Result = v
	return s
}

func (s *PluginAnalysisResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PluginAnalysisResponseBodyResult struct {
	// example:
	//
	// plugin description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 8.17.0
	ElasticsearchVersion *string `json:"elasticsearchVersion,omitempty" xml:"elasticsearchVersion,omitempty"`
	// example:
	//
	// plugin_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// null
	SecurityPolicy *string `json:"securityPolicy,omitempty" xml:"securityPolicy,omitempty"`
	// example:
	//
	// 8.17.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s PluginAnalysisResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s PluginAnalysisResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PluginAnalysisResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *PluginAnalysisResponseBodyResult) GetElasticsearchVersion() *string {
	return s.ElasticsearchVersion
}

func (s *PluginAnalysisResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *PluginAnalysisResponseBodyResult) GetSecurityPolicy() *string {
	return s.SecurityPolicy
}

func (s *PluginAnalysisResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *PluginAnalysisResponseBodyResult) SetDescription(v string) *PluginAnalysisResponseBodyResult {
	s.Description = &v
	return s
}

func (s *PluginAnalysisResponseBodyResult) SetElasticsearchVersion(v string) *PluginAnalysisResponseBodyResult {
	s.ElasticsearchVersion = &v
	return s
}

func (s *PluginAnalysisResponseBodyResult) SetName(v string) *PluginAnalysisResponseBodyResult {
	s.Name = &v
	return s
}

func (s *PluginAnalysisResponseBodyResult) SetSecurityPolicy(v string) *PluginAnalysisResponseBodyResult {
	s.SecurityPolicy = &v
	return s
}

func (s *PluginAnalysisResponseBodyResult) SetVersion(v string) *PluginAnalysisResponseBodyResult {
	s.Version = &v
	return s
}

func (s *PluginAnalysisResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
