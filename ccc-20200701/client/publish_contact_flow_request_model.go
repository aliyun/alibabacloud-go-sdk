// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishContactFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactFlowId(v string) *PublishContactFlowRequest
	GetContactFlowId() *string
	SetDraftId(v string) *PublishContactFlowRequest
	GetDraftId() *string
	SetInstanceId(v string) *PublishContactFlowRequest
	GetInstanceId() *string
}

type PublishContactFlowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b0a063bf-f138-42a4-ad9f-9babe3ec1a9e
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0e0b8e78-af3e-4360-a5c5-f9bb5c2b8af2
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s PublishContactFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishContactFlowRequest) GoString() string {
	return s.String()
}

func (s *PublishContactFlowRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *PublishContactFlowRequest) GetDraftId() *string {
	return s.DraftId
}

func (s *PublishContactFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PublishContactFlowRequest) SetContactFlowId(v string) *PublishContactFlowRequest {
	s.ContactFlowId = &v
	return s
}

func (s *PublishContactFlowRequest) SetDraftId(v string) *PublishContactFlowRequest {
	s.DraftId = &v
	return s
}

func (s *PublishContactFlowRequest) SetInstanceId(v string) *PublishContactFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *PublishContactFlowRequest) Validate() error {
	return dara.Validate(s)
}
