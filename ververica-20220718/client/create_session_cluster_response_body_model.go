// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SessionCluster) *CreateSessionClusterResponseBody
	GetData() *SessionCluster
	SetErrorCode(v string) *CreateSessionClusterResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateSessionClusterResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *CreateSessionClusterResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CreateSessionClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSessionClusterResponseBody
	GetSuccess() *bool
}

type CreateSessionClusterResponseBody struct {
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
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateSessionClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterResponseBody) GetData() *SessionCluster {
	return s.Data
}

func (s *CreateSessionClusterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateSessionClusterResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateSessionClusterResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateSessionClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSessionClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSessionClusterResponseBody) SetData(v *SessionCluster) *CreateSessionClusterResponseBody {
	s.Data = v
	return s
}

func (s *CreateSessionClusterResponseBody) SetErrorCode(v string) *CreateSessionClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSessionClusterResponseBody) SetErrorMessage(v string) *CreateSessionClusterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSessionClusterResponseBody) SetHttpCode(v int32) *CreateSessionClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateSessionClusterResponseBody) SetRequestId(v string) *CreateSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSessionClusterResponseBody) SetSuccess(v bool) *CreateSessionClusterResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSessionClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
