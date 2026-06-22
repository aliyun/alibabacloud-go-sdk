// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotListConversationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *KopilotListConversationsRequest
	GetPage() *int32
	SetRegionId(v string) *KopilotListConversationsRequest
	GetRegionId() *string
	SetSize(v int32) *KopilotListConversationsRequest
	GetSize() *int32
}

type KopilotListConversationsRequest struct {
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Size     *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s KopilotListConversationsRequest) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsRequest) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsRequest) GetPage() *int32 {
	return s.Page
}

func (s *KopilotListConversationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *KopilotListConversationsRequest) GetSize() *int32 {
	return s.Size
}

func (s *KopilotListConversationsRequest) SetPage(v int32) *KopilotListConversationsRequest {
	s.Page = &v
	return s
}

func (s *KopilotListConversationsRequest) SetRegionId(v string) *KopilotListConversationsRequest {
	s.RegionId = &v
	return s
}

func (s *KopilotListConversationsRequest) SetSize(v int32) *KopilotListConversationsRequest {
	s.Size = &v
	return s
}

func (s *KopilotListConversationsRequest) Validate() error {
	return dara.Validate(s)
}
