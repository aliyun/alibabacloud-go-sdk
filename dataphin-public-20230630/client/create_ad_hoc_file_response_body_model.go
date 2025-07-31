// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdHocFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAdHocFileResponseBody
	GetCode() *string
	SetFileId(v int64) *CreateAdHocFileResponseBody
	GetFileId() *int64
	SetHttpStatusCode(v int32) *CreateAdHocFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateAdHocFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAdHocFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAdHocFileResponseBody
	GetSuccess() *bool
}

type CreateAdHocFileResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1212313222
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
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
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAdHocFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAdHocFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAdHocFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAdHocFileResponseBody) GetFileId() *int64 {
	return s.FileId
}

func (s *CreateAdHocFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAdHocFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAdHocFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAdHocFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAdHocFileResponseBody) SetCode(v string) *CreateAdHocFileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAdHocFileResponseBody) SetFileId(v int64) *CreateAdHocFileResponseBody {
	s.FileId = &v
	return s
}

func (s *CreateAdHocFileResponseBody) SetHttpStatusCode(v int32) *CreateAdHocFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAdHocFileResponseBody) SetMessage(v string) *CreateAdHocFileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAdHocFileResponseBody) SetRequestId(v string) *CreateAdHocFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAdHocFileResponseBody) SetSuccess(v bool) *CreateAdHocFileResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAdHocFileResponseBody) Validate() error {
	return dara.Validate(s)
}
