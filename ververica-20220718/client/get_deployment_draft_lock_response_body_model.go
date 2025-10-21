// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentDraftLockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Lock) *GetDeploymentDraftLockResponseBody
	GetData() *Lock
	SetErrorCode(v string) *GetDeploymentDraftLockResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDeploymentDraftLockResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetDeploymentDraftLockResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetDeploymentDraftLockResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDeploymentDraftLockResponseBody
	GetSuccess() *bool
}

type GetDeploymentDraftLockResponseBody struct {
	Data *Lock `json:"data,omitempty" xml:"data,omitempty"`
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

func (s GetDeploymentDraftLockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentDraftLockResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftLockResponseBody) GetData() *Lock {
	return s.Data
}

func (s *GetDeploymentDraftLockResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDeploymentDraftLockResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDeploymentDraftLockResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetDeploymentDraftLockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeploymentDraftLockResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeploymentDraftLockResponseBody) SetData(v *Lock) *GetDeploymentDraftLockResponseBody {
	s.Data = v
	return s
}

func (s *GetDeploymentDraftLockResponseBody) SetErrorCode(v string) *GetDeploymentDraftLockResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDeploymentDraftLockResponseBody) SetErrorMessage(v string) *GetDeploymentDraftLockResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDeploymentDraftLockResponseBody) SetHttpCode(v int32) *GetDeploymentDraftLockResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDeploymentDraftLockResponseBody) SetRequestId(v string) *GetDeploymentDraftLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeploymentDraftLockResponseBody) SetSuccess(v bool) *GetDeploymentDraftLockResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeploymentDraftLockResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
