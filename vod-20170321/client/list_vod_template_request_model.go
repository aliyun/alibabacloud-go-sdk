// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListVodTemplateRequest
	GetAppId() *string
	SetTemplateType(v string) *ListVodTemplateRequest
	GetTemplateType() *string
}

type ListVodTemplateRequest struct {
	// The ID of the application. Set the value to **app-1000000**. For more information, see [Overview](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The type of the template. Set the value to **Snapshot**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Snapshot
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListVodTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVodTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListVodTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListVodTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListVodTemplateRequest) SetAppId(v string) *ListVodTemplateRequest {
	s.AppId = &v
	return s
}

func (s *ListVodTemplateRequest) SetTemplateType(v string) *ListVodTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *ListVodTemplateRequest) Validate() error {
	return dara.Validate(s)
}
