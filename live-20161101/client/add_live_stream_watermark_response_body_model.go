// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveStreamWatermarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveStreamWatermarkResponseBody
	GetRequestId() *string
	SetTemplateId(v string) *AddLiveStreamWatermarkResponseBody
	GetTemplateId() *string
}

type AddLiveStreamWatermarkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0df228-4a64- af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the watermark template.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AddLiveStreamWatermarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveStreamWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveStreamWatermarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveStreamWatermarkResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AddLiveStreamWatermarkResponseBody) SetRequestId(v string) *AddLiveStreamWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveStreamWatermarkResponseBody) SetTemplateId(v string) *AddLiveStreamWatermarkResponseBody {
	s.TemplateId = &v
	return s
}

func (s *AddLiveStreamWatermarkResponseBody) Validate() error {
	return dara.Validate(s)
}
