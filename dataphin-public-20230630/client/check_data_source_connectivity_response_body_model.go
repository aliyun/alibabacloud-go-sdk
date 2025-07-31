// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataSourceConnectivityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckDataSourceConnectivityResponseBody
	GetCode() *string
	SetData(v bool) *CheckDataSourceConnectivityResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *CheckDataSourceConnectivityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CheckDataSourceConnectivityResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckDataSourceConnectivityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckDataSourceConnectivityResponseBody
	GetSuccess() *bool
}

type CheckDataSourceConnectivityResponseBody struct {
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

func (s CheckDataSourceConnectivityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDataSourceConnectivityResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckDataSourceConnectivityResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckDataSourceConnectivityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckDataSourceConnectivityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckDataSourceConnectivityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDataSourceConnectivityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckDataSourceConnectivityResponseBody) SetCode(v string) *CheckDataSourceConnectivityResponseBody {
	s.Code = &v
	return s
}

func (s *CheckDataSourceConnectivityResponseBody) SetData(v bool) *CheckDataSourceConnectivityResponseBody {
	s.Data = &v
	return s
}

func (s *CheckDataSourceConnectivityResponseBody) SetHttpStatusCode(v int32) *CheckDataSourceConnectivityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckDataSourceConnectivityResponseBody) SetMessage(v string) *CheckDataSourceConnectivityResponseBody {
	s.Message = &v
	return s
}

func (s *CheckDataSourceConnectivityResponseBody) SetRequestId(v string) *CheckDataSourceConnectivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDataSourceConnectivityResponseBody) SetSuccess(v bool) *CheckDataSourceConnectivityResponseBody {
	s.Success = &v
	return s
}

func (s *CheckDataSourceConnectivityResponseBody) Validate() error {
	return dara.Validate(s)
}
