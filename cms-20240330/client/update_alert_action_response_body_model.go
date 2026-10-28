// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAlertActionResponseBody
	GetRequestId() *string
}

type UpdateAlertActionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAlertActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertActionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlertActionResponseBody) SetRequestId(v string) *UpdateAlertActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertActionResponseBody) Validate() error {
	return dara.Validate(s)
}
