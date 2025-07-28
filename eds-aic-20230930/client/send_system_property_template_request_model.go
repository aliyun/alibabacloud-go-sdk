// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSystemPropertyTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *SendSystemPropertyTemplateRequest
	GetAndroidInstanceIds() []*string
	SetTemplateId(v string) *SendSystemPropertyTemplateRequest
	GetTemplateId() *string
	SetTemplateIds(v []*string) *SendSystemPropertyTemplateRequest
	GetTemplateIds() []*string
}

type SendSystemPropertyTemplateRequest struct {
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// ap-angyvganxlf****
	TemplateId  *string   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
}

func (s SendSystemPropertyTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s SendSystemPropertyTemplateRequest) GoString() string {
	return s.String()
}

func (s *SendSystemPropertyTemplateRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *SendSystemPropertyTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SendSystemPropertyTemplateRequest) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *SendSystemPropertyTemplateRequest) SetAndroidInstanceIds(v []*string) *SendSystemPropertyTemplateRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *SendSystemPropertyTemplateRequest) SetTemplateId(v string) *SendSystemPropertyTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *SendSystemPropertyTemplateRequest) SetTemplateIds(v []*string) *SendSystemPropertyTemplateRequest {
	s.TemplateIds = v
	return s
}

func (s *SendSystemPropertyTemplateRequest) Validate() error {
	return dara.Validate(s)
}
