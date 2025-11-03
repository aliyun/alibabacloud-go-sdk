// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentInstallRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAgentInstallRecordsResponseBody
	GetRequestId() *string
	SetCode(v string) *ListAgentInstallRecordsResponseBody
	GetCode() *string
	SetData(v []*ListAgentInstallRecordsResponseBodyData) *ListAgentInstallRecordsResponseBody
	GetData() []*ListAgentInstallRecordsResponseBodyData
	SetMessage(v string) *ListAgentInstallRecordsResponseBody
	GetMessage() *string
	SetTotal(v int64) *ListAgentInstallRecordsResponseBody
	GetTotal() *int64
}

type ListAgentInstallRecordsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListAgentInstallRecordsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 64
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAgentInstallRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstallRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentInstallRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentInstallRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentInstallRecordsResponseBody) GetData() []*ListAgentInstallRecordsResponseBodyData {
	return s.Data
}

func (s *ListAgentInstallRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentInstallRecordsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAgentInstallRecordsResponseBody) SetRequestId(v string) *ListAgentInstallRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBody) SetCode(v string) *ListAgentInstallRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBody) SetData(v []*ListAgentInstallRecordsResponseBodyData) *ListAgentInstallRecordsResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentInstallRecordsResponseBody) SetMessage(v string) *ListAgentInstallRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBody) SetTotal(v int64) *ListAgentInstallRecordsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBody) Validate() error {
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

type ListAgentInstallRecordsResponseBodyData struct {
	CreatedAt     *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	InstanceId    *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	PluginId      *string `json:"plugin_id,omitempty" xml:"plugin_id,omitempty"`
	PluginVersion *string `json:"plugin_version,omitempty" xml:"plugin_version,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt     *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s ListAgentInstallRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstallRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentInstallRecordsResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListAgentInstallRecordsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentInstallRecordsResponseBodyData) GetPluginId() *string {
	return s.PluginId
}

func (s *ListAgentInstallRecordsResponseBodyData) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *ListAgentInstallRecordsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListAgentInstallRecordsResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListAgentInstallRecordsResponseBodyData) SetCreatedAt(v string) *ListAgentInstallRecordsResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBodyData) SetInstanceId(v string) *ListAgentInstallRecordsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBodyData) SetPluginId(v string) *ListAgentInstallRecordsResponseBodyData {
	s.PluginId = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBodyData) SetPluginVersion(v string) *ListAgentInstallRecordsResponseBodyData {
	s.PluginVersion = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBodyData) SetStatus(v string) *ListAgentInstallRecordsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBodyData) SetUpdatedAt(v string) *ListAgentInstallRecordsResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
