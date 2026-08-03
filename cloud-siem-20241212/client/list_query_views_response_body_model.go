// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryViewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogProjectName(v string) *ListQueryViewsResponseBody
	GetLogProjectName() *string
	SetLogRegionId(v string) *ListQueryViewsResponseBody
	GetLogRegionId() *string
	SetLogStoreName(v string) *ListQueryViewsResponseBody
	GetLogStoreName() *string
	SetMaxResults(v int32) *ListQueryViewsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListQueryViewsResponseBody
	GetNextToken() *string
	SetQueryViews(v []*ListQueryViewsResponseBodyQueryViews) *ListQueryViewsResponseBody
	GetQueryViews() []*ListQueryViewsResponseBodyQueryViews
	SetRequestId(v string) *ListQueryViewsResponseBody
	GetRequestId() *string
}

type ListQueryViewsResponseBody struct {
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	LogRegionId    *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	LogStoreName   *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The maximum number of results to return when you use the NextToken-based pagination method. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request or if no more results exist. If more results exist, set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// AAAAASLVeIxed4466E0LVmGkzwS6hJKd9DGVGMDRM6Lu****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of query views.
	QueryViews []*ListQueryViewsResponseBodyQueryViews `json:"QueryViews,omitempty" xml:"QueryViews,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQueryViewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQueryViewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueryViewsResponseBody) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *ListQueryViewsResponseBody) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *ListQueryViewsResponseBody) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *ListQueryViewsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQueryViewsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQueryViewsResponseBody) GetQueryViews() []*ListQueryViewsResponseBodyQueryViews {
	return s.QueryViews
}

func (s *ListQueryViewsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQueryViewsResponseBody) SetLogProjectName(v string) *ListQueryViewsResponseBody {
	s.LogProjectName = &v
	return s
}

func (s *ListQueryViewsResponseBody) SetLogRegionId(v string) *ListQueryViewsResponseBody {
	s.LogRegionId = &v
	return s
}

func (s *ListQueryViewsResponseBody) SetLogStoreName(v string) *ListQueryViewsResponseBody {
	s.LogStoreName = &v
	return s
}

func (s *ListQueryViewsResponseBody) SetMaxResults(v int32) *ListQueryViewsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListQueryViewsResponseBody) SetNextToken(v string) *ListQueryViewsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQueryViewsResponseBody) SetQueryViews(v []*ListQueryViewsResponseBodyQueryViews) *ListQueryViewsResponseBody {
	s.QueryViews = v
	return s
}

func (s *ListQueryViewsResponseBody) SetRequestId(v string) *ListQueryViewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueryViewsResponseBody) Validate() error {
	if s.QueryViews != nil {
		for _, item := range s.QueryViews {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQueryViewsResponseBodyQueryViews struct {
	// The custom query condition of the view.
	//
	// example:
	//
	// preset
	QueryViewCondition *string `json:"QueryViewCondition,omitempty" xml:"QueryViewCondition,omitempty"`
	// The alert filter statement of the view.
	//
	// example:
	//
	// module_code: alibaba_cloud_sas_custom_rule
	QueryViewCriteria *string `json:"QueryViewCriteria,omitempty" xml:"QueryViewCriteria,omitempty"`
	// The list of displayed fields.
	//
	// example:
	//
	// ["alert_name_cn","alert_type_cn","entity_list","alert_level","alert_status","start_time","end_time"]
	QueryViewFields *string `json:"QueryViewFields,omitempty" xml:"QueryViewFields,omitempty"`
	// The unique identifier of the query view.
	//
	// example:
	//
	// qv-a1b2c3d4e5f6g7h8****
	QueryViewId *string `json:"QueryViewId,omitempty" xml:"QueryViewId,omitempty"`
	// The view name.
	//
	// example:
	//
	// alert
	QueryViewName *string `json:"QueryViewName,omitempty" xml:"QueryViewName,omitempty"`
	// The display order.
	//
	// example:
	//
	// 1
	QueryViewOrder *string `json:"QueryViewOrder,omitempty" xml:"QueryViewOrder,omitempty"`
	// The scene to which the view belongs.
	//
	// example:
	//
	// Alert
	QueryViewScene *string `json:"QueryViewScene,omitempty" xml:"QueryViewScene,omitempty"`
	// The view status.
	//
	// example:
	//
	// 2
	QueryViewStatus *string `json:"QueryViewStatus,omitempty" xml:"QueryViewStatus,omitempty"`
	// The view type.
	//
	// example:
	//
	// preset
	QueryViewType *string `json:"QueryViewType,omitempty" xml:"QueryViewType,omitempty"`
}

func (s ListQueryViewsResponseBodyQueryViews) String() string {
	return dara.Prettify(s)
}

func (s ListQueryViewsResponseBodyQueryViews) GoString() string {
	return s.String()
}

func (s *ListQueryViewsResponseBodyQueryViews) GetQueryViewCondition() *string {
	return s.QueryViewCondition
}

func (s *ListQueryViewsResponseBodyQueryViews) GetQueryViewCriteria() *string {
	return s.QueryViewCriteria
}

func (s *ListQueryViewsResponseBodyQueryViews) GetQueryViewFields() *string {
	return s.QueryViewFields
}

func (s *ListQueryViewsResponseBodyQueryViews) GetQueryViewId() *string {
	return s.QueryViewId
}

func (s *ListQueryViewsResponseBodyQueryViews) GetQueryViewName() *string {
	return s.QueryViewName
}

func (s *ListQueryViewsResponseBodyQueryViews) GetQueryViewOrder() *string {
	return s.QueryViewOrder
}

func (s *ListQueryViewsResponseBodyQueryViews) GetQueryViewScene() *string {
	return s.QueryViewScene
}

func (s *ListQueryViewsResponseBodyQueryViews) GetQueryViewStatus() *string {
	return s.QueryViewStatus
}

func (s *ListQueryViewsResponseBodyQueryViews) GetQueryViewType() *string {
	return s.QueryViewType
}

func (s *ListQueryViewsResponseBodyQueryViews) SetQueryViewCondition(v string) *ListQueryViewsResponseBodyQueryViews {
	s.QueryViewCondition = &v
	return s
}

func (s *ListQueryViewsResponseBodyQueryViews) SetQueryViewCriteria(v string) *ListQueryViewsResponseBodyQueryViews {
	s.QueryViewCriteria = &v
	return s
}

func (s *ListQueryViewsResponseBodyQueryViews) SetQueryViewFields(v string) *ListQueryViewsResponseBodyQueryViews {
	s.QueryViewFields = &v
	return s
}

func (s *ListQueryViewsResponseBodyQueryViews) SetQueryViewId(v string) *ListQueryViewsResponseBodyQueryViews {
	s.QueryViewId = &v
	return s
}

func (s *ListQueryViewsResponseBodyQueryViews) SetQueryViewName(v string) *ListQueryViewsResponseBodyQueryViews {
	s.QueryViewName = &v
	return s
}

func (s *ListQueryViewsResponseBodyQueryViews) SetQueryViewOrder(v string) *ListQueryViewsResponseBodyQueryViews {
	s.QueryViewOrder = &v
	return s
}

func (s *ListQueryViewsResponseBodyQueryViews) SetQueryViewScene(v string) *ListQueryViewsResponseBodyQueryViews {
	s.QueryViewScene = &v
	return s
}

func (s *ListQueryViewsResponseBodyQueryViews) SetQueryViewStatus(v string) *ListQueryViewsResponseBodyQueryViews {
	s.QueryViewStatus = &v
	return s
}

func (s *ListQueryViewsResponseBodyQueryViews) SetQueryViewType(v string) *ListQueryViewsResponseBodyQueryViews {
	s.QueryViewType = &v
	return s
}

func (s *ListQueryViewsResponseBodyQueryViews) Validate() error {
	return dara.Validate(s)
}
