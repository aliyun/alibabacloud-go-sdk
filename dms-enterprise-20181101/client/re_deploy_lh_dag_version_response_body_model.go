// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReDeployLhDagVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeployId(v int64) *ReDeployLhDagVersionResponseBody
	GetDeployId() *int64
	SetErrorCode(v string) *ReDeployLhDagVersionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ReDeployLhDagVersionResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ReDeployLhDagVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReDeployLhDagVersionResponseBody
	GetSuccess() *bool
}

type ReDeployLhDagVersionResponseBody struct {
	// The ID of the deployment record.
	//
	// example:
	//
	// 15990
	DeployId *int64 `json:"DeployId,omitempty" xml:"DeployId,omitempty"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 035C92E0-2EAD-50E5-A6DD-550F5F73D7CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReDeployLhDagVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReDeployLhDagVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ReDeployLhDagVersionResponseBody) GetDeployId() *int64 {
	return s.DeployId
}

func (s *ReDeployLhDagVersionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ReDeployLhDagVersionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ReDeployLhDagVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReDeployLhDagVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReDeployLhDagVersionResponseBody) SetDeployId(v int64) *ReDeployLhDagVersionResponseBody {
	s.DeployId = &v
	return s
}

func (s *ReDeployLhDagVersionResponseBody) SetErrorCode(v string) *ReDeployLhDagVersionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReDeployLhDagVersionResponseBody) SetErrorMessage(v string) *ReDeployLhDagVersionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReDeployLhDagVersionResponseBody) SetRequestId(v string) *ReDeployLhDagVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReDeployLhDagVersionResponseBody) SetSuccess(v bool) *ReDeployLhDagVersionResponseBody {
	s.Success = &v
	return s
}

func (s *ReDeployLhDagVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
