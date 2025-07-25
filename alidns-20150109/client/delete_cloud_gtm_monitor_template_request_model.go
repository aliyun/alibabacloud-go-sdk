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
	// The language in which the returned results are displayed. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the health check template. This ID uniquely identifies a health check template.
	//
	// This parameter is required.
	//
	// example:
	//
	// mtp-89518052425100**80
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
