// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdcProbeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuid(v string) *DeleteIdcProbeRequest
	GetUuid() *string
}

type DeleteIdcProbeRequest struct {
	// The UUID of the server.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// inet-1234567****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DeleteIdcProbeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdcProbeRequest) GoString() string {
	return s.String()
}

func (s *DeleteIdcProbeRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DeleteIdcProbeRequest) SetUuid(v string) *DeleteIdcProbeRequest {
	s.Uuid = &v
	return s
}

func (s *DeleteIdcProbeRequest) Validate() error {
	return dara.Validate(s)
}
