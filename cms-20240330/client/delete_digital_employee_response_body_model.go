// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDigitalEmployeeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDigitalEmployeeResponseBody
	GetRequestId() *string
}

type DeleteDigitalEmployeeResponseBody struct {
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteDigitalEmployeeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDigitalEmployeeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDigitalEmployeeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDigitalEmployeeResponseBody) SetRequestId(v string) *DeleteDigitalEmployeeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDigitalEmployeeResponseBody) Validate() error {
	return dara.Validate(s)
}
