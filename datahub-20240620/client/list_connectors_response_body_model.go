// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v *ListConnectorsResponseBodyList) *ListConnectorsResponseBody
	GetList() *ListConnectorsResponseBodyList
	SetMaxResults(v int32) *ListConnectorsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListConnectorsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListConnectorsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListConnectorsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListConnectorsResponseBody
	GetTotalCount() *int32
}

type ListConnectorsResponseBody struct {
	List *ListConnectorsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 9892074a2a89600ae4b0d5a34fb99a3f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// A20A7093-8FE0-058C-BE0C-3C8057D5F1A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListConnectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBody) GetList() *ListConnectorsResponseBodyList {
	return s.List
}

func (s *ListConnectorsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConnectorsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConnectorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConnectorsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListConnectorsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListConnectorsResponseBody) SetList(v *ListConnectorsResponseBodyList) *ListConnectorsResponseBody {
	s.List = v
	return s
}

func (s *ListConnectorsResponseBody) SetMaxResults(v int32) *ListConnectorsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListConnectorsResponseBody) SetNextToken(v string) *ListConnectorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConnectorsResponseBody) SetRequestId(v string) *ListConnectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectorsResponseBody) SetSuccess(v bool) *ListConnectorsResponseBody {
	s.Success = &v
	return s
}

func (s *ListConnectorsResponseBody) SetTotalCount(v int32) *ListConnectorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListConnectorsResponseBody) Validate() error {
	if s.List != nil {
		if err := s.List.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConnectorsResponseBodyList struct {
	Connector []*ListConnectorsResponseBodyListConnector `json:"Connector,omitempty" xml:"Connector,omitempty" type:"Repeated"`
}

func (s ListConnectorsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyList) GetConnector() []*ListConnectorsResponseBodyListConnector {
	return s.Connector
}

func (s *ListConnectorsResponseBodyList) SetConnector(v []*ListConnectorsResponseBodyListConnector) *ListConnectorsResponseBodyList {
	s.Connector = v
	return s
}

func (s *ListConnectorsResponseBodyList) Validate() error {
	if s.Connector != nil {
		for _, item := range s.Connector {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConnectorsResponseBodyListConnector struct {
	// example:
	//
	// [\\"field1\\",\\"field2\\"]
	ColumnFields *string `json:"ColumnFields,omitempty" xml:"ColumnFields,omitempty"`
	// example:
	//
	// {\\"Table\\":\\"r3\\",\\"Endpoint\\":\\"https://Device-data.cn-beijing.ots-internal.aliyuncs.com\\",\\"Instance\\":\\"Device-data\\",\\"WriteMode\\":\\"PUT\\",\\"AuthMode\\":\\"STS\\"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// fa44384c-0ac5-4d3e-8acd-76e18636ab10
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// example:
	//
	// 1724041098000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1696648921408952
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 2025-12-04 16:45:00
	DoneTime *string `json:"DoneTime,omitempty" xml:"DoneTime,omitempty"`
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// RUNNING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1745824636429WZ2EE
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	// example:
	//
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// example:
	//
	// SINK_ODPS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1708171905000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListConnectorsResponseBodyListConnector) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsResponseBodyListConnector) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyListConnector) GetColumnFields() *string {
	return s.ColumnFields
}

func (s *ListConnectorsResponseBodyListConnector) GetConfig() *string {
	return s.Config
}

func (s *ListConnectorsResponseBodyListConnector) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *ListConnectorsResponseBodyListConnector) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListConnectorsResponseBodyListConnector) GetCreator() *string {
	return s.Creator
}

func (s *ListConnectorsResponseBodyListConnector) GetDoneTime() *string {
	return s.DoneTime
}

func (s *ListConnectorsResponseBodyListConnector) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListConnectorsResponseBodyListConnector) GetState() *string {
	return s.State
}

func (s *ListConnectorsResponseBodyListConnector) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *ListConnectorsResponseBodyListConnector) GetTopicName() *string {
	return s.TopicName
}

func (s *ListConnectorsResponseBodyListConnector) GetType() *string {
	return s.Type
}

func (s *ListConnectorsResponseBodyListConnector) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListConnectorsResponseBodyListConnector) SetColumnFields(v string) *ListConnectorsResponseBodyListConnector {
	s.ColumnFields = &v
	return s
}

func (s *ListConnectorsResponseBodyListConnector) SetConfig(v string) *ListConnectorsResponseBodyListConnector {
	s.Config = &v
	return s
}

func (s *ListConnectorsResponseBodyListConnector) SetConnectorId(v string) *ListConnectorsResponseBodyListConnector {
	s.ConnectorId = &v
	return s
}

func (s *ListConnectorsResponseBodyListConnector) SetCreateTime(v string) *ListConnectorsResponseBodyListConnector {
	s.CreateTime = &v
	return s
}

func (s *ListConnectorsResponseBodyListConnector) SetCreator(v string) *ListConnectorsResponseBodyListConnector {
	s.Creator = &v
	return s
}

func (s *ListConnectorsResponseBodyListConnector) SetDoneTime(v string) *ListConnectorsResponseBodyListConnector {
	s.DoneTime = &v
	return s
}

func (s *ListConnectorsResponseBodyListConnector) SetProjectName(v string) *ListConnectorsResponseBodyListConnector {
	s.ProjectName = &v
	return s
}

func (s *ListConnectorsResponseBodyListConnector) SetState(v string) *ListConnectorsResponseBodyListConnector {
	s.State = &v
	return s
}

func (s *ListConnectorsResponseBodyListConnector) SetSubscriptionId(v string) *ListConnectorsResponseBodyListConnector {
	s.SubscriptionId = &v
	return s
}

func (s *ListConnectorsResponseBodyListConnector) SetTopicName(v string) *ListConnectorsResponseBodyListConnector {
	s.TopicName = &v
	return s
}

func (s *ListConnectorsResponseBodyListConnector) SetType(v string) *ListConnectorsResponseBodyListConnector {
	s.Type = &v
	return s
}

func (s *ListConnectorsResponseBodyListConnector) SetUpdateTime(v string) *ListConnectorsResponseBodyListConnector {
	s.UpdateTime = &v
	return s
}

func (s *ListConnectorsResponseBodyListConnector) Validate() error {
	return dara.Validate(s)
}
