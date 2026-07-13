// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudGtmMonitorTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteCloudGtmMonitorTemplateRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *DeleteCloudGtmMonitorTemplateRequest
	GetClientToken() *string
	SetTemplateId(v string) *DeleteCloudGtmMonitorTemplateRequest
	GetTemplateId() *string
}

type DeleteCloudGtmMonitorTemplateRequest struct {
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US (default): English.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// A client-generated token that is used to ensure the idempotence of the request. This token must be unique for each request and can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The unique ID of the health check template.
	//
	// This parameter is required.
	//
	// example:
	//
	// mtp-89518052425100****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteCloudGtmMonitorTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudGtmMonitorTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudGtmMonitorTemplateRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteCloudGtmMonitorTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCloudGtmMonitorTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteCloudGtmMonitorTemplateRequest) SetAcceptLanguage(v string) *DeleteCloudGtmMonitorTemplateRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteCloudGtmMonitorTemplateRequest) SetClientToken(v string) *DeleteCloudGtmMonitorTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCloudGtmMonitorTemplateRequest) SetTemplateId(v string) *DeleteCloudGtmMonitorTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteCloudGtmMonitorTemplateRequest) Validate() error {
	return dara.Validate(s)
}
