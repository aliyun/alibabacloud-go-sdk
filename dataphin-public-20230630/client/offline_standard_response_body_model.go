// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineStandardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OfflineStandardResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *OfflineStandardResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *OfflineStandardResponseBody
	GetMessage() *string
	SetRequestId(v string) *OfflineStandardResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OfflineStandardResponseBody
	GetSuccess() *bool
}

type OfflineStandardResponseBody struct {
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
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OfflineStandardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineStandardResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineStandardResponseBody) GetCode() *string {
	return s.Code
}

func (s *OfflineStandardResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OfflineStandardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OfflineStandardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineStandardResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OfflineStandardResponseBody) SetCode(v string) *OfflineStandardResponseBody {
	s.Code = &v
	return s
}

func (s *OfflineStandardResponseBody) SetHttpStatusCode(v int32) *OfflineStandardResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OfflineStandardResponseBody) SetMessage(v string) *OfflineStandardResponseBody {
	s.Message = &v
	return s
}

func (s *OfflineStandardResponseBody) SetRequestId(v string) *OfflineStandardResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineStandardResponseBody) SetSuccess(v bool) *OfflineStandardResponseBody {
	s.Success = &v
	return s
}

func (s *OfflineStandardResponseBody) Validate() error {
	return dara.Validate(s)
}
