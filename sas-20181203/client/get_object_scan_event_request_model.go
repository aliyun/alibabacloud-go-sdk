// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectScanEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *GetObjectScanEventRequest
	GetEventId() *string
	SetLang(v string) *GetObjectScanEventRequest
	GetLang() *string
}

type GetObjectScanEventRequest struct {
	// The ID of the alert event.
	//
	// example:
	//
	// 81****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The language of the content in the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetObjectScanEventRequest) String() string {
	return dara.Prettify(s)
}

func (s GetObjectScanEventRequest) GoString() string {
	return s.String()
}

func (s *GetObjectScanEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *GetObjectScanEventRequest) GetLang() *string {
	return s.Lang
}

func (s *GetObjectScanEventRequest) SetEventId(v string) *GetObjectScanEventRequest {
	s.EventId = &v
	return s
}

func (s *GetObjectScanEventRequest) SetLang(v string) *GetObjectScanEventRequest {
	s.Lang = &v
	return s
}

func (s *GetObjectScanEventRequest) Validate() error {
	return dara.Validate(s)
}
