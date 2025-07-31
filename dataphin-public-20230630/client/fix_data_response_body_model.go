// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFixDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FixDataResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *FixDataResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *FixDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *FixDataResponseBody
	GetRequestId() *string
	SetSubmitId(v string) *FixDataResponseBody
	GetSubmitId() *string
	SetSuccess(v bool) *FixDataResponseBody
	GetSuccess() *bool
}

type FixDataResponseBody struct {
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
	// example:
	//
	// 12324234
	SubmitId *string `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
	Success  *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FixDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FixDataResponseBody) GoString() string {
	return s.String()
}

func (s *FixDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *FixDataResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *FixDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FixDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FixDataResponseBody) GetSubmitId() *string {
	return s.SubmitId
}

func (s *FixDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FixDataResponseBody) SetCode(v string) *FixDataResponseBody {
	s.Code = &v
	return s
}

func (s *FixDataResponseBody) SetHttpStatusCode(v int32) *FixDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FixDataResponseBody) SetMessage(v string) *FixDataResponseBody {
	s.Message = &v
	return s
}

func (s *FixDataResponseBody) SetRequestId(v string) *FixDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *FixDataResponseBody) SetSubmitId(v string) *FixDataResponseBody {
	s.SubmitId = &v
	return s
}

func (s *FixDataResponseBody) SetSuccess(v bool) *FixDataResponseBody {
	s.Success = &v
	return s
}

func (s *FixDataResponseBody) Validate() error {
	return dara.Validate(s)
}
