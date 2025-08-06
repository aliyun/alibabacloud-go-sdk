// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVodTemplateId(v string) *GetVodTemplateRequest
	GetVodTemplateId() *string
}

type GetVodTemplateRequest struct {
	// The ID of the snapshot template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c49f2f4c0969*****fcd446690
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s GetVodTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVodTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetVodTemplateRequest) GetVodTemplateId() *string {
	return s.VodTemplateId
}

func (s *GetVodTemplateRequest) SetVodTemplateId(v string) *GetVodTemplateRequest {
	s.VodTemplateId = &v
	return s
}

func (s *GetVodTemplateRequest) Validate() error {
	return dara.Validate(s)
}
