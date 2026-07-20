// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceControlEventShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *GetResourceControlEventShrinkRequest
	GetAliyunLang() *string
	SetEventId(v string) *GetResourceControlEventShrinkRequest
	GetEventId() *string
	SetEventIdListShrink(v string) *GetResourceControlEventShrinkRequest
	GetEventIdListShrink() *string
}

type GetResourceControlEventShrinkRequest struct {
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 09C-2PpwIzkpx2zG2fuFrAH55CpJaTK
	EventId           *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventIdListShrink *string `json:"EventIdList,omitempty" xml:"EventIdList,omitempty"`
}

func (s GetResourceControlEventShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceControlEventShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetResourceControlEventShrinkRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *GetResourceControlEventShrinkRequest) GetEventId() *string {
	return s.EventId
}

func (s *GetResourceControlEventShrinkRequest) GetEventIdListShrink() *string {
	return s.EventIdListShrink
}

func (s *GetResourceControlEventShrinkRequest) SetAliyunLang(v string) *GetResourceControlEventShrinkRequest {
	s.AliyunLang = &v
	return s
}

func (s *GetResourceControlEventShrinkRequest) SetEventId(v string) *GetResourceControlEventShrinkRequest {
	s.EventId = &v
	return s
}

func (s *GetResourceControlEventShrinkRequest) SetEventIdListShrink(v string) *GetResourceControlEventShrinkRequest {
	s.EventIdListShrink = &v
	return s
}

func (s *GetResourceControlEventShrinkRequest) Validate() error {
	return dara.Validate(s)
}
