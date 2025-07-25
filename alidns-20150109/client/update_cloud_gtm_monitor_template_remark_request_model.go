// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmMonitorTemplateRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmMonitorTemplateRemarkRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *UpdateCloudGtmMonitorTemplateRemarkRequest
	GetClientToken() *string
	SetRemark(v string) *UpdateCloudGtmMonitorTemplateRemarkRequest
	GetRemark() *string
	SetTemplateId(v string) *UpdateCloudGtmMonitorTemplateRemarkRequest
	GetTemplateId() *string
}

type UpdateCloudGtmMonitorTemplateRemarkRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
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
	// The new description of the template. If you do not specify this parameter, the original description is deleted.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the health check template. This ID uniquely identifies a health check template.
	//
	// This parameter is required.
	//
	// example:
	//
	// mtp-89518052455928**00
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateCloudGtmMonitorTemplateRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmMonitorTemplateRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmMonitorTemplateRemarkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmMonitorTemplateRemarkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmMonitorTemplateRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateCloudGtmMonitorTemplateRemarkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateCloudGtmMonitorTemplateRemarkRequest) SetAcceptLanguage(v string) *UpdateCloudGtmMonitorTemplateRemarkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRemarkRequest) SetClientToken(v string) *UpdateCloudGtmMonitorTemplateRemarkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRemarkRequest) SetRemark(v string) *UpdateCloudGtmMonitorTemplateRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRemarkRequest) SetTemplateId(v string) *UpdateCloudGtmMonitorTemplateRemarkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRemarkRequest) Validate() error {
	return dara.Validate(s)
}
