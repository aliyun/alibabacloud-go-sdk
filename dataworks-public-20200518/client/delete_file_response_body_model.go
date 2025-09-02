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
	// The ID of the request. You can troubleshoot errors based on the ID.
	//
	// example:
	//
	// 1000000001
	DeploymentId *int64 `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	// The ID of the file. You can use the ListFiles interface to query the ID of the corresponding file.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the deployment task that deploys the file. If the file has been committed, an asynchronous process is triggered to delete the file in the scheduling system. The value of this parameter is used to call the GetDeployment operation to poll the status of the asynchronous process.
	//
	// If this parameter is empty, the file is deleted and the polling is not required.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error code returned.
	//
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
