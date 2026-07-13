// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmMonitorTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudGtmMonitorTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmMonitorTemplateResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmMonitorTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloudGtmMonitorTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmMonitorTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmMonitorTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmMonitorTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmMonitorTemplateResponseBody) SetRequestId(v string) *UpdateCloudGtmMonitorTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateResponseBody) SetSuccess(v bool) *UpdateCloudGtmMonitorTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
