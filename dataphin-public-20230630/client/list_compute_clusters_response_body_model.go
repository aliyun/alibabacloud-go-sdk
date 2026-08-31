// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListComputeClustersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListComputeClustersResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListComputeClustersResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListComputeClustersResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListComputeClustersResponseBody
	GetNextToken() *string
	SetPageResult(v *ListComputeClustersResponseBodyPageResult) *ListComputeClustersResponseBody
	GetPageResult() *ListComputeClustersResponseBodyPageResult
	SetRequestId(v string) *ListComputeClustersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListComputeClustersResponseBody
	GetSuccess() *bool
}

type ListComputeClustersResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The page size. Valid values: 1 to 50. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The details of the backend exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether a token for the next query exists. If NextToken is empty, no more results are available. If NextToken is returned, the value indicates the token used to start the next query.
	//
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh3adOhYj269roQctwr/Eik+
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The paged query result.
	PageResult *ListComputeClustersResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListComputeClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComputeClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeClustersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListComputeClustersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListComputeClustersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComputeClustersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListComputeClustersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComputeClustersResponseBody) GetPageResult() *ListComputeClustersResponseBodyPageResult {
	return s.PageResult
}

func (s *ListComputeClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComputeClustersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListComputeClustersResponseBody) SetCode(v string) *ListComputeClustersResponseBody {
	s.Code = &v
	return s
}

func (s *ListComputeClustersResponseBody) SetHttpStatusCode(v int32) *ListComputeClustersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListComputeClustersResponseBody) SetMaxResults(v int32) *ListComputeClustersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListComputeClustersResponseBody) SetMessage(v string) *ListComputeClustersResponseBody {
	s.Message = &v
	return s
}

func (s *ListComputeClustersResponseBody) SetNextToken(v string) *ListComputeClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListComputeClustersResponseBody) SetPageResult(v *ListComputeClustersResponseBodyPageResult) *ListComputeClustersResponseBody {
	s.PageResult = v
	return s
}

func (s *ListComputeClustersResponseBody) SetRequestId(v string) *ListComputeClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeClustersResponseBody) SetSuccess(v bool) *ListComputeClustersResponseBody {
	s.Success = &v
	return s
}

func (s *ListComputeClustersResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComputeClustersResponseBodyPageResult struct {
	// The paged list of clusters.
	ClusterList []*ListComputeClustersResponseBodyPageResultClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListComputeClustersResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListComputeClustersResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListComputeClustersResponseBodyPageResult) GetClusterList() []*ListComputeClustersResponseBodyPageResultClusterList {
	return s.ClusterList
}

func (s *ListComputeClustersResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListComputeClustersResponseBodyPageResult) SetClusterList(v []*ListComputeClustersResponseBodyPageResultClusterList) *ListComputeClustersResponseBodyPageResult {
	s.ClusterList = v
	return s
}

func (s *ListComputeClustersResponseBodyPageResult) SetTotalCount(v int32) *ListComputeClustersResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListComputeClustersResponseBodyPageResult) Validate() error {
	if s.ClusterList != nil {
		for _, item := range s.ClusterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComputeClustersResponseBodyPageResultClusterList struct {
	// The time when the cluster was created.
	//
	// example:
	//
	// 2025-06-30 08:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator.
	//
	// example:
	//
	// 30012211
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The creator.
	//
	// example:
	//
	// John Doe
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// test
	Des *string `json:"Des,omitempty" xml:"Des,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// 102311
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The user who last modified the cluster.
	//
	// example:
	//
	// 30012211
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The user who last modified the cluster.
	//
	// example:
	//
	// John Doe
	ModifierName *string `json:"ModifierName,omitempty" xml:"ModifierName,omitempty"`
	// The time when the cluster was last updated.
	//
	// example:
	//
	// 2025-06-30 08:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cluster_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cluster version.
	//
	// example:
	//
	// CDH6
	TypeVersion *string `json:"TypeVersion,omitempty" xml:"TypeVersion,omitempty"`
}

func (s ListComputeClustersResponseBodyPageResultClusterList) String() string {
	return dara.Prettify(s)
}

func (s ListComputeClustersResponseBodyPageResultClusterList) GoString() string {
	return s.String()
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) GetCreator() *string {
	return s.Creator
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) GetDes() *string {
	return s.Des
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) GetId() *int64 {
	return s.Id
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) GetModifier() *string {
	return s.Modifier
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) GetModifierName() *string {
	return s.ModifierName
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) GetName() *string {
	return s.Name
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) GetTypeVersion() *string {
	return s.TypeVersion
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) SetCreateTime(v string) *ListComputeClustersResponseBodyPageResultClusterList {
	s.CreateTime = &v
	return s
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) SetCreator(v string) *ListComputeClustersResponseBodyPageResultClusterList {
	s.Creator = &v
	return s
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) SetCreatorName(v string) *ListComputeClustersResponseBodyPageResultClusterList {
	s.CreatorName = &v
	return s
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) SetDes(v string) *ListComputeClustersResponseBodyPageResultClusterList {
	s.Des = &v
	return s
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) SetId(v int64) *ListComputeClustersResponseBodyPageResultClusterList {
	s.Id = &v
	return s
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) SetModifier(v string) *ListComputeClustersResponseBodyPageResultClusterList {
	s.Modifier = &v
	return s
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) SetModifierName(v string) *ListComputeClustersResponseBodyPageResultClusterList {
	s.ModifierName = &v
	return s
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) SetModifyTime(v string) *ListComputeClustersResponseBodyPageResultClusterList {
	s.ModifyTime = &v
	return s
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) SetName(v string) *ListComputeClustersResponseBodyPageResultClusterList {
	s.Name = &v
	return s
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) SetTypeVersion(v string) *ListComputeClustersResponseBodyPageResultClusterList {
	s.TypeVersion = &v
	return s
}

func (s *ListComputeClustersResponseBodyPageResultClusterList) Validate() error {
	return dara.Validate(s)
}
