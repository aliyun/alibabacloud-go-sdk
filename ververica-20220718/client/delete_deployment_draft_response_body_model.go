// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentDraftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteDeploymentDraftResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDeploymentDraftResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteDeploymentDraftResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteDeploymentDraftResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDeploymentDraftResponseBody
	GetSuccess() *bool
}

type DeleteDeploymentDraftResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDeploymentDraftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentDraftResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentDraftResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDeploymentDraftResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDeploymentDraftResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteDeploymentDraftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeploymentDraftResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDeploymentDraftResponseBody) SetErrorCode(v string) *DeleteDeploymentDraftResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDeploymentDraftResponseBody) SetErrorMessage(v string) *DeleteDeploymentDraftResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDeploymentDraftResponseBody) SetHttpCode(v int32) *DeleteDeploymentDraftResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteDeploymentDraftResponseBody) SetRequestId(v string) *DeleteDeploymentDraftResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeploymentDraftResponseBody) SetSuccess(v bool) *DeleteDeploymentDraftResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDeploymentDraftResponseBody) Validate() error {
	return dara.Validate(s)
}
