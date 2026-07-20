// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceControlEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *GetResourceControlEventRequest
	GetAliyunLang() *string
	SetEventId(v string) *GetResourceControlEventRequest
	GetEventId() *string
	SetEventIdList(v []*string) *GetResourceControlEventRequest
	GetEventIdList() []*string
}

type GetResourceControlEventRequest struct {
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 09C-2PpwIzkpx2zG2fuFrAH55CpJaTK
	EventId     *string   `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventIdList []*string `json:"EventIdList,omitempty" xml:"EventIdList,omitempty" type:"Repeated"`
}

func (s GetResourceControlEventRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceControlEventRequest) GoString() string {
	return s.String()
}

func (s *GetResourceControlEventRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *GetResourceControlEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *GetResourceControlEventRequest) GetEventIdList() []*string {
	return s.EventIdList
}

func (s *GetResourceControlEventRequest) SetAliyunLang(v string) *GetResourceControlEventRequest {
	s.AliyunLang = &v
	return s
}

func (s *GetResourceControlEventRequest) SetEventId(v string) *GetResourceControlEventRequest {
	s.EventId = &v
	return s
}

func (s *GetResourceControlEventRequest) SetEventIdList(v []*string) *GetResourceControlEventRequest {
	s.EventIdList = v
	return s
}

func (s *GetResourceControlEventRequest) Validate() error {
	return dara.Validate(s)
}
