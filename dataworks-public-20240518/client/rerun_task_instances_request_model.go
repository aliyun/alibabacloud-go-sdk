// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunTaskInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *RerunTaskInstancesRequest
	GetComment() *string
	SetIds(v []*int64) *RerunTaskInstancesRequest
	GetIds() []*int64
	SetUseLatestConfig(v bool) *RerunTaskInstancesRequest
	GetUseLatestConfig() *bool
}

type RerunTaskInstancesRequest struct {
	// The remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The list of node instance IDs.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// Specifies whether to rerun instances with the latest configuration.
	//
	// example:
	//
	// false
	UseLatestConfig *bool `json:"UseLatestConfig,omitempty" xml:"UseLatestConfig,omitempty"`
}

func (s RerunTaskInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RerunTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *RerunTaskInstancesRequest) GetComment() *string {
	return s.Comment
}

func (s *RerunTaskInstancesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *RerunTaskInstancesRequest) GetUseLatestConfig() *bool {
	return s.UseLatestConfig
}

func (s *RerunTaskInstancesRequest) SetComment(v string) *RerunTaskInstancesRequest {
	s.Comment = &v
	return s
}

func (s *RerunTaskInstancesRequest) SetIds(v []*int64) *RerunTaskInstancesRequest {
	s.Ids = v
	return s
}

func (s *RerunTaskInstancesRequest) SetUseLatestConfig(v bool) *RerunTaskInstancesRequest {
	s.UseLatestConfig = &v
	return s
}

func (s *RerunTaskInstancesRequest) Validate() error {
	return dara.Validate(s)
}
