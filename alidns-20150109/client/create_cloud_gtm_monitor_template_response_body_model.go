// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmMonitorTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCloudGtmMonitorTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCloudGtmMonitorTemplateResponseBody
	GetSuccess() *bool
	SetTemplateId(v string) *CreateCloudGtmMonitorTemplateResponseBody
	GetTemplateId() *string
}

type CreateCloudGtmMonitorTemplateResponseBody struct {
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the health check template. This ID uniquely identifies the health check template.
	//
	// example:
	//
	// mtp-89518052425100**80
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateCloudGtmMonitorTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmMonitorTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmMonitorTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudGtmMonitorTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCloudGtmMonitorTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateCloudGtmMonitorTemplateResponseBody) SetRequestId(v string) *CreateCloudGtmMonitorTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateResponseBody) SetSuccess(v bool) *CreateCloudGtmMonitorTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateResponseBody) SetTemplateId(v string) *CreateCloudGtmMonitorTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateCloudGtmMonitorTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
