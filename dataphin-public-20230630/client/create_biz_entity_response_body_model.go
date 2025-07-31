// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateBizEntityResponseBody
	GetCode() *string
	SetCreateResult(v *CreateBizEntityResponseBodyCreateResult) *CreateBizEntityResponseBody
	GetCreateResult() *CreateBizEntityResponseBodyCreateResult
	SetHttpStatusCode(v int32) *CreateBizEntityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateBizEntityResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBizEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBizEntityResponseBody
	GetSuccess() *bool
}

type CreateBizEntityResponseBody struct {
	// example:
	//
	// OK
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateResult *CreateBizEntityResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
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

func (s CreateBizEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBizEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBizEntityResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBizEntityResponseBody) GetCreateResult() *CreateBizEntityResponseBodyCreateResult {
	return s.CreateResult
}

func (s *CreateBizEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBizEntityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBizEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBizEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBizEntityResponseBody) SetCode(v string) *CreateBizEntityResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBizEntityResponseBody) SetCreateResult(v *CreateBizEntityResponseBodyCreateResult) *CreateBizEntityResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateBizEntityResponseBody) SetHttpStatusCode(v int32) *CreateBizEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBizEntityResponseBody) SetMessage(v string) *CreateBizEntityResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBizEntityResponseBody) SetRequestId(v string) *CreateBizEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBizEntityResponseBody) SetSuccess(v bool) *CreateBizEntityResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBizEntityResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateBizEntityResponseBodyCreateResult struct {
	// example:
	//
	// 12113111
	BizEntityId *int64 `json:"BizEntityId,omitempty" xml:"BizEntityId,omitempty"`
}

func (s CreateBizEntityResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s CreateBizEntityResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateBizEntityResponseBodyCreateResult) GetBizEntityId() *int64 {
	return s.BizEntityId
}

func (s *CreateBizEntityResponseBodyCreateResult) SetBizEntityId(v int64) *CreateBizEntityResponseBodyCreateResult {
	s.BizEntityId = &v
	return s
}

func (s *CreateBizEntityResponseBodyCreateResult) Validate() error {
	return dara.Validate(s)
}
