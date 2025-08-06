// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVodTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateVodTemplateRequest
	GetName() *string
	SetTemplateConfig(v string) *UpdateVodTemplateRequest
	GetTemplateConfig() *string
	SetVodTemplateId(v string) *UpdateVodTemplateRequest
	GetVodTemplateId() *string
}

type UpdateVodTemplateRequest struct {
	// The name of the template.
	//
	// 	- The name can be up to 128 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of the snapshot template. The value is a JSON-formatted string. For more information about the data structure, see the "SnapshotTemplateConfig" section of the [Media processing parameters](https://help.aliyun.com/document_detail/98618.html) topic.
	//
	// example:
	//
	// {"SnapshotConfig":{"Count":10,"SpecifiedOffsetTime":0,"Interval":1}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The ID of the snapshot template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8c75a02e339b*****0b0d2c48171a22
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s UpdateVodTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVodTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateVodTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateVodTemplateRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *UpdateVodTemplateRequest) GetVodTemplateId() *string {
	return s.VodTemplateId
}

func (s *UpdateVodTemplateRequest) SetName(v string) *UpdateVodTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateVodTemplateRequest) SetTemplateConfig(v string) *UpdateVodTemplateRequest {
	s.TemplateConfig = &v
	return s
}

func (s *UpdateVodTemplateRequest) SetVodTemplateId(v string) *UpdateVodTemplateRequest {
	s.VodTemplateId = &v
	return s
}

func (s *UpdateVodTemplateRequest) Validate() error {
	return dara.Validate(s)
}
