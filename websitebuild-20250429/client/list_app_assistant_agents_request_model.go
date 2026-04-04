// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppAssistantAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListAppAssistantAgentsRequest
	GetBizId() *string
	SetPlatformType(v string) *ListAppAssistantAgentsRequest
	GetPlatformType() *string
}

type ListAppAssistantAgentsRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// ChannelOps
	PlatformType *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
}

func (s ListAppAssistantAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppAssistantAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAppAssistantAgentsRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListAppAssistantAgentsRequest) GetPlatformType() *string {
	return s.PlatformType
}

func (s *ListAppAssistantAgentsRequest) SetBizId(v string) *ListAppAssistantAgentsRequest {
	s.BizId = &v
	return s
}

func (s *ListAppAssistantAgentsRequest) SetPlatformType(v string) *ListAppAssistantAgentsRequest {
	s.PlatformType = &v
	return s
}

func (s *ListAppAssistantAgentsRequest) Validate() error {
	return dara.Validate(s)
}
