// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumnFields(v string) *GetConnectorResponseBody
	GetColumnFields() *string
	SetConfig(v string) *GetConnectorResponseBody
	GetConfig() *string
	SetConnectorId(v string) *GetConnectorResponseBody
	GetConnectorId() *string
	SetCreateTime(v string) *GetConnectorResponseBody
	GetCreateTime() *string
	SetCreator(v string) *GetConnectorResponseBody
	GetCreator() *string
	SetDoneTime(v string) *GetConnectorResponseBody
	GetDoneTime() *string
	SetProjectName(v string) *GetConnectorResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetConnectorResponseBody
	GetRequestId() *string
	SetState(v string) *GetConnectorResponseBody
	GetState() *string
	SetSubscriptionId(v string) *GetConnectorResponseBody
	GetSubscriptionId() *string
	SetSuccess(v bool) *GetConnectorResponseBody
	GetSuccess() *bool
	SetTopicName(v string) *GetConnectorResponseBody
	GetTopicName() *string
	SetType(v string) *GetConnectorResponseBody
	GetType() *string
	SetUpdateTime(v string) *GetConnectorResponseBody
	GetUpdateTime() *string
}

type GetConnectorResponseBody struct {
	// The list of fields for the synchronization task.
	//
	// example:
	//
	// [\\"field1\\",\\"field2\\"]
	ColumnFields *string `json:"ColumnFields,omitempty" xml:"ColumnFields,omitempty"`
	// The configuration information of the synchronization task.
	//
	// example:
	//
	// {\\"TimestampUnit\\":\\"MICROSECOND\\",\\"PartitionConfig\\":{\\"hh\\":\\"%H\\",\\"mm\\":\\"%M\\",\\"ds\\":\\"%Y%m%d\\"},\\"Project\\":\\"xxx\\",\\"TimeRange\\":15,\\"TimeZone\\":\\"Asia/Shanghai\\",\\"Table\\":\\"xxx\\",\\"OdpsEndpoint\\":\\"xxx\\",\\"AccessId\\":\\"xxx\\",\\"PartitionMode\\":\\"SYSTEM_TIME\\",\\"AuthMode\\":\\"ak\\"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The synchronization task ID.
	//
	// example:
	//
	// c5e07a96-5069-4486-87c3-0d281951f772
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1724041098000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator of the synchronization task.
	//
	// example:
	//
	// 270523390948438349
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The time when the synchronization task is marked as done. This parameter takes effect only for ODPS synchronization (SINK_ODPS) tasks.
	//
	// example:
	//
	// 2025-06-06 15:45:00
	DoneTime *string `json:"DoneTime,omitempty" xml:"DoneTime,omitempty"`
	// The project name.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A20A7093-8FE0-058C-BE0C-3C8057D5F1A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the synchronization task.
	//
	// example:
	//
	// RUNNING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The subscription ID.
	//
	// example:
	//
	// 1764123132492KO88A
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The topic name.
	//
	// example:
	//
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The type of the synchronization task.
	//
	// example:
	//
	// SINK_ODPS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The last update time.
	//
	// example:
	//
	// 1724041098000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectorResponseBody) GetColumnFields() *string {
	return s.ColumnFields
}

func (s *GetConnectorResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetConnectorResponseBody) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *GetConnectorResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetConnectorResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetConnectorResponseBody) GetDoneTime() *string {
	return s.DoneTime
}

func (s *GetConnectorResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConnectorResponseBody) GetState() *string {
	return s.State
}

func (s *GetConnectorResponseBody) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *GetConnectorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConnectorResponseBody) GetTopicName() *string {
	return s.TopicName
}

func (s *GetConnectorResponseBody) GetType() *string {
	return s.Type
}

func (s *GetConnectorResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetConnectorResponseBody) SetColumnFields(v string) *GetConnectorResponseBody {
	s.ColumnFields = &v
	return s
}

func (s *GetConnectorResponseBody) SetConfig(v string) *GetConnectorResponseBody {
	s.Config = &v
	return s
}

func (s *GetConnectorResponseBody) SetConnectorId(v string) *GetConnectorResponseBody {
	s.ConnectorId = &v
	return s
}

func (s *GetConnectorResponseBody) SetCreateTime(v string) *GetConnectorResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetConnectorResponseBody) SetCreator(v string) *GetConnectorResponseBody {
	s.Creator = &v
	return s
}

func (s *GetConnectorResponseBody) SetDoneTime(v string) *GetConnectorResponseBody {
	s.DoneTime = &v
	return s
}

func (s *GetConnectorResponseBody) SetProjectName(v string) *GetConnectorResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetConnectorResponseBody) SetRequestId(v string) *GetConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectorResponseBody) SetState(v string) *GetConnectorResponseBody {
	s.State = &v
	return s
}

func (s *GetConnectorResponseBody) SetSubscriptionId(v string) *GetConnectorResponseBody {
	s.SubscriptionId = &v
	return s
}

func (s *GetConnectorResponseBody) SetSuccess(v bool) *GetConnectorResponseBody {
	s.Success = &v
	return s
}

func (s *GetConnectorResponseBody) SetTopicName(v string) *GetConnectorResponseBody {
	s.TopicName = &v
	return s
}

func (s *GetConnectorResponseBody) SetType(v string) *GetConnectorResponseBody {
	s.Type = &v
	return s
}

func (s *GetConnectorResponseBody) SetUpdateTime(v string) *GetConnectorResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetConnectorResponseBody) Validate() error {
	return dara.Validate(s)
}
