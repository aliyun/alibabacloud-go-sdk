// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteDiskResponseBody
	GetCode() *int32
	SetRequestId(v string) *DeleteDiskResponseBody
	GetRequestId() *string
}

type DeleteDiskResponseBody struct {
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
	// 3ABEEB76-1976-55AB-B884-3D65CA6A4743
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDiskResponseBody) SetCode(v int32) *DeleteDiskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDiskResponseBody) SetRequestId(v string) *DeleteDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
