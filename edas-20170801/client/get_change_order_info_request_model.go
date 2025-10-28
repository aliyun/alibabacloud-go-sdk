// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChangeOrderInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *GetChangeOrderInfoRequest
	GetChangeOrderId() *string
}

type GetChangeOrderInfoRequest struct {
	// The ID of the change process.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1074f3e2-e974-4a0e-****-************
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s GetChangeOrderInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoRequest) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoRequest) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *GetChangeOrderInfoRequest) SetChangeOrderId(v string) *GetChangeOrderInfoRequest {
	s.ChangeOrderId = &v
	return s
}

func (s *GetChangeOrderInfoRequest) Validate() error {
	return dara.Validate(s)
}
