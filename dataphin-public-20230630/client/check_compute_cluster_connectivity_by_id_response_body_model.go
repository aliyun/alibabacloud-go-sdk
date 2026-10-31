// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckComputeClusterConnectivityByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckComputeClusterConnectivityByIdResponseBody
	GetCode() *string
	SetData(v bool) *CheckComputeClusterConnectivityByIdResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *CheckComputeClusterConnectivityByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CheckComputeClusterConnectivityByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckComputeClusterConnectivityByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckComputeClusterConnectivityByIdResponseBody
	GetSuccess() *bool
}

type CheckComputeClusterConnectivityByIdResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The cluster connectivity check result.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend response error details.
	//
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
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckComputeClusterConnectivityByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckComputeClusterConnectivityByIdResponseBody) GoString() string {
	return s.String()
}

func (s *CheckComputeClusterConnectivityByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckComputeClusterConnectivityByIdResponseBody) GetData() *bool {
	return s.Data
}

func (s *CheckComputeClusterConnectivityByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckComputeClusterConnectivityByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckComputeClusterConnectivityByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckComputeClusterConnectivityByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckComputeClusterConnectivityByIdResponseBody) SetCode(v string) *CheckComputeClusterConnectivityByIdResponseBody {
	s.Code = &v
	return s
}

func (s *CheckComputeClusterConnectivityByIdResponseBody) SetData(v bool) *CheckComputeClusterConnectivityByIdResponseBody {
	s.Data = &v
	return s
}

func (s *CheckComputeClusterConnectivityByIdResponseBody) SetHttpStatusCode(v int32) *CheckComputeClusterConnectivityByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckComputeClusterConnectivityByIdResponseBody) SetMessage(v string) *CheckComputeClusterConnectivityByIdResponseBody {
	s.Message = &v
	return s
}

func (s *CheckComputeClusterConnectivityByIdResponseBody) SetRequestId(v string) *CheckComputeClusterConnectivityByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckComputeClusterConnectivityByIdResponseBody) SetSuccess(v bool) *CheckComputeClusterConnectivityByIdResponseBody {
	s.Success = &v
	return s
}

func (s *CheckComputeClusterConnectivityByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
