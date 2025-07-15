// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscardEditingContactFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactFlowId(v string) *DiscardEditingContactFlowRequest
	GetContactFlowId() *string
	SetDraftId(v string) *DiscardEditingContactFlowRequest
	GetDraftId() *string
	SetInstanceId(v string) *DiscardEditingContactFlowRequest
	GetInstanceId() *string
}

type DiscardEditingContactFlowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3ff4e021-fd63-4572-ad8c-10ed69972965
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0aa493d6-58eb-4290-9ba2-e1c2c615b46b
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DiscardEditingContactFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DiscardEditingContactFlowRequest) GoString() string {
	return s.String()
}

func (s *DiscardEditingContactFlowRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *DiscardEditingContactFlowRequest) GetDraftId() *string {
	return s.DraftId
}

func (s *DiscardEditingContactFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DiscardEditingContactFlowRequest) SetContactFlowId(v string) *DiscardEditingContactFlowRequest {
	s.ContactFlowId = &v
	return s
}

func (s *DiscardEditingContactFlowRequest) SetDraftId(v string) *DiscardEditingContactFlowRequest {
	s.DraftId = &v
	return s
}

func (s *DiscardEditingContactFlowRequest) SetInstanceId(v string) *DiscardEditingContactFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *DiscardEditingContactFlowRequest) Validate() error {
	return dara.Validate(s)
}
