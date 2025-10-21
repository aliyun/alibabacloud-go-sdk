// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SessionCluster) *UpdateSessionClusterResponseBody
	GetData() *SessionCluster
	SetErrorCode(v string) *UpdateSessionClusterResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateSessionClusterResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *UpdateSessionClusterResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateSessionClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSessionClusterResponseBody
	GetSuccess() *bool
}

type UpdateSessionClusterResponseBody struct {
	Data *SessionCluster `json:"data,omitempty" xml:"data,omitempty"`
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
	// 1EF03B0C-F44F-47AD-BB48-D002D0F7B8C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateSessionClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSessionClusterResponseBody) GetData() *SessionCluster {
	return s.Data
}

func (s *UpdateSessionClusterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateSessionClusterResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateSessionClusterResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateSessionClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSessionClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSessionClusterResponseBody) SetData(v *SessionCluster) *UpdateSessionClusterResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSessionClusterResponseBody) SetErrorCode(v string) *UpdateSessionClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateSessionClusterResponseBody) SetErrorMessage(v string) *UpdateSessionClusterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateSessionClusterResponseBody) SetHttpCode(v int32) *UpdateSessionClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateSessionClusterResponseBody) SetRequestId(v string) *UpdateSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSessionClusterResponseBody) SetSuccess(v bool) *UpdateSessionClusterResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSessionClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
