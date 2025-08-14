// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveTranscodeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *GetLiveTranscodeTemplateRequest
	GetTemplateId() *string
}

type GetLiveTranscodeTemplateRequest struct {
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287666****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetLiveTranscodeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLiveTranscodeTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetLiveTranscodeTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLiveTranscodeTemplateRequest) SetTemplateId(v string) *GetLiveTranscodeTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetLiveTranscodeTemplateRequest) Validate() error {
	return dara.Validate(s)
}
