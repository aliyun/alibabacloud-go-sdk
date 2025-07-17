// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *DeployFileResponseBody
	GetData() *int64
	SetErrorCode(v string) *DeployFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeployFileResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeployFileResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeployFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeployFileResponseBody
	GetSuccess() *bool
}

type DeployFileResponseBody struct {
	// example:
	//
	// 30000001
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeployFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeployFileResponseBody) GetData() *int64 {
	return s.Data
}

func (s *DeployFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeployFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeployFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeployFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeployFileResponseBody) SetData(v int64) *DeployFileResponseBody {
	s.Data = &v
	return s
}

func (s *DeployFileResponseBody) SetErrorCode(v string) *DeployFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeployFileResponseBody) SetErrorMessage(v string) *DeployFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeployFileResponseBody) SetHttpStatusCode(v int32) *DeployFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeployFileResponseBody) SetRequestId(v string) *DeployFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployFileResponseBody) SetSuccess(v bool) *DeployFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeployFileResponseBody) Validate() error {
	return dara.Validate(s)
}
