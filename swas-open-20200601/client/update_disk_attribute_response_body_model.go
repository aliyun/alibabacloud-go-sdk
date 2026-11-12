// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDiskAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDiskAttributeResponseBody
	GetRequestId() *string
}

type UpdateDiskAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28D****534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDiskAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDiskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDiskAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDiskAttributeResponseBody) SetRequestId(v string) *UpdateDiskAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDiskAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
