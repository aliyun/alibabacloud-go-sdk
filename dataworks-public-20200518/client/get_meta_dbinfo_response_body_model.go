// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaDBInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaDBInfoResponseBodyData) *GetMetaDBInfoResponseBody
	GetData() *GetMetaDBInfoResponseBodyData
	SetRequestId(v string) *GetMetaDBInfoResponseBody
	GetRequestId() *string
}

type GetMetaDBInfoResponseBody struct {
	// The basic metadata information.
	Data *GetMetaDBInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0bc1411515937****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMetaDBInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaDBInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaDBInfoResponseBody) GetData() *GetMetaDBInfoResponseBodyData {
	return s.Data
}

func (s *GetMetaDBInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaDBInfoResponseBody) SetData(v *GetMetaDBInfoResponseBodyData) *GetMetaDBInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaDBInfoResponseBody) SetRequestId(v string) *GetMetaDBInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaDBInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaDBInfoResponseBodyData struct {
	// The compute engine instance ID. Specify the ID in the `Engine type.Engine name` format.
	//
	// example:
	//
	// odps.engine_name
	AppGuid *string `json:"AppGuid,omitempty" xml:"AppGuid,omitempty"`
	// The EMR cluster ID.
	//
	// example:
	//
	// abc
	ClusterBizId *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	// The comment.
	//
	// example:
	//
	// The ID of the compute engine instance. The ID is in the Engine type.Engine name format.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the compute engine instance was created.
	//
	// example:
	//
	// 1541576644000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The endpoint of the service.
	//
	// example:
	//
	// http://service.odpsstg.aliyun-inc.com/
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The type of the environment. Valid values: 0 and 1. The value 0 indicates the development environment. The value 1 indicates the production environment.
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The storage path of the metadatabase of the EMR cluster.
	//
	// example:
	//
	// hdfs://
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The time when the compute engine instance was modified.
	//
	// example:
	//
	// 1541576644000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the Alibaba Cloud account used by the workspace owner.
	//
	// example:
	//
	// 23
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the workspace owner.
	//
	// example:
	//
	// 323
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 22
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The display name of the workspace.
	//
	// example:
	//
	// The storage path of the metadatabase of the EMR cluster.
	ProjectNameCn *string `json:"ProjectNameCn,omitempty" xml:"ProjectNameCn,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 233
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The type of the metadatabase.
	//
	// example:
	//
	// hive
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMetaDBInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaDBInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaDBInfoResponseBodyData) GetAppGuid() *string {
	return s.AppGuid
}

func (s *GetMetaDBInfoResponseBodyData) GetClusterBizId() *string {
	return s.ClusterBizId
}

func (s *GetMetaDBInfoResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *GetMetaDBInfoResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetMetaDBInfoResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetMetaDBInfoResponseBodyData) GetEnvType() *int32 {
	return s.EnvType
}

func (s *GetMetaDBInfoResponseBodyData) GetLocation() *string {
	return s.Location
}

func (s *GetMetaDBInfoResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetMetaDBInfoResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMetaDBInfoResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetMetaDBInfoResponseBodyData) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetMetaDBInfoResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetMetaDBInfoResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetMetaDBInfoResponseBodyData) GetProjectNameCn() *string {
	return s.ProjectNameCn
}

func (s *GetMetaDBInfoResponseBodyData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetMetaDBInfoResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetMetaDBInfoResponseBodyData) SetAppGuid(v string) *GetMetaDBInfoResponseBodyData {
	s.AppGuid = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetClusterBizId(v string) *GetMetaDBInfoResponseBodyData {
	s.ClusterBizId = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetComment(v string) *GetMetaDBInfoResponseBodyData {
	s.Comment = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetCreateTime(v int64) *GetMetaDBInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetEndpoint(v string) *GetMetaDBInfoResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetEnvType(v int32) *GetMetaDBInfoResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetLocation(v string) *GetMetaDBInfoResponseBodyData {
	s.Location = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetModifyTime(v int64) *GetMetaDBInfoResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetName(v string) *GetMetaDBInfoResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetOwnerId(v string) *GetMetaDBInfoResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetOwnerName(v string) *GetMetaDBInfoResponseBodyData {
	s.OwnerName = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetProjectId(v int64) *GetMetaDBInfoResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetProjectName(v string) *GetMetaDBInfoResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetProjectNameCn(v string) *GetMetaDBInfoResponseBodyData {
	s.ProjectNameCn = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetTenantId(v int64) *GetMetaDBInfoResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) SetType(v string) *GetMetaDBInfoResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMetaDBInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
