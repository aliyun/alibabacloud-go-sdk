// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataSourceConnectivityByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckDataSourceConnectivityByIdResponseBody
	GetCode() *string
	SetData(v bool) *CheckDataSourceConnectivityByIdResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *CheckDataSourceConnectivityByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CheckDataSourceConnectivityByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckDataSourceConnectivityByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckDataSourceConnectivityByIdResponseBody
	GetSuccess() *bool
}

type CheckDataSourceConnectivityByIdResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The connectivity check result of the data source.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckDataSourceConnectivityByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityByIdResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckDataSourceConnectivityByIdResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckDataSourceConnectivityByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckDataSourceConnectivityByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckDataSourceConnectivityByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDataSourceConnectivityByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckDataSourceConnectivityByIdResponseBody) SetCode(v string) *CheckDataSourceConnectivityByIdResponseBody {
	s.Code = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponseBody) SetData(v bool) *CheckDataSourceConnectivityByIdResponseBody {
	s.Data = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponseBody) SetHttpStatusCode(v int32) *CheckDataSourceConnectivityByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponseBody) SetMessage(v string) *CheckDataSourceConnectivityByIdResponseBody {
	s.Message = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponseBody) SetRequestId(v string) *CheckDataSourceConnectivityByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponseBody) SetSuccess(v bool) *CheckDataSourceConnectivityByIdResponseBody {
	s.Success = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
