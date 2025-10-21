// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*SessionCluster) *ListSessionClustersResponseBody
	GetData() []*SessionCluster
	SetErrorCode(v string) *ListSessionClustersResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSessionClustersResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListSessionClustersResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListSessionClustersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSessionClustersResponseBody
	GetSuccess() *bool
}

type ListSessionClustersResponseBody struct {
	Data []*SessionCluster `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s ListSessionClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSessionClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListSessionClustersResponseBody) GetData() []*SessionCluster {
	return s.Data
}

func (s *ListSessionClustersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSessionClustersResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSessionClustersResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListSessionClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSessionClustersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSessionClustersResponseBody) SetData(v []*SessionCluster) *ListSessionClustersResponseBody {
	s.Data = v
	return s
}

func (s *ListSessionClustersResponseBody) SetErrorCode(v string) *ListSessionClustersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetErrorMessage(v string) *ListSessionClustersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetHttpCode(v int32) *ListSessionClustersResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetRequestId(v string) *ListSessionClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSessionClustersResponseBody) SetSuccess(v bool) *ListSessionClustersResponseBody {
	s.Success = &v
	return s
}

func (s *ListSessionClustersResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
