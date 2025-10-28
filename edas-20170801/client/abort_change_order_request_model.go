// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortChangeOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *AbortChangeOrderRequest
	GetChangeOrderId() *string
}

type AbortChangeOrderRequest struct {
	// The ID of the change process. You can call the [GetChangeOrderInfo](https://help.aliyun.com/document_detail/62072.html) operation to query the change process ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4f038ddf-b27b-****-****-88e44375****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s AbortChangeOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s AbortChangeOrderRequest) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderRequest) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *AbortChangeOrderRequest) SetChangeOrderId(v string) *AbortChangeOrderRequest {
	s.ChangeOrderId = &v
	return s
}

func (s *AbortChangeOrderRequest) Validate() error {
	return dara.Validate(s)
}
