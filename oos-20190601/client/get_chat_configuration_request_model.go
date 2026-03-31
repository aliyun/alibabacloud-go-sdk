// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetChatConfigurationRequest
	GetName() *string
	SetRegionId(v string) *GetChatConfigurationRequest
	GetRegionId() *string
}

type GetChatConfigurationRequest struct {
	// example:
	//
	// chatOps-qiwei
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetChatConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetChatConfigurationRequest) GetName() *string {
	return s.Name
}

func (s *GetChatConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetChatConfigurationRequest) SetName(v string) *GetChatConfigurationRequest {
	s.Name = &v
	return s
}

func (s *GetChatConfigurationRequest) SetRegionId(v string) *GetChatConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *GetChatConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
