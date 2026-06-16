// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *GetDatasetResponseBody
	GetAgentSpace() *string
	SetCreateTime(v string) *GetDatasetResponseBody
	GetCreateTime() *string
	SetDatasetName(v string) *GetDatasetResponseBody
	GetDatasetName() *string
	SetDescription(v string) *GetDatasetResponseBody
	GetDescription() *string
	SetRegionId(v string) *GetDatasetResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetDatasetResponseBody
	GetRequestId() *string
	SetSchema(v map[string]*IndexKey) *GetDatasetResponseBody
	GetSchema() map[string]*IndexKey
	SetUpdateTime(v string) *GetDatasetResponseBody
	GetUpdateTime() *string
}

type GetDatasetResponseBody struct {
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	CreateTime  *string              `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DatasetName *string              `json:"datasetName,omitempty" xml:"datasetName,omitempty"`
	Description *string              `json:"description,omitempty" xml:"description,omitempty"`
	RegionId    *string              `json:"regionId,omitempty" xml:"regionId,omitempty"`
	RequestId   *string              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Schema      map[string]*IndexKey `json:"schema,omitempty" xml:"schema,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBody) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetDatasetResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDatasetResponseBody) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetDatasetResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetDatasetResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetResponseBody) GetSchema() map[string]*IndexKey {
	return s.Schema
}

func (s *GetDatasetResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetDatasetResponseBody) SetAgentSpace(v string) *GetDatasetResponseBody {
	s.AgentSpace = &v
	return s
}

func (s *GetDatasetResponseBody) SetCreateTime(v string) *GetDatasetResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetDatasetResponseBody) SetDatasetName(v string) *GetDatasetResponseBody {
	s.DatasetName = &v
	return s
}

func (s *GetDatasetResponseBody) SetDescription(v string) *GetDatasetResponseBody {
	s.Description = &v
	return s
}

func (s *GetDatasetResponseBody) SetRegionId(v string) *GetDatasetResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetDatasetResponseBody) SetRequestId(v string) *GetDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetResponseBody) SetSchema(v map[string]*IndexKey) *GetDatasetResponseBody {
	s.Schema = v
	return s
}

func (s *GetDatasetResponseBody) SetUpdateTime(v string) *GetDatasetResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
