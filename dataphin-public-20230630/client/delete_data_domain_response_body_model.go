// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDataDomainResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteDataDomainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDataDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDataDomainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataDomainResponseBody
	GetSuccess() *bool
}

type DeleteDataDomainResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataDomainResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDataDomainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDataDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDataDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataDomainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataDomainResponseBody) SetCode(v string) *DeleteDataDomainResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDataDomainResponseBody) SetHttpStatusCode(v int32) *DeleteDataDomainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDataDomainResponseBody) SetMessage(v string) *DeleteDataDomainResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDataDomainResponseBody) SetRequestId(v string) *DeleteDataDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataDomainResponseBody) SetSuccess(v bool) *DeleteDataDomainResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
