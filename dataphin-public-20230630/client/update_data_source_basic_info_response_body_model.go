// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceBasicInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDataSourceBasicInfoResponseBody
	GetCode() *string
	SetData(v bool) *UpdateDataSourceBasicInfoResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateDataSourceBasicInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDataSourceBasicInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDataSourceBasicInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataSourceBasicInfoResponseBody
	GetSuccess() *bool
}

type UpdateDataSourceBasicInfoResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s UpdateDataSourceBasicInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceBasicInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDataSourceBasicInfoResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateDataSourceBasicInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDataSourceBasicInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDataSourceBasicInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataSourceBasicInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataSourceBasicInfoResponseBody) SetCode(v string) *UpdateDataSourceBasicInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDataSourceBasicInfoResponseBody) SetData(v bool) *UpdateDataSourceBasicInfoResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDataSourceBasicInfoResponseBody) SetHttpStatusCode(v int32) *UpdateDataSourceBasicInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDataSourceBasicInfoResponseBody) SetMessage(v string) *UpdateDataSourceBasicInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDataSourceBasicInfoResponseBody) SetRequestId(v string) *UpdateDataSourceBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSourceBasicInfoResponseBody) SetSuccess(v bool) *UpdateDataSourceBasicInfoResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataSourceBasicInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
