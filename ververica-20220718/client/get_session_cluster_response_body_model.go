// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SessionCluster) *GetSessionClusterResponseBody
	GetData() *SessionCluster
	SetErrorCode(v string) *GetSessionClusterResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetSessionClusterResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetSessionClusterResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetSessionClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSessionClusterResponseBody
	GetSuccess() *bool
}

type GetSessionClusterResponseBody struct {
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

func (s GetSessionClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionClusterResponseBody) GetData() *SessionCluster {
	return s.Data
}

func (s *GetSessionClusterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSessionClusterResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSessionClusterResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetSessionClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSessionClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSessionClusterResponseBody) SetData(v *SessionCluster) *GetSessionClusterResponseBody {
	s.Data = v
	return s
}

func (s *GetSessionClusterResponseBody) SetErrorCode(v string) *GetSessionClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSessionClusterResponseBody) SetErrorMessage(v string) *GetSessionClusterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSessionClusterResponseBody) SetHttpCode(v int32) *GetSessionClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetSessionClusterResponseBody) SetRequestId(v string) *GetSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSessionClusterResponseBody) SetSuccess(v bool) *GetSessionClusterResponseBody {
	s.Success = &v
	return s
}

func (s *GetSessionClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
