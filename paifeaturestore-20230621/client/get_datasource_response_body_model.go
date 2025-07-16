// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *GetDatasourceResponseBody
	GetConfig() *string
	SetDatasourceId(v string) *GetDatasourceResponseBody
	GetDatasourceId() *string
	SetGmtCreateTime(v string) *GetDatasourceResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetDatasourceResponseBody
	GetGmtModifiedTime() *string
	SetName(v string) *GetDatasourceResponseBody
	GetName() *string
	SetRequestId(v string) *GetDatasourceResponseBody
	GetRequestId() *string
	SetType(v string) *GetDatasourceResponseBody
	GetType() *string
	SetUri(v string) *GetDatasourceResponseBody
	GetUri() *string
	SetWorkspaceId(v string) *GetDatasourceResponseBody
	GetWorkspaceId() *string
}

type GetDatasourceResponseBody struct {
	// example:
	//
	// {"address": ""}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// datasource1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// AD7D9E95-BD31-53F2-B710-6C01866FCB05
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// igraph_instance1
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// 32244
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasourceResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetDatasourceResponseBody) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *GetDatasourceResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetDatasourceResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetDatasourceResponseBody) GetName() *string {
	return s.Name
}

func (s *GetDatasourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasourceResponseBody) GetType() *string {
	return s.Type
}

func (s *GetDatasourceResponseBody) GetUri() *string {
	return s.Uri
}

func (s *GetDatasourceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDatasourceResponseBody) SetConfig(v string) *GetDatasourceResponseBody {
	s.Config = &v
	return s
}

func (s *GetDatasourceResponseBody) SetDatasourceId(v string) *GetDatasourceResponseBody {
	s.DatasourceId = &v
	return s
}

func (s *GetDatasourceResponseBody) SetGmtCreateTime(v string) *GetDatasourceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetDatasourceResponseBody) SetGmtModifiedTime(v string) *GetDatasourceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetDatasourceResponseBody) SetName(v string) *GetDatasourceResponseBody {
	s.Name = &v
	return s
}

func (s *GetDatasourceResponseBody) SetRequestId(v string) *GetDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasourceResponseBody) SetType(v string) *GetDatasourceResponseBody {
	s.Type = &v
	return s
}

func (s *GetDatasourceResponseBody) SetUri(v string) *GetDatasourceResponseBody {
	s.Uri = &v
	return s
}

func (s *GetDatasourceResponseBody) SetWorkspaceId(v string) *GetDatasourceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetDatasourceResponseBody) Validate() error {
	return dara.Validate(s)
}
