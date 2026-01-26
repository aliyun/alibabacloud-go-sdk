// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTraceAppByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *SearchTraceAppByNameRequest
	GetRegionId() *string
	SetTags(v []*SearchTraceAppByNameRequestTags) *SearchTraceAppByNameRequest
	GetTags() []*SearchTraceAppByNameRequestTags
	SetTraceAppName(v string) *SearchTraceAppByNameRequest
	GetTraceAppName() *string
}

type SearchTraceAppByNameRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of tags.
	Tags []*SearchTraceAppByNameRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The name of the application.
	//
	// > If you do not specify this parameter, all application monitoring tasks in the specified region are queried.
	//
	// example:
	//
	// test-app
	TraceAppName *string `json:"TraceAppName,omitempty" xml:"TraceAppName,omitempty"`
}

func (s SearchTraceAppByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByNameRequest) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchTraceAppByNameRequest) GetTags() []*SearchTraceAppByNameRequestTags {
	return s.Tags
}

func (s *SearchTraceAppByNameRequest) GetTraceAppName() *string {
	return s.TraceAppName
}

func (s *SearchTraceAppByNameRequest) SetRegionId(v string) *SearchTraceAppByNameRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByNameRequest) SetTags(v []*SearchTraceAppByNameRequestTags) *SearchTraceAppByNameRequest {
	s.Tags = v
	return s
}

func (s *SearchTraceAppByNameRequest) SetTraceAppName(v string) *SearchTraceAppByNameRequest {
	s.TraceAppName = &v
	return s
}

func (s *SearchTraceAppByNameRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchTraceAppByNameRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTraceAppByNameRequestTags) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByNameRequestTags) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameRequestTags) GetKey() *string {
	return s.Key
}

func (s *SearchTraceAppByNameRequestTags) GetValue() *string {
	return s.Value
}

func (s *SearchTraceAppByNameRequestTags) SetKey(v string) *SearchTraceAppByNameRequestTags {
	s.Key = &v
	return s
}

func (s *SearchTraceAppByNameRequestTags) SetValue(v string) *SearchTraceAppByNameRequestTags {
	s.Value = &v
	return s
}

func (s *SearchTraceAppByNameRequestTags) Validate() error {
	return dara.Validate(s)
}
