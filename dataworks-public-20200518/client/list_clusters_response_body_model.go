// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListClustersResponseBodyData) *ListClustersResponseBody
	GetData() *ListClustersResponseBodyData
	SetErrorCode(v string) *ListClustersResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListClustersResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListClustersResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListClustersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListClustersResponseBody
	GetSuccess() *bool
}

type ListClustersResponseBody struct {
	// The returns data.
	Data *ListClustersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The request ID. You can use the request ID to query logs and troubleshoot issues.
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

func (s ListClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) GetData() *ListClustersResponseBodyData {
	return s.Data
}

func (s *ListClustersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListClustersResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListClustersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClustersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListClustersResponseBody) SetData(v *ListClustersResponseBodyData) *ListClustersResponseBody {
	s.Data = v
	return s
}

func (s *ListClustersResponseBody) SetErrorCode(v string) *ListClustersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListClustersResponseBody) SetErrorMessage(v string) *ListClustersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListClustersResponseBody) SetHttpStatusCode(v int32) *ListClustersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetSuccess(v bool) *ListClustersResponseBody {
	s.Success = &v
	return s
}

func (s *ListClustersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyData struct {
	// List of cluster information.
	Clusters []*Cluster `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClustersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyData) GetClusters() []*Cluster {
	return s.Clusters
}

func (s *ListClustersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClustersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClustersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListClustersResponseBodyData) SetClusters(v []*Cluster) *ListClustersResponseBodyData {
	s.Clusters = v
	return s
}

func (s *ListClustersResponseBodyData) SetPageNumber(v int32) *ListClustersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListClustersResponseBodyData) SetPageSize(v int32) *ListClustersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListClustersResponseBodyData) SetTotalCount(v int32) *ListClustersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListClustersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
