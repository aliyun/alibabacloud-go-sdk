// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachRCDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachRCDiskResponseBody
	GetRequestId() *string
}

type DetachRCDiskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C8E88DED-533F-4B3C-9207-731FBF394CCA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachRCDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachRCDiskResponseBody) GoString() string {
	return s.String()
}

func (s *DetachRCDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachRCDiskResponseBody) SetRequestId(v string) *DetachRCDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachRCDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
