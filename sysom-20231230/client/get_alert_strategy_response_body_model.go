// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAlertStrategyResponseBody
	GetCode() *string
	SetData(v *GetAlertStrategyResponseBodyData) *GetAlertStrategyResponseBody
	GetData() *GetAlertStrategyResponseBodyData
	SetMessage(v string) *GetAlertStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAlertStrategyResponseBody
	GetRequestId() *string
}

type GetAlertStrategyResponseBody struct {
	// Status code:
	//
	// - `code == Success` indicates successful authorization;
	//
	// - Other status codes indicate failed authorization. When authorization fails, view the `message` field to obtain detailed error message.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned data.
	Data *GetAlertStrategyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Error message
	//
	// - If `code == Success`, this field is empty;
	//
	// - Otherwise, this field contains the request error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAlertStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlertStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlertStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAlertStrategyResponseBody) GetData() *GetAlertStrategyResponseBodyData {
	return s.Data
}

func (s *GetAlertStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAlertStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlertStrategyResponseBody) SetCode(v string) *GetAlertStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlertStrategyResponseBody) SetData(v *GetAlertStrategyResponseBodyData) *GetAlertStrategyResponseBody {
	s.Data = v
	return s
}

func (s *GetAlertStrategyResponseBody) SetMessage(v string) *GetAlertStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlertStrategyResponseBody) SetRequestId(v string) *GetAlertStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlertStrategyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertStrategyResponseBodyData struct {
	// Creation Time.
	//
	// example:
	//
	// 1751520976660
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Indicates whether the alert policy is enabled
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Alert policy ID
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// k8s label
	K8sLabel *bool `json:"k8sLabel,omitempty" xml:"k8sLabel,omitempty"`
	// Policy Name
	//
	// example:
	//
	// strategy1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Details of the alert policy
	Strategy *GetAlertStrategyResponseBodyDataStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
	// User ID
	//
	// example:
	//
	// 1222933234714935
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// Update Time.
	//
	// example:
	//
	// 1751254826285
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetAlertStrategyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAlertStrategyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAlertStrategyResponseBodyData) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetAlertStrategyResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetAlertStrategyResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetAlertStrategyResponseBodyData) GetK8sLabel() *bool {
	return s.K8sLabel
}

func (s *GetAlertStrategyResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetAlertStrategyResponseBodyData) GetStrategy() *GetAlertStrategyResponseBodyDataStrategy {
	return s.Strategy
}

func (s *GetAlertStrategyResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *GetAlertStrategyResponseBodyData) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *GetAlertStrategyResponseBodyData) SetCreatedAt(v int64) *GetAlertStrategyResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetAlertStrategyResponseBodyData) SetEnabled(v bool) *GetAlertStrategyResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetAlertStrategyResponseBodyData) SetId(v int64) *GetAlertStrategyResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAlertStrategyResponseBodyData) SetK8sLabel(v bool) *GetAlertStrategyResponseBodyData {
	s.K8sLabel = &v
	return s
}

func (s *GetAlertStrategyResponseBodyData) SetName(v string) *GetAlertStrategyResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAlertStrategyResponseBodyData) SetStrategy(v *GetAlertStrategyResponseBodyDataStrategy) *GetAlertStrategyResponseBodyData {
	s.Strategy = v
	return s
}

func (s *GetAlertStrategyResponseBodyData) SetUid(v string) *GetAlertStrategyResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetAlertStrategyResponseBodyData) SetUpdatedAt(v int64) *GetAlertStrategyResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetAlertStrategyResponseBodyData) Validate() error {
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertStrategyResponseBodyDataStrategy struct {
	// Set of clusters that accept alerts
	Clusters     []*string   `json:"clusters,omitempty" xml:"clusters,omitempty" type:"Repeated"`
	Destinations interface{} `json:"destinations,omitempty" xml:"destinations,omitempty"`
	// List of abnormal items that accept alerts
	//
	// example:
	//
	// 节点CPU使用率检测
	Items interface{} `json:"items,omitempty" xml:"items,omitempty"`
}

func (s GetAlertStrategyResponseBodyDataStrategy) String() string {
	return dara.Prettify(s)
}

func (s GetAlertStrategyResponseBodyDataStrategy) GoString() string {
	return s.String()
}

func (s *GetAlertStrategyResponseBodyDataStrategy) GetClusters() []*string {
	return s.Clusters
}

func (s *GetAlertStrategyResponseBodyDataStrategy) GetDestinations() interface{} {
	return s.Destinations
}

func (s *GetAlertStrategyResponseBodyDataStrategy) GetItems() interface{} {
	return s.Items
}

func (s *GetAlertStrategyResponseBodyDataStrategy) SetClusters(v []*string) *GetAlertStrategyResponseBodyDataStrategy {
	s.Clusters = v
	return s
}

func (s *GetAlertStrategyResponseBodyDataStrategy) SetDestinations(v interface{}) *GetAlertStrategyResponseBodyDataStrategy {
	s.Destinations = v
	return s
}

func (s *GetAlertStrategyResponseBodyDataStrategy) SetItems(v interface{}) *GetAlertStrategyResponseBodyDataStrategy {
	s.Items = v
	return s
}

func (s *GetAlertStrategyResponseBodyDataStrategy) Validate() error {
	return dara.Validate(s)
}
