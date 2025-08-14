// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRecordTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *DeleteLiveRecordTemplateRequest
	GetTemplateId() *string
}

type DeleteLiveRecordTemplateRequest struct {
	// The ID of the template to be deleted. To obtain the template ID, log on to the [Intelligent Media Services (IMS) console](https://ice.console.aliyun.com/live-processing/template/list/record), choose Real-time Media Processing > Template Management, and then click the Recording tab. Alternatively, find the ID from the response parameters of the [CreateLiveRecordTemplate](https://help.aliyun.com/document_detail/448213.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 69e1f9fe-1e97-11ed-ba64-0c42a1b73d66
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteLiveRecordTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteLiveRecordTemplateRequest) SetTemplateId(v string) *DeleteLiveRecordTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteLiveRecordTemplateRequest) Validate() error {
	return dara.Validate(s)
}
