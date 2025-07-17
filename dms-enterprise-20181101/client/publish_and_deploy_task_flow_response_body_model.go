// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishAndDeployTaskFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeployId(v int64) *PublishAndDeployTaskFlowResponseBody
	GetDeployId() *int64
	SetErrorCode(v string) *PublishAndDeployTaskFlowResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *PublishAndDeployTaskFlowResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *PublishAndDeployTaskFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishAndDeployTaskFlowResponseBody
	GetSuccess() *bool
}

type PublishAndDeployTaskFlowResponseBody struct {
	// The ID of the deployment record.
	//
	// example:
	//
	// 12**
	DeployId *int64 `json:"DeployId,omitempty" xml:"DeployId,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Unknown server error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 64E26249-B61F-51C6-B6DF-47EFF50128CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishAndDeployTaskFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishAndDeployTaskFlowResponseBody) GoString() string {
	return s.String()
}

func (s *PublishAndDeployTaskFlowResponseBody) GetDeployId() *int64 {
	return s.DeployId
}

func (s *PublishAndDeployTaskFlowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PublishAndDeployTaskFlowResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PublishAndDeployTaskFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishAndDeployTaskFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishAndDeployTaskFlowResponseBody) SetDeployId(v int64) *PublishAndDeployTaskFlowResponseBody {
	s.DeployId = &v
	return s
}

func (s *PublishAndDeployTaskFlowResponseBody) SetErrorCode(v string) *PublishAndDeployTaskFlowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PublishAndDeployTaskFlowResponseBody) SetErrorMessage(v string) *PublishAndDeployTaskFlowResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PublishAndDeployTaskFlowResponseBody) SetRequestId(v string) *PublishAndDeployTaskFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishAndDeployTaskFlowResponseBody) SetSuccess(v bool) *PublishAndDeployTaskFlowResponseBody {
	s.Success = &v
	return s
}

func (s *PublishAndDeployTaskFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
