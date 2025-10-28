// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackChangeOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *RollbackChangeOrderRequest
	GetChangeOrderId() *string
}

type RollbackChangeOrderRequest struct {
	// The ID of the change process.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc5133d7-773f-4c81-****-e2103dce****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s RollbackChangeOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackChangeOrderRequest) GoString() string {
	return s.String()
}

func (s *RollbackChangeOrderRequest) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *RollbackChangeOrderRequest) SetChangeOrderId(v string) *RollbackChangeOrderRequest {
	s.ChangeOrderId = &v
	return s
}

func (s *RollbackChangeOrderRequest) Validate() error {
	return dara.Validate(s)
}
