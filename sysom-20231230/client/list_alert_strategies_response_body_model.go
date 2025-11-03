// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertStrategiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAlertStrategiesResponseBody
	GetCode() *string
	SetData(v []*ListAlertStrategiesResponseBodyData) *ListAlertStrategiesResponseBody
	GetData() []*ListAlertStrategiesResponseBodyData
	SetMaxResults(v int32) *ListAlertStrategiesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAlertStrategiesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAlertStrategiesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAlertStrategiesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListAlertStrategiesResponseBody
	GetTotal() *int64
}

type ListAlertStrategiesResponseBody struct {
	// example:
	//
	// Success
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListAlertStrategiesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// c2f78a783f49457caba6bace6f6f79e4
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 92
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAlertStrategiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertStrategiesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAlertStrategiesResponseBody) GetData() []*ListAlertStrategiesResponseBodyData {
	return s.Data
}

func (s *ListAlertStrategiesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlertStrategiesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAlertStrategiesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlertStrategiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertStrategiesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAlertStrategiesResponseBody) SetCode(v string) *ListAlertStrategiesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlertStrategiesResponseBody) SetData(v []*ListAlertStrategiesResponseBodyData) *ListAlertStrategiesResponseBody {
	s.Data = v
	return s
}

func (s *ListAlertStrategiesResponseBody) SetMaxResults(v int32) *ListAlertStrategiesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAlertStrategiesResponseBody) SetMessage(v string) *ListAlertStrategiesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlertStrategiesResponseBody) SetNextToken(v string) *ListAlertStrategiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAlertStrategiesResponseBody) SetRequestId(v string) *ListAlertStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertStrategiesResponseBody) SetTotal(v int64) *ListAlertStrategiesResponseBody {
	s.Total = &v
	return s
}

func (s *ListAlertStrategiesResponseBody) Validate() error {
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

type ListAlertStrategiesResponseBodyData struct {
	// example:
	//
	// 1753669116286
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// 1
	Id       *int64 `json:"id,omitempty" xml:"id,omitempty"`
	K8sLabel *bool  `json:"k8sLabel,omitempty" xml:"k8sLabel,omitempty"`
	// example:
	//
	// strategy1
	Name     *string                                      `json:"name,omitempty" xml:"name,omitempty"`
	Strategy *ListAlertStrategiesResponseBodyDataStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
	// example:
	//
	// 1880327028143673
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// example:
	//
	// 1753237017710
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListAlertStrategiesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAlertStrategiesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlertStrategiesResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListAlertStrategiesResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListAlertStrategiesResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListAlertStrategiesResponseBodyData) GetK8sLabel() *bool {
	return s.K8sLabel
}

func (s *ListAlertStrategiesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAlertStrategiesResponseBodyData) GetStrategy() *ListAlertStrategiesResponseBodyDataStrategy {
	return s.Strategy
}

func (s *ListAlertStrategiesResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *ListAlertStrategiesResponseBodyData) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *ListAlertStrategiesResponseBodyData) SetCreatedAt(v string) *ListAlertStrategiesResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListAlertStrategiesResponseBodyData) SetEnabled(v bool) *ListAlertStrategiesResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *ListAlertStrategiesResponseBodyData) SetId(v int64) *ListAlertStrategiesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAlertStrategiesResponseBodyData) SetK8sLabel(v bool) *ListAlertStrategiesResponseBodyData {
	s.K8sLabel = &v
	return s
}

func (s *ListAlertStrategiesResponseBodyData) SetName(v string) *ListAlertStrategiesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAlertStrategiesResponseBodyData) SetStrategy(v *ListAlertStrategiesResponseBodyDataStrategy) *ListAlertStrategiesResponseBodyData {
	s.Strategy = v
	return s
}

func (s *ListAlertStrategiesResponseBodyData) SetUid(v string) *ListAlertStrategiesResponseBodyData {
	s.Uid = &v
	return s
}

func (s *ListAlertStrategiesResponseBodyData) SetUpdatedAt(v int64) *ListAlertStrategiesResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *ListAlertStrategiesResponseBodyData) Validate() error {
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAlertStrategiesResponseBodyDataStrategy struct {
	Clusters []*string `json:"clusters,omitempty" xml:"clusters,omitempty" type:"Repeated"`
	Items    []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListAlertStrategiesResponseBodyDataStrategy) String() string {
	return dara.Prettify(s)
}

func (s ListAlertStrategiesResponseBodyDataStrategy) GoString() string {
	return s.String()
}

func (s *ListAlertStrategiesResponseBodyDataStrategy) GetClusters() []*string {
	return s.Clusters
}

func (s *ListAlertStrategiesResponseBodyDataStrategy) GetItems() []*string {
	return s.Items
}

func (s *ListAlertStrategiesResponseBodyDataStrategy) SetClusters(v []*string) *ListAlertStrategiesResponseBodyDataStrategy {
	s.Clusters = v
	return s
}

func (s *ListAlertStrategiesResponseBodyDataStrategy) SetItems(v []*string) *ListAlertStrategiesResponseBodyDataStrategy {
	s.Items = v
	return s
}

func (s *ListAlertStrategiesResponseBodyDataStrategy) Validate() error {
	return dara.Validate(s)
}
