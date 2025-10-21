// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSessionClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SessionCluster) *DeleteSessionClusterResponseBody
	GetData() *SessionCluster
	SetErrorCode(v string) *DeleteSessionClusterResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteSessionClusterResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteSessionClusterResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteSessionClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSessionClusterResponseBody
	GetSuccess() *bool
}

type DeleteSessionClusterResponseBody struct {
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
	// CBC799F0-ABCD-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteSessionClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSessionClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSessionClusterResponseBody) GetData() *SessionCluster {
	return s.Data
}

func (s *DeleteSessionClusterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteSessionClusterResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteSessionClusterResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteSessionClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSessionClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSessionClusterResponseBody) SetData(v *SessionCluster) *DeleteSessionClusterResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSessionClusterResponseBody) SetErrorCode(v string) *DeleteSessionClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSessionClusterResponseBody) SetErrorMessage(v string) *DeleteSessionClusterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteSessionClusterResponseBody) SetHttpCode(v int32) *DeleteSessionClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteSessionClusterResponseBody) SetRequestId(v string) *DeleteSessionClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSessionClusterResponseBody) SetSuccess(v bool) *DeleteSessionClusterResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSessionClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
