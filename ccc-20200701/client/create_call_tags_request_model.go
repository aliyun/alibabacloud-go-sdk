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
	// This parameter is required.
	//
	// example:
	//
	// ["TagA","TagB"]
	CallTagNameList *string `json:"CallTagNameList,omitempty" xml:"CallTagNameList,omitempty"`
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
