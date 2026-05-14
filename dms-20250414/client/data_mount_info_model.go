// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataMountInfo interface {
	dara.Model
	String() string
	GoString() string
	SetMountFolderName(v string) *DataMountInfo
	GetMountFolderName() *string
	SetOssBucket(v string) *DataMountInfo
	GetOssBucket() *string
	SetPrefix(v string) *DataMountInfo
	GetPrefix() *string
	SetReadOnly(v bool) *DataMountInfo
	GetReadOnly() *bool
}

type DataMountInfo struct {
	MountFolderName *string `json:"MountFolderName,omitempty" xml:"MountFolderName,omitempty"`
	OssBucket       *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	Prefix          *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	ReadOnly        *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
}

func (s DataMountInfo) String() string {
	return dara.Prettify(s)
}

func (s DataMountInfo) GoString() string {
	return s.String()
}

func (s *DataMountInfo) GetMountFolderName() *string {
	return s.MountFolderName
}

func (s *DataMountInfo) GetOssBucket() *string {
	return s.OssBucket
}

func (s *DataMountInfo) GetPrefix() *string {
	return s.Prefix
}

func (s *DataMountInfo) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DataMountInfo) SetMountFolderName(v string) *DataMountInfo {
	s.MountFolderName = &v
	return s
}

func (s *DataMountInfo) SetOssBucket(v string) *DataMountInfo {
	s.OssBucket = &v
	return s
}

func (s *DataMountInfo) SetPrefix(v string) *DataMountInfo {
	s.Prefix = &v
	return s
}

func (s *DataMountInfo) SetReadOnly(v bool) *DataMountInfo {
	s.ReadOnly = &v
	return s
}

func (s *DataMountInfo) Validate() error {
	return dara.Validate(s)
}
