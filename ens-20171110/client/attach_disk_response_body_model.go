// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AttachDiskResponseBody
	GetCode() *int32
	SetRequestId(v string) *AttachDiskResponseBody
	GetRequestId() *string
}

type AttachDiskResponseBody struct {
	// The returned service code. 0 indicates that the request was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 916777D9-42D3-5928-92CE-373B1874B674
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachDiskResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDiskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AttachDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachDiskResponseBody) SetCode(v int32) *AttachDiskResponseBody {
	s.Code = &v
	return s
}

func (s *AttachDiskResponseBody) SetRequestId(v string) *AttachDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
