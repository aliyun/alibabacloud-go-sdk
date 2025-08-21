// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnsServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateEnsServiceResponseBody
	GetCode() *int32
	SetRequestId(v string) *CreateEnsServiceResponseBody
	GetRequestId() *string
}

type CreateEnsServiceResponseBody struct {
	// The service code. 0 is returned for a successful request. An error code is returned for a failed request.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AF02B43-2D08-49D3-8AAF-65B9C792ED14
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEnsServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEnsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnsServiceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateEnsServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEnsServiceResponseBody) SetCode(v int32) *CreateEnsServiceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEnsServiceResponseBody) SetRequestId(v string) *CreateEnsServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnsServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
