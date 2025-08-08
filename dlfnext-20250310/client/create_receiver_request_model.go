// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReceiverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateReceiverRequest
	GetComment() *string
	SetReceiverName(v string) *CreateReceiverRequest
	GetReceiverName() *string
	SetReceiverTenantId(v int64) *CreateReceiverRequest
	GetReceiverTenantId() *int64
}

type CreateReceiverRequest struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// receiver_name
	ReceiverName *string `json:"receiverName,omitempty" xml:"receiverName,omitempty"`
	// example:
	//
	// 1111
	ReceiverTenantId *int64 `json:"receiverTenantId,omitempty" xml:"receiverTenantId,omitempty"`
}

func (s CreateReceiverRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReceiverRequest) GoString() string {
	return s.String()
}

func (s *CreateReceiverRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateReceiverRequest) GetReceiverName() *string {
	return s.ReceiverName
}

func (s *CreateReceiverRequest) GetReceiverTenantId() *int64 {
	return s.ReceiverTenantId
}

func (s *CreateReceiverRequest) SetComment(v string) *CreateReceiverRequest {
	s.Comment = &v
	return s
}

func (s *CreateReceiverRequest) SetReceiverName(v string) *CreateReceiverRequest {
	s.ReceiverName = &v
	return s
}

func (s *CreateReceiverRequest) SetReceiverTenantId(v int64) *CreateReceiverRequest {
	s.ReceiverTenantId = &v
	return s
}

func (s *CreateReceiverRequest) Validate() error {
	return dara.Validate(s)
}
