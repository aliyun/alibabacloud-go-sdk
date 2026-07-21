// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasets(v []*ListDatasetsResponseBodyDatasets) *ListDatasetsResponseBody
	GetDatasets() []*ListDatasetsResponseBodyDatasets
	SetMaxResults(v int32) *ListDatasetsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDatasetsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDatasetsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListDatasetsResponseBody
	GetTotal() *int32
}

type ListDatasetsResponseBody struct {
	// The result set.
	Datasets []*ListDatasetsResponseBodyDatasets `json:"datasets,omitempty" xml:"datasets,omitempty" type:"Repeated"`
	// The maximum number of results specified in this request.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page of results.
	//
	// If the total number of results exceeds the maxResults limit, the data is truncated. You can use nextToken to query the next page of data.
	//
	// example:
	//
	// umaQfI7x758Ns4TgWrj8yA3fYlnk7dJgsfhMrSViRY8=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90F52F93-8800-5A71-8737-18F34BA90166
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 33
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListDatasetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBody) GetDatasets() []*ListDatasetsResponseBodyDatasets {
	return s.Datasets
}

func (s *ListDatasetsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDatasetsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatasetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasetsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListDatasetsResponseBody) SetDatasets(v []*ListDatasetsResponseBodyDatasets) *ListDatasetsResponseBody {
	s.Datasets = v
	return s
}

func (s *ListDatasetsResponseBody) SetMaxResults(v int32) *ListDatasetsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDatasetsResponseBody) SetNextToken(v string) *ListDatasetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDatasetsResponseBody) SetRequestId(v string) *ListDatasetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetsResponseBody) SetTotal(v int32) *ListDatasetsResponseBody {
	s.Total = &v
	return s
}

func (s *ListDatasetsResponseBody) Validate() error {
	if s.Datasets != nil {
		for _, item := range s.Datasets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDatasetsResponseBodyDatasets struct {
	// The agent space name.
	//
	// example:
	//
	// sop-agent
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The creation time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-19T02:11:02Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The dataset name.
	//
	// example:
	//
	// product_faq_dataset
	DatasetName *string `json:"datasetName,omitempty" xml:"datasetName,omitempty"`
	// The dataset description.
	//
	// example:
	//
	// Product FAQ dataset for semantic search
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	IsFavorite  *bool   `json:"isFavorite,omitempty" xml:"isFavorite,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The update time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-05-18T02:21:32Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListDatasetsResponseBodyDatasets) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyDatasets) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyDatasets) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *ListDatasetsResponseBodyDatasets) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDatasetsResponseBodyDatasets) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ListDatasetsResponseBodyDatasets) GetDescription() *string {
	return s.Description
}

func (s *ListDatasetsResponseBodyDatasets) GetIsFavorite() *bool {
	return s.IsFavorite
}

func (s *ListDatasetsResponseBodyDatasets) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDatasetsResponseBodyDatasets) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListDatasetsResponseBodyDatasets) SetAgentSpace(v string) *ListDatasetsResponseBodyDatasets {
	s.AgentSpace = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetCreateTime(v string) *ListDatasetsResponseBodyDatasets {
	s.CreateTime = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetDatasetName(v string) *ListDatasetsResponseBodyDatasets {
	s.DatasetName = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetDescription(v string) *ListDatasetsResponseBodyDatasets {
	s.Description = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetIsFavorite(v bool) *ListDatasetsResponseBodyDatasets {
	s.IsFavorite = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetRegionId(v string) *ListDatasetsResponseBodyDatasets {
	s.RegionId = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) SetUpdateTime(v string) *ListDatasetsResponseBodyDatasets {
	s.UpdateTime = &v
	return s
}

func (s *ListDatasetsResponseBodyDatasets) Validate() error {
	return dara.Validate(s)
}
