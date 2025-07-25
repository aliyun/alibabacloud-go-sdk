// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudGtmMonitorTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudGtmMonitorTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCloudGtmMonitorTemplateResponseBody
	GetSuccess() *bool
}

type DeleteCloudGtmMonitorTemplateResponseBody struct {
	// Unique request identification code.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation to delete the health check template was successful:
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

func (s DeleteCloudGtmMonitorTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudGtmMonitorTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudGtmMonitorTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudGtmMonitorTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCloudGtmMonitorTemplateResponseBody) SetRequestId(v string) *DeleteCloudGtmMonitorTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudGtmMonitorTemplateResponseBody) SetSuccess(v bool) *DeleteCloudGtmMonitorTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCloudGtmMonitorTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
