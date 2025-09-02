// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterConfigs(v []*ClusterConfig) *ListClusterConfigsResponseBody
	GetClusterConfigs() []*ClusterConfig
	SetErrorCode(v string) *ListClusterConfigsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListClusterConfigsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListClusterConfigsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListClusterConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListClusterConfigsResponseBody
	GetSuccess() *bool
}

type ListClusterConfigsResponseBody struct {
	// A list of configurations of cluster modules.
	ClusterConfigs []*ClusterConfig `json:"ClusterConfigs,omitempty" xml:"ClusterConfigs,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 101011005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid.Cluster.ClusterNotFound
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 0000-ABCD-E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListClusterConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterConfigsResponseBody) GetClusterConfigs() []*ClusterConfig {
	return s.ClusterConfigs
}

func (s *ListClusterConfigsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListClusterConfigsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListClusterConfigsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListClusterConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListClusterConfigsResponseBody) SetClusterConfigs(v []*ClusterConfig) *ListClusterConfigsResponseBody {
	s.ClusterConfigs = v
	return s
}

func (s *ListClusterConfigsResponseBody) SetErrorCode(v string) *ListClusterConfigsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListClusterConfigsResponseBody) SetErrorMessage(v string) *ListClusterConfigsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListClusterConfigsResponseBody) SetHttpStatusCode(v int32) *ListClusterConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListClusterConfigsResponseBody) SetRequestId(v string) *ListClusterConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterConfigsResponseBody) SetSuccess(v bool) *ListClusterConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListClusterConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
