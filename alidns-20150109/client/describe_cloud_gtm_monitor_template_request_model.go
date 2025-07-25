// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmMonitorTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeCloudGtmMonitorTemplateRequest
	GetAcceptLanguage() *string
	SetTemplateId(v string) *DescribeCloudGtmMonitorTemplateRequest
	GetTemplateId() *string
}

type DescribeCloudGtmMonitorTemplateRequest struct {
	// The language of the response. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US*	- (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the health check template that you want to query. This ID uniquely identifies the health check template.
	//
	// This parameter is required.
	//
	// example:
	//
	// mtp-89518052425100**80
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeCloudGtmMonitorTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmMonitorTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmMonitorTemplateRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeCloudGtmMonitorTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeCloudGtmMonitorTemplateRequest) SetAcceptLanguage(v string) *DescribeCloudGtmMonitorTemplateRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateRequest) SetTemplateId(v string) *DescribeCloudGtmMonitorTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeCloudGtmMonitorTemplateRequest) Validate() error {
	return dara.Validate(s)
}
