// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckResourceCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *PrecheckResourceCountRequest
	GetResourceType() *string
	SetTagResourceMatchers(v []*PrecheckResourceCountRequestTagResourceMatchers) *PrecheckResourceCountRequest
	GetTagResourceMatchers() []*PrecheckResourceCountRequestTagResourceMatchers
}

type PrecheckResourceCountRequest struct {
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	TagResourceMatchers []*PrecheckResourceCountRequestTagResourceMatchers `json:"TagResourceMatchers,omitempty" xml:"TagResourceMatchers,omitempty" type:"Repeated"`
}

func (s PrecheckResourceCountRequest) String() string {
	return dara.Prettify(s)
}

func (s PrecheckResourceCountRequest) GoString() string {
	return s.String()
}

func (s *PrecheckResourceCountRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *PrecheckResourceCountRequest) GetTagResourceMatchers() []*PrecheckResourceCountRequestTagResourceMatchers {
	return s.TagResourceMatchers
}

func (s *PrecheckResourceCountRequest) SetResourceType(v string) *PrecheckResourceCountRequest {
	s.ResourceType = &v
	return s
}

func (s *PrecheckResourceCountRequest) SetTagResourceMatchers(v []*PrecheckResourceCountRequestTagResourceMatchers) *PrecheckResourceCountRequest {
	s.TagResourceMatchers = v
	return s
}

func (s *PrecheckResourceCountRequest) Validate() error {
	if s.TagResourceMatchers != nil {
		for _, item := range s.TagResourceMatchers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PrecheckResourceCountRequestTagResourceMatchers struct {
	// This parameter is required.
	//
	// example:
	//
	// CreatedBy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EQUAL
	Operator *string   `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Values   []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s PrecheckResourceCountRequestTagResourceMatchers) String() string {
	return dara.Prettify(s)
}

func (s PrecheckResourceCountRequestTagResourceMatchers) GoString() string {
	return s.String()
}

func (s *PrecheckResourceCountRequestTagResourceMatchers) GetKey() *string {
	return s.Key
}

func (s *PrecheckResourceCountRequestTagResourceMatchers) GetOperator() *string {
	return s.Operator
}

func (s *PrecheckResourceCountRequestTagResourceMatchers) GetValues() []*string {
	return s.Values
}

func (s *PrecheckResourceCountRequestTagResourceMatchers) SetKey(v string) *PrecheckResourceCountRequestTagResourceMatchers {
	s.Key = &v
	return s
}

func (s *PrecheckResourceCountRequestTagResourceMatchers) SetOperator(v string) *PrecheckResourceCountRequestTagResourceMatchers {
	s.Operator = &v
	return s
}

func (s *PrecheckResourceCountRequestTagResourceMatchers) SetValues(v []*string) *PrecheckResourceCountRequestTagResourceMatchers {
	s.Values = v
	return s
}

func (s *PrecheckResourceCountRequestTagResourceMatchers) Validate() error {
	return dara.Validate(s)
}
