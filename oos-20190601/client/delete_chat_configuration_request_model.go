// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteChatConfigurationRequest
	GetName() *string
	SetRegionId(v string) *DeleteChatConfigurationRequest
	GetRegionId() *string
}

type DeleteChatConfigurationRequest struct {
	// example:
	//
	// TestChatConfig-9be97b40
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteChatConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteChatConfigurationRequest) GetName() *string {
	return s.Name
}

func (s *DeleteChatConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteChatConfigurationRequest) SetName(v string) *DeleteChatConfigurationRequest {
	s.Name = &v
	return s
}

func (s *DeleteChatConfigurationRequest) SetRegionId(v string) *DeleteChatConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteChatConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
