// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmMonitorTemplateRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudGtmMonitorTemplateRemarkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmMonitorTemplateRemarkResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmMonitorTemplateRemarkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Modify the health check template remark operation status:
	//
	// - true: Operation successful
	//
	// - false: Operation failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloudGtmMonitorTemplateRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmMonitorTemplateRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmMonitorTemplateRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmMonitorTemplateRemarkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmMonitorTemplateRemarkResponseBody) SetRequestId(v string) *UpdateCloudGtmMonitorTemplateRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRemarkResponseBody) SetSuccess(v bool) *UpdateCloudGtmMonitorTemplateRemarkResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
