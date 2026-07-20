// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepayResource interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogId(v string) *PrepayResource
	GetCatalogId() *string
	SetCatalogName(v string) *PrepayResource
	GetCatalogName() *string
	SetCu(v int32) *PrepayResource
	GetCu() *int32
	SetExpireTime(v int64) *PrepayResource
	GetExpireTime() *int64
	SetGmtCreate(v int64) *PrepayResource
	GetGmtCreate() *int64
	SetInstanceId(v string) *PrepayResource
	GetInstanceId() *string
	SetInstanceStatus(v string) *PrepayResource
	GetInstanceStatus() *string
}

type PrepayResource struct {
	CatalogId      *string `json:"catalogId,omitempty" xml:"catalogId,omitempty"`
	CatalogName    *string `json:"catalogName,omitempty" xml:"catalogName,omitempty"`
	Cu             *int32  `json:"cu,omitempty" xml:"cu,omitempty"`
	ExpireTime     *int64  `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	GmtCreate      *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	InstanceId     *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
}

func (s PrepayResource) String() string {
	return dara.Prettify(s)
}

func (s PrepayResource) GoString() string {
	return s.String()
}

func (s *PrepayResource) GetCatalogId() *string {
	return s.CatalogId
}

func (s *PrepayResource) GetCatalogName() *string {
	return s.CatalogName
}

func (s *PrepayResource) GetCu() *int32 {
	return s.Cu
}

func (s *PrepayResource) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *PrepayResource) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *PrepayResource) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PrepayResource) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *PrepayResource) SetCatalogId(v string) *PrepayResource {
	s.CatalogId = &v
	return s
}

func (s *PrepayResource) SetCatalogName(v string) *PrepayResource {
	s.CatalogName = &v
	return s
}

func (s *PrepayResource) SetCu(v int32) *PrepayResource {
	s.Cu = &v
	return s
}

func (s *PrepayResource) SetExpireTime(v int64) *PrepayResource {
	s.ExpireTime = &v
	return s
}

func (s *PrepayResource) SetGmtCreate(v int64) *PrepayResource {
	s.GmtCreate = &v
	return s
}

func (s *PrepayResource) SetInstanceId(v string) *PrepayResource {
	s.InstanceId = &v
	return s
}

func (s *PrepayResource) SetInstanceStatus(v string) *PrepayResource {
	s.InstanceStatus = &v
	return s
}

func (s *PrepayResource) Validate() error {
	return dara.Validate(s)
}
