// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v int64) *DeleteFileResponseBody
	GetDeploymentId() *int64
	SetErrorCode(v string) *DeleteFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteFileResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteFileResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFileResponseBody
	GetSuccess() *bool
}

type DeleteFileResponseBody struct {
	// example:
	//
	// 1000000001
	DeploymentId *int64 `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
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

func (s DeleteFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) GetDeploymentId() *int64 {
	return s.DeploymentId
}

func (s *DeleteFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFileResponseBody) SetDeploymentId(v int64) *DeleteFileResponseBody {
	s.DeploymentId = &v
	return s
}

func (s *DeleteFileResponseBody) SetErrorCode(v string) *DeleteFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFileResponseBody) SetErrorMessage(v string) *DeleteFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFileResponseBody) SetHttpStatusCode(v int32) *DeleteFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteFileResponseBody) SetRequestId(v string) *DeleteFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileResponseBody) SetSuccess(v bool) *DeleteFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFileResponseBody) Validate() error {
	return dara.Validate(s)
}
