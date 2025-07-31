// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckComputeSourceConnectivityByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckComputeSourceConnectivityByIdResponseBody
	GetCode() *string
	SetData(v bool) *CheckComputeSourceConnectivityByIdResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *CheckComputeSourceConnectivityByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CheckComputeSourceConnectivityByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckComputeSourceConnectivityByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckComputeSourceConnectivityByIdResponseBody
	GetSuccess() *bool
}

type CheckComputeSourceConnectivityByIdResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CheckComputeSourceConnectivityByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckComputeSourceConnectivityByIdResponseBody) GoString() string {
	return s.String()
}

func (s *CheckComputeSourceConnectivityByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckComputeSourceConnectivityByIdResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckComputeSourceConnectivityByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckComputeSourceConnectivityByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckComputeSourceConnectivityByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckComputeSourceConnectivityByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckComputeSourceConnectivityByIdResponseBody) SetCode(v string) *CheckComputeSourceConnectivityByIdResponseBody {
	s.Code = &v
	return s
}

func (s *CheckComputeSourceConnectivityByIdResponseBody) SetData(v bool) *CheckComputeSourceConnectivityByIdResponseBody {
	s.Data = &v
	return s
}

func (s *CheckComputeSourceConnectivityByIdResponseBody) SetHttpStatusCode(v int32) *CheckComputeSourceConnectivityByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckComputeSourceConnectivityByIdResponseBody) SetMessage(v string) *CheckComputeSourceConnectivityByIdResponseBody {
	s.Message = &v
	return s
}

func (s *CheckComputeSourceConnectivityByIdResponseBody) SetRequestId(v string) *CheckComputeSourceConnectivityByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckComputeSourceConnectivityByIdResponseBody) SetSuccess(v bool) *CheckComputeSourceConnectivityByIdResponseBody {
	s.Success = &v
	return s
}

func (s *CheckComputeSourceConnectivityByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
