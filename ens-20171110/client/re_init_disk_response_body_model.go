// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReInitDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ReInitDiskResponseBody
	GetCode() *int32
	SetRequestId(v string) *ReInitDiskResponseBody
	GetRequestId() *string
}

type ReInitDiskResponseBody struct {
	// The returned service code. A value of 0 indicates that the operation was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4DD66F05-3116-4BAA-B588-52EB2E7F431D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReInitDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReInitDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ReInitDiskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ReInitDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReInitDiskResponseBody) SetCode(v int32) *ReInitDiskResponseBody {
	s.Code = &v
	return s
}

func (s *ReInitDiskResponseBody) SetRequestId(v string) *ReInitDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReInitDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
