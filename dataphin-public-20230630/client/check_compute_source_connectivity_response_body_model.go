// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckComputeSourceConnectivityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckComputeSourceConnectivityResponseBody
	GetCode() *string
	SetData(v bool) *CheckComputeSourceConnectivityResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *CheckComputeSourceConnectivityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CheckComputeSourceConnectivityResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckComputeSourceConnectivityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckComputeSourceConnectivityResponseBody
	GetSuccess() *bool
}

type CheckComputeSourceConnectivityResponseBody struct {
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

func (s CheckComputeSourceConnectivityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckComputeSourceConnectivityResponseBody) GoString() string {
	return s.String()
}

func (s *CheckComputeSourceConnectivityResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckComputeSourceConnectivityResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckComputeSourceConnectivityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckComputeSourceConnectivityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckComputeSourceConnectivityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckComputeSourceConnectivityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckComputeSourceConnectivityResponseBody) SetCode(v string) *CheckComputeSourceConnectivityResponseBody {
	s.Code = &v
	return s
}

func (s *CheckComputeSourceConnectivityResponseBody) SetData(v bool) *CheckComputeSourceConnectivityResponseBody {
	s.Data = &v
	return s
}

func (s *CheckComputeSourceConnectivityResponseBody) SetHttpStatusCode(v int32) *CheckComputeSourceConnectivityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckComputeSourceConnectivityResponseBody) SetMessage(v string) *CheckComputeSourceConnectivityResponseBody {
	s.Message = &v
	return s
}

func (s *CheckComputeSourceConnectivityResponseBody) SetRequestId(v string) *CheckComputeSourceConnectivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckComputeSourceConnectivityResponseBody) SetSuccess(v bool) *CheckComputeSourceConnectivityResponseBody {
	s.Success = &v
	return s
}

func (s *CheckComputeSourceConnectivityResponseBody) Validate() error {
	return dara.Validate(s)
}
