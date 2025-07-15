// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContactFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactFlowId(v string) *GetContactFlowRequest
	GetContactFlowId() *string
	SetDraftId(v string) *GetContactFlowRequest
	GetDraftId() *string
	SetInstanceId(v string) *GetContactFlowRequest
	GetInstanceId() *string
}

type GetContactFlowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 274601be-a6d5-4429-bcef-32b51d031c6e
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 566399d7-5558-447c-a72f-9be2768b6a82
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetContactFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContactFlowRequest) GoString() string {
	return s.String()
}

func (s *GetContactFlowRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *GetContactFlowRequest) GetDraftId() *string {
	return s.DraftId
}

func (s *GetContactFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetContactFlowRequest) SetContactFlowId(v string) *GetContactFlowRequest {
	s.ContactFlowId = &v
	return s
}

func (s *GetContactFlowRequest) SetDraftId(v string) *GetContactFlowRequest {
	s.DraftId = &v
	return s
}

func (s *GetContactFlowRequest) SetInstanceId(v string) *GetContactFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *GetContactFlowRequest) Validate() error {
	return dara.Validate(s)
}
