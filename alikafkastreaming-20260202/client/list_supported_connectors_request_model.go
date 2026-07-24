// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportedConnectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListSupportedConnectorsRequest
	GetInstanceId() *string
	SetRegionId(v string) *ListSupportedConnectorsRequest
	GetRegionId() *string
}

type ListSupportedConnectorsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListSupportedConnectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSupportedConnectorsRequest) GoString() string {
	return s.String()
}

func (s *ListSupportedConnectorsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSupportedConnectorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSupportedConnectorsRequest) SetInstanceId(v string) *ListSupportedConnectorsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSupportedConnectorsRequest) SetRegionId(v string) *ListSupportedConnectorsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSupportedConnectorsRequest) Validate() error {
	return dara.Validate(s)
}
