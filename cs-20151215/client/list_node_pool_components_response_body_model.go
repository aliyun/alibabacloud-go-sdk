// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePoolComponentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponents(v []*ListNodePoolComponentsResponseBodyComponents) *ListNodePoolComponentsResponseBody
	GetComponents() []*ListNodePoolComponentsResponseBodyComponents
	SetMaxResults(v int32) *ListNodePoolComponentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodePoolComponentsResponseBody
	GetNextToken() *string
}

type ListNodePoolComponentsResponseBody struct {
	// The node components.
	Components []*ListNodePoolComponentsResponseBodyComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// The maximum number of results per page. Default value: 200.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"max_results,omitempty" xml:"max_results,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// 5c0a1c0f91c14c6****
	NextToken *string `json:"next_token,omitempty" xml:"next_token,omitempty"`
}

func (s ListNodePoolComponentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentsResponseBody) GetComponents() []*ListNodePoolComponentsResponseBodyComponents {
	return s.Components
}

func (s *ListNodePoolComponentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodePoolComponentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodePoolComponentsResponseBody) SetComponents(v []*ListNodePoolComponentsResponseBodyComponents) *ListNodePoolComponentsResponseBody {
	s.Components = v
	return s
}

func (s *ListNodePoolComponentsResponseBody) SetMaxResults(v int32) *ListNodePoolComponentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNodePoolComponentsResponseBody) SetNextToken(v string) *ListNodePoolComponentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNodePoolComponentsResponseBody) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodePoolComponentsResponseBodyComponents struct {
	// The parameter schema of the node component, including custom parameters and system parameters.
	//
	// example:
	//
	// "{\\"$schema\\": \\"\\",\\"properties\\": {}"
	ConfigSchema *string `json:"config_schema,omitempty" xml:"config_schema,omitempty"`
	// The name of the node component.
	//
	// example:
	//
	// kubelet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the node component.
	//
	// example:
	//
	// 1.28.9-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListNodePoolComponentsResponseBodyComponents) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentsResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentsResponseBodyComponents) GetConfigSchema() *string {
	return s.ConfigSchema
}

func (s *ListNodePoolComponentsResponseBodyComponents) GetName() *string {
	return s.Name
}

func (s *ListNodePoolComponentsResponseBodyComponents) GetVersion() *string {
	return s.Version
}

func (s *ListNodePoolComponentsResponseBodyComponents) SetConfigSchema(v string) *ListNodePoolComponentsResponseBodyComponents {
	s.ConfigSchema = &v
	return s
}

func (s *ListNodePoolComponentsResponseBodyComponents) SetName(v string) *ListNodePoolComponentsResponseBodyComponents {
	s.Name = &v
	return s
}

func (s *ListNodePoolComponentsResponseBodyComponents) SetVersion(v string) *ListNodePoolComponentsResponseBodyComponents {
	s.Version = &v
	return s
}

func (s *ListNodePoolComponentsResponseBodyComponents) Validate() error {
	return dara.Validate(s)
}
