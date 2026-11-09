// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTags(v string) *ListInstancesRequest
	GetTags() *string
}

type ListInstancesRequest struct {
	// A list of resource tags in JSON format.
	//
	// example:
	//
	// [{"key":"test","value":"test"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetTags() *string {
	return s.Tags
}

func (s *ListInstancesRequest) SetTags(v string) *ListInstancesRequest {
	s.Tags = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
