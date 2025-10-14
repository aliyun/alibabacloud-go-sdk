// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStoragePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateStoragePoolResponseBody
	GetCode() *string
	SetRequestId(v string) *CreateStoragePoolResponseBody
	GetRequestId() *string
}

type CreateStoragePoolResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateStoragePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStoragePoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStoragePoolResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateStoragePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStoragePoolResponseBody) SetCode(v string) *CreateStoragePoolResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStoragePoolResponseBody) SetRequestId(v string) *CreateStoragePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStoragePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
