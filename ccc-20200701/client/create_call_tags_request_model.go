// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallTagNameList(v string) *CreateCallTagsRequest
	GetCallTagNameList() *string
	SetInstanceId(v string) *CreateCallTagsRequest
	GetInstanceId() *string
}

type CreateCallTagsRequest struct {
	// A JSON-formatted string representing an array of call tag names. Each array element is a call tag name to be created. The length of each call tag name must be between 1 and 10 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["TagA","TagB"]
	CallTagNameList *string `json:"CallTagNameList,omitempty" xml:"CallTagNameList,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateCallTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCallTagsRequest) GoString() string {
	return s.String()
}

func (s *CreateCallTagsRequest) GetCallTagNameList() *string {
	return s.CallTagNameList
}

func (s *CreateCallTagsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCallTagsRequest) SetCallTagNameList(v string) *CreateCallTagsRequest {
	s.CallTagNameList = &v
	return s
}

func (s *CreateCallTagsRequest) SetInstanceId(v string) *CreateCallTagsRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCallTagsRequest) Validate() error {
	return dara.Validate(s)
}
