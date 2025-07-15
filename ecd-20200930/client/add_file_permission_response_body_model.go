// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFilePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddFilePermissionResponseBody
	GetRequestId() *string
}

type AddFilePermissionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFilePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFilePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *AddFilePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFilePermissionResponseBody) SetRequestId(v string) *AddFilePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFilePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
