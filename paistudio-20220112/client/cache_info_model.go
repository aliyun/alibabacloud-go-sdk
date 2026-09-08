// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCacheInfo interface {
	dara.Model
	String() string
	GoString() string
	SetMountPoint(v string) *CacheInfo
	GetMountPoint() *string
	SetPort(v string) *CacheInfo
	GetPort() *string
}

type CacheInfo struct {
	// The mount point address of the datasource for service registration in the cache service, such as an OSS Bucket path or a CPFS path.
	//
	// example:
	//
	// oss://your-bucket.oss-cn-wulanchabu-internal.aliyuncs.com/
	MountPoint *string `json:"MountPoint,omitempty" xml:"MountPoint,omitempty"`
	// The port number that the cache service provides for external access to the datasource. The client must access cached data through this port.
	//
	// example:
	//
	// 10080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CacheInfo) String() string {
	return dara.Prettify(s)
}

func (s CacheInfo) GoString() string {
	return s.String()
}

func (s *CacheInfo) GetMountPoint() *string {
	return s.MountPoint
}

func (s *CacheInfo) GetPort() *string {
	return s.Port
}

func (s *CacheInfo) SetMountPoint(v string) *CacheInfo {
	s.MountPoint = &v
	return s
}

func (s *CacheInfo) SetPort(v string) *CacheInfo {
	s.Port = &v
	return s
}

func (s *CacheInfo) Validate() error {
	return dara.Validate(s)
}
