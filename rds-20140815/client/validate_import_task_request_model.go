// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateImportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstanceId(v string) *ValidateImportTaskRequest
	GetDbInstanceId() *string
	SetEstimatedSize(v int32) *ValidateImportTaskRequest
	GetEstimatedSize() *int32
	SetHost(v string) *ValidateImportTaskRequest
	GetHost() *string
	SetOwnerId(v int64) *ValidateImportTaskRequest
	GetOwnerId() *int64
	SetPassword(v string) *ValidateImportTaskRequest
	GetPassword() *string
	SetPort(v int32) *ValidateImportTaskRequest
	GetPort() *int32
	SetRegionId(v string) *ValidateImportTaskRequest
	GetRegionId() *string
	SetSourceInstanceId(v string) *ValidateImportTaskRequest
	GetSourceInstanceId() *string
	SetSourcePlatform(v string) *ValidateImportTaskRequest
	GetSourcePlatform() *string
	SetStreamPort(v int32) *ValidateImportTaskRequest
	GetStreamPort() *int32
	SetUser(v string) *ValidateImportTaskRequest
	GetUser() *string
	SetXtrabackupPath(v string) *ValidateImportTaskRequest
	GetXtrabackupPath() *string
}

type ValidateImportTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-sdfljk123****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// 100
	EstimatedSize *int32 `json:"EstimatedSize,omitempty" xml:"EstimatedSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 192.168.10.1
	Host    *string `json:"Host,omitempty" xml:"Host,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UGFzc3dvcmQxMjMK
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
	// i-wz9ff3acy500io5wdf5s
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
	// /usr/local/bin/xtrabackup
	XtrabackupPath *string `json:"XtrabackupPath,omitempty" xml:"XtrabackupPath,omitempty"`
}

func (s ValidateImportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateImportTaskRequest) GoString() string {
	return s.String()
}

func (s *ValidateImportTaskRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *ValidateImportTaskRequest) GetEstimatedSize() *int32 {
	return s.EstimatedSize
}

func (s *ValidateImportTaskRequest) GetHost() *string {
	return s.Host
}

func (s *ValidateImportTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ValidateImportTaskRequest) GetPassword() *string {
	return s.Password
}

func (s *ValidateImportTaskRequest) GetPort() *int32 {
	return s.Port
}

func (s *ValidateImportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ValidateImportTaskRequest) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *ValidateImportTaskRequest) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *ValidateImportTaskRequest) GetStreamPort() *int32 {
	return s.StreamPort
}

func (s *ValidateImportTaskRequest) GetUser() *string {
	return s.User
}

func (s *ValidateImportTaskRequest) GetXtrabackupPath() *string {
	return s.XtrabackupPath
}

func (s *ValidateImportTaskRequest) SetDbInstanceId(v string) *ValidateImportTaskRequest {
	s.DbInstanceId = &v
	return s
}

func (s *ValidateImportTaskRequest) SetEstimatedSize(v int32) *ValidateImportTaskRequest {
	s.EstimatedSize = &v
	return s
}

func (s *ValidateImportTaskRequest) SetHost(v string) *ValidateImportTaskRequest {
	s.Host = &v
	return s
}

func (s *ValidateImportTaskRequest) SetOwnerId(v int64) *ValidateImportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ValidateImportTaskRequest) SetPassword(v string) *ValidateImportTaskRequest {
	s.Password = &v
	return s
}

func (s *ValidateImportTaskRequest) SetPort(v int32) *ValidateImportTaskRequest {
	s.Port = &v
	return s
}

func (s *ValidateImportTaskRequest) SetRegionId(v string) *ValidateImportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ValidateImportTaskRequest) SetSourceInstanceId(v string) *ValidateImportTaskRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *ValidateImportTaskRequest) SetSourcePlatform(v string) *ValidateImportTaskRequest {
	s.SourcePlatform = &v
	return s
}

func (s *ValidateImportTaskRequest) SetStreamPort(v int32) *ValidateImportTaskRequest {
	s.StreamPort = &v
	return s
}

func (s *ValidateImportTaskRequest) SetUser(v string) *ValidateImportTaskRequest {
	s.User = &v
	return s
}

func (s *ValidateImportTaskRequest) SetXtrabackupPath(v string) *ValidateImportTaskRequest {
	s.XtrabackupPath = &v
	return s
}

func (s *ValidateImportTaskRequest) Validate() error {
	return dara.Validate(s)
}
