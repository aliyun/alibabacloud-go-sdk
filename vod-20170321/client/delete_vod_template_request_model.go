// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVodTemplateId(v string) *DeleteVodTemplateRequest
	GetVodTemplateId() *string
}

type DeleteVodTemplateRequest struct {
	// The ID of the snapshot template.
	//
	// This parameter is required.
	//
	// example:
	//
	// f5b228fe6930e*****d6bf55bd87789
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s DeleteVodTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteVodTemplateRequest) GetVodTemplateId() *string {
	return s.VodTemplateId
}

func (s *DeleteVodTemplateRequest) SetVodTemplateId(v string) *DeleteVodTemplateRequest {
	s.VodTemplateId = &v
	return s
}

func (s *DeleteVodTemplateRequest) Validate() error {
	return dara.Validate(s)
}
