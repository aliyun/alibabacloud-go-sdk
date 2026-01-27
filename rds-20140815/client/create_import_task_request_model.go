// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstanceId(v string) *CreateImportTaskRequest
	GetDbInstanceId() *string
	SetEstimatedSize(v int32) *CreateImportTaskRequest
	GetEstimatedSize() *int32
	SetHost(v string) *CreateImportTaskRequest
	GetHost() *string
	SetOwnerId(v int64) *CreateImportTaskRequest
	GetOwnerId() *int64
	SetPassword(v string) *CreateImportTaskRequest
	GetPassword() *string
	SetPort(v int32) *CreateImportTaskRequest
	GetPort() *int32
	SetRegionId(v string) *CreateImportTaskRequest
	GetRegionId() *string
	SetSourceInstanceId(v string) *CreateImportTaskRequest
	GetSourceInstanceId() *string
	SetSourcePlatform(v string) *CreateImportTaskRequest
	GetSourcePlatform() *string
	SetStreamPort(v int32) *CreateImportTaskRequest
	GetStreamPort() *int32
	SetUser(v string) *CreateImportTaskRequest
	GetUser() *string
	SetXtrabackupPath(v string) *CreateImportTaskRequest
	GetXtrabackupPath() *string
}

type CreateImportTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1u*****ggm7j9j
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// 1000
	EstimatedSize *int32 `json:"EstimatedSize,omitempty" xml:"EstimatedSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 172.20.246.90
	Host    *string `json:"Host,omitempty" xml:"Host,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OEF5JjVOM2pzZXFKRw==
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// i-bp1fe296n52ub3chezpg
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// example:
	//
	// ECS
	SourcePlatform *string `json:"SourcePlatform,omitempty" xml:"SourcePlatform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9999
	StreamPort *int32 `json:"StreamPort,omitempty" xml:"StreamPort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// myadmin
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// example:
	//
	// /usr/bin/xtrabackup
	XtrabackupPath *string `json:"XtrabackupPath,omitempty" xml:"XtrabackupPath,omitempty"`
}

func (s CreateImportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImportTaskRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *CreateImportTaskRequest) GetEstimatedSize() *int32 {
	return s.EstimatedSize
}

func (s *CreateImportTaskRequest) GetHost() *string {
	return s.Host
}

func (s *CreateImportTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateImportTaskRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateImportTaskRequest) GetPort() *int32 {
	return s.Port
}

func (s *CreateImportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImportTaskRequest) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *CreateImportTaskRequest) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *CreateImportTaskRequest) GetStreamPort() *int32 {
	return s.StreamPort
}

func (s *CreateImportTaskRequest) GetUser() *string {
	return s.User
}

func (s *CreateImportTaskRequest) GetXtrabackupPath() *string {
	return s.XtrabackupPath
}

func (s *CreateImportTaskRequest) SetDbInstanceId(v string) *CreateImportTaskRequest {
	s.DbInstanceId = &v
	return s
}

func (s *CreateImportTaskRequest) SetEstimatedSize(v int32) *CreateImportTaskRequest {
	s.EstimatedSize = &v
	return s
}

func (s *CreateImportTaskRequest) SetHost(v string) *CreateImportTaskRequest {
	s.Host = &v
	return s
}

func (s *CreateImportTaskRequest) SetOwnerId(v int64) *CreateImportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateImportTaskRequest) SetPassword(v string) *CreateImportTaskRequest {
	s.Password = &v
	return s
}

func (s *CreateImportTaskRequest) SetPort(v int32) *CreateImportTaskRequest {
	s.Port = &v
	return s
}

func (s *CreateImportTaskRequest) SetRegionId(v string) *CreateImportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImportTaskRequest) SetSourceInstanceId(v string) *CreateImportTaskRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *CreateImportTaskRequest) SetSourcePlatform(v string) *CreateImportTaskRequest {
	s.SourcePlatform = &v
	return s
}

func (s *CreateImportTaskRequest) SetStreamPort(v int32) *CreateImportTaskRequest {
	s.StreamPort = &v
	return s
}

func (s *CreateImportTaskRequest) SetUser(v string) *CreateImportTaskRequest {
	s.User = &v
	return s
}

func (s *CreateImportTaskRequest) SetXtrabackupPath(v string) *CreateImportTaskRequest {
	s.XtrabackupPath = &v
	return s
}

func (s *CreateImportTaskRequest) Validate() error {
	return dara.Validate(s)
}
