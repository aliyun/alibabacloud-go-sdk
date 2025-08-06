// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVodTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *AddVodTemplateRequest
	GetAppId() *string
	SetName(v string) *AddVodTemplateRequest
	GetName() *string
	SetTemplateConfig(v string) *AddVodTemplateRequest
	GetTemplateConfig() *string
	SetTemplateType(v string) *AddVodTemplateRequest
	GetTemplateType() *string
}

type AddVodTemplateRequest struct {
	// The ID of the application. Default value: **app-1000000**. For more information, see [Multi-application service](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the template.
	//
	// 	- The name cannot exceed 128 bytes.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of the snapshot template. The value must be a JSON string. For more information about the data structure, see [SnapshotTemplateConfig](https://help.aliyun.com/document_detail/98618.html) and [DynamicImageTemplateConfig](https://help.aliyun.com/document_detail/98618.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"SnapshotConfig":{"Count":10,"SpecifiedOffsetTime":0,"Interval":1,"FrameType":"normal"},"SnapshotType":"NormalSnapshot"}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The type of the template. Set the value to **Snapshot**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Snapshot
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s AddVodTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddVodTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddVodTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddVodTemplateRequest) GetName() *string {
	return s.Name
}

func (s *AddVodTemplateRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *AddVodTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *AddVodTemplateRequest) SetAppId(v string) *AddVodTemplateRequest {
	s.AppId = &v
	return s
}

func (s *AddVodTemplateRequest) SetName(v string) *AddVodTemplateRequest {
	s.Name = &v
	return s
}

func (s *AddVodTemplateRequest) SetTemplateConfig(v string) *AddVodTemplateRequest {
	s.TemplateConfig = &v
	return s
}

func (s *AddVodTemplateRequest) SetTemplateType(v string) *AddVodTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *AddVodTemplateRequest) Validate() error {
	return dara.Validate(s)
}
