// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachRCDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachRCDiskResponseBody
	GetRequestId() *string
}

type AttachRCDiskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C19D1668-70CB-5421-AA91-D6D8EE3AB664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachRCDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachRCDiskResponseBody) GoString() string {
	return s.String()
}

func (s *AttachRCDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachRCDiskResponseBody) SetRequestId(v string) *AttachRCDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachRCDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
