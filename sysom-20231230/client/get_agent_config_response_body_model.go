// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentConfigResponseBody
	GetCode() *string
	SetData(v []*GetAgentConfigResponseBodyData) *GetAgentConfigResponseBody
	GetData() []*GetAgentConfigResponseBodyData
	SetMessage(v string) *GetAgentConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentConfigResponseBody
	GetRequestId() *string
	SetTotal(v int64) *GetAgentConfigResponseBody
	GetTotal() *int64
}

type GetAgentConfigResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data []*GetAgentConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the error message of the request.
	//
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of returned records.
	//
	// example:
	//
	// 10
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetAgentConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentConfigResponseBody) GetData() []*GetAgentConfigResponseBodyData {
	return s.Data
}

func (s *GetAgentConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentConfigResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetAgentConfigResponseBody) SetCode(v string) *GetAgentConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentConfigResponseBody) SetData(v []*GetAgentConfigResponseBodyData) *GetAgentConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentConfigResponseBody) SetMessage(v string) *GetAgentConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentConfigResponseBody) SetRequestId(v string) *GetAgentConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentConfigResponseBody) SetTotal(v int64) *GetAgentConfigResponseBody {
	s.Total = &v
	return s
}

func (s *GetAgentConfigResponseBody) Validate() error {
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

type GetAgentConfigResponseBodyData struct {
	// The ID of the component configuration.
	//
	// example:
	//
	// 4d5ecf02-3d01-11f0-8c90-00163e1029af
	ConfigId *string `json:"configId,omitempty" xml:"configId,omitempty"`
	// The name of the configuration.
	//
	// example:
	//
	// Configuration-x
	ConfigName *string `json:"configName,omitempty" xml:"configName,omitempty"`
	// The type of the configuration.
	//
	// example:
	//
	// Custom
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// The text of the chat content.
	//
	// example:
	//
	// The specific content of the copilot reply
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-07-10T17:57:10
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The information about the resources to which the configuration is applied.
	Resource *GetAgentConfigResponseBodyDataResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Struct"`
	// The update time.
	//
	// example:
	//
	// 2026-07-10T17:57:10
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetAgentConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentConfigResponseBodyData) GetConfigId() *string {
	return s.ConfigId
}

func (s *GetAgentConfigResponseBodyData) GetConfigName() *string {
	return s.ConfigName
}

func (s *GetAgentConfigResponseBodyData) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetAgentConfigResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetAgentConfigResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetAgentConfigResponseBodyData) GetResource() *GetAgentConfigResponseBodyDataResource {
	return s.Resource
}

func (s *GetAgentConfigResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetAgentConfigResponseBodyData) SetConfigId(v string) *GetAgentConfigResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *GetAgentConfigResponseBodyData) SetConfigName(v string) *GetAgentConfigResponseBodyData {
	s.ConfigName = &v
	return s
}

func (s *GetAgentConfigResponseBodyData) SetConfigType(v string) *GetAgentConfigResponseBodyData {
	s.ConfigType = &v
	return s
}

func (s *GetAgentConfigResponseBodyData) SetContent(v string) *GetAgentConfigResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetAgentConfigResponseBodyData) SetCreatedAt(v string) *GetAgentConfigResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetAgentConfigResponseBodyData) SetResource(v *GetAgentConfigResponseBodyDataResource) *GetAgentConfigResponseBodyData {
	s.Resource = v
	return s
}

func (s *GetAgentConfigResponseBodyData) SetUpdatedAt(v string) *GetAgentConfigResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetAgentConfigResponseBodyData) Validate() error {
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentConfigResponseBodyDataResource struct {
	// The number of clusters to which the configuration is applied.
	//
	// example:
	//
	// 0
	Cluster *int32 `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The number of nodes to which the configuration is applied.
	//
	// example:
	//
	// 0
	Node *int32 `json:"node,omitempty" xml:"node,omitempty"`
}

func (s GetAgentConfigResponseBodyDataResource) String() string {
	return dara.Prettify(s)
}

func (s GetAgentConfigResponseBodyDataResource) GoString() string {
	return s.String()
}

func (s *GetAgentConfigResponseBodyDataResource) GetCluster() *int32 {
	return s.Cluster
}

func (s *GetAgentConfigResponseBodyDataResource) GetNode() *int32 {
	return s.Node
}

func (s *GetAgentConfigResponseBodyDataResource) SetCluster(v int32) *GetAgentConfigResponseBodyDataResource {
	s.Cluster = &v
	return s
}

func (s *GetAgentConfigResponseBodyDataResource) SetNode(v int32) *GetAgentConfigResponseBodyDataResource {
	s.Node = &v
	return s
}

func (s *GetAgentConfigResponseBodyDataResource) Validate() error {
	return dara.Validate(s)
}
