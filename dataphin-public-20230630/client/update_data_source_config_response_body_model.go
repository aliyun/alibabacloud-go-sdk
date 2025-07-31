// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDataSourceConfigResponseBody
	GetCode() *string
	SetData(v bool) *UpdateDataSourceConfigResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateDataSourceConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDataSourceConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDataSourceConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataSourceConfigResponseBody
	GetSuccess() *bool
}

type UpdateDataSourceConfigResponseBody struct {
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

func (s UpdateDataSourceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDataSourceConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateDataSourceConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDataSourceConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDataSourceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataSourceConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataSourceConfigResponseBody) SetCode(v string) *UpdateDataSourceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDataSourceConfigResponseBody) SetData(v bool) *UpdateDataSourceConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDataSourceConfigResponseBody) SetHttpStatusCode(v int32) *UpdateDataSourceConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDataSourceConfigResponseBody) SetMessage(v string) *UpdateDataSourceConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDataSourceConfigResponseBody) SetRequestId(v string) *UpdateDataSourceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSourceConfigResponseBody) SetSuccess(v bool) *UpdateDataSourceConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataSourceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
