// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationMaintainConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyActiveOperationMaintainConfigResponseBody
	GetRequestId() *string
}

type ModifyActiveOperationMaintainConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 561AFBF1-BE20-44DB-9BD1-6988B53E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyActiveOperationMaintainConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationMaintainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationMaintainConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyActiveOperationMaintainConfigResponseBody) SetRequestId(v string) *ModifyActiveOperationMaintainConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
