// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeInfo interface {
	dara.Model
	String() string
	GoString() string
	SetHost(v string) *NodeInfo
	GetHost() *string
	SetHostName(v string) *NodeInfo
	GetHostName() *string
	SetPort(v int64) *NodeInfo
	GetPort() *int64
	SetZoneId(v string) *NodeInfo
	GetZoneId() *string
}

type NodeInfo struct {
	Host     *string `json:"host,omitempty" xml:"host,omitempty"`
	HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
	Port     *int64  `json:"port,omitempty" xml:"port,omitempty"`
	ZoneId   *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s NodeInfo) String() string {
	return dara.Prettify(s)
}

func (s NodeInfo) GoString() string {
	return s.String()
}

func (s *NodeInfo) GetHost() *string {
	return s.Host
}

func (s *NodeInfo) GetHostName() *string {
	return s.HostName
}

func (s *NodeInfo) GetPort() *int64 {
	return s.Port
}

func (s *NodeInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *NodeInfo) SetHost(v string) *NodeInfo {
	s.Host = &v
	return s
}

func (s *NodeInfo) SetHostName(v string) *NodeInfo {
	s.HostName = &v
	return s
}

func (s *NodeInfo) SetPort(v int64) *NodeInfo {
	s.Port = &v
	return s
}

func (s *NodeInfo) SetZoneId(v string) *NodeInfo {
	s.ZoneId = &v
	return s
}

func (s *NodeInfo) Validate() error {
	return dara.Validate(s)
}
