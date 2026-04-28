// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataCName interface {
	dara.Model
	String() string
	GoString() string
	SetCertExpireTime(v int64) *DataCName
	GetCertExpireTime() *int64
	SetCertName(v string) *DataCName
	GetCertName() *string
	SetCname(v string) *DataCName
	GetCname() *string
	SetCnameType(v string) *DataCName
	GetCnameType() *string
	SetLocation(v string) *DataCName
	GetLocation() *string
	SetStoreId(v string) *DataCName
	GetStoreId() *string
}

type DataCName struct {
	CertExpireTime *int64  `json:"cert_expire_time,omitempty" xml:"cert_expire_time,omitempty"`
	CertName       *string `json:"cert_name,omitempty" xml:"cert_name,omitempty"`
	Cname          *string `json:"cname,omitempty" xml:"cname,omitempty"`
	CnameType      *string `json:"cname_type,omitempty" xml:"cname_type,omitempty"`
	Location       *string `json:"location,omitempty" xml:"location,omitempty"`
	StoreId        *string `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

func (s DataCName) String() string {
	return dara.Prettify(s)
}

func (s DataCName) GoString() string {
	return s.String()
}

func (s *DataCName) GetCertExpireTime() *int64 {
	return s.CertExpireTime
}

func (s *DataCName) GetCertName() *string {
	return s.CertName
}

func (s *DataCName) GetCname() *string {
	return s.Cname
}

func (s *DataCName) GetCnameType() *string {
	return s.CnameType
}

func (s *DataCName) GetLocation() *string {
	return s.Location
}

func (s *DataCName) GetStoreId() *string {
	return s.StoreId
}

func (s *DataCName) SetCertExpireTime(v int64) *DataCName {
	s.CertExpireTime = &v
	return s
}

func (s *DataCName) SetCertName(v string) *DataCName {
	s.CertName = &v
	return s
}

func (s *DataCName) SetCname(v string) *DataCName {
	s.Cname = &v
	return s
}

func (s *DataCName) SetCnameType(v string) *DataCName {
	s.CnameType = &v
	return s
}

func (s *DataCName) SetLocation(v string) *DataCName {
	s.Location = &v
	return s
}

func (s *DataCName) SetStoreId(v string) *DataCName {
	s.StoreId = &v
	return s
}

func (s *DataCName) Validate() error {
	return dara.Validate(s)
}
