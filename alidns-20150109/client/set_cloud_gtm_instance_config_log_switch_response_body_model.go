// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCloudGtmInstanceConfigLogSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetCloudGtmInstanceConfigLogSwitchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetCloudGtmInstanceConfigLogSwitchResponseBody
	GetSuccess() *bool
}

type SetCloudGtmInstanceConfigLogSwitchResponseBody struct {
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetCloudGtmInstanceConfigLogSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCloudGtmInstanceConfigLogSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *SetCloudGtmInstanceConfigLogSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCloudGtmInstanceConfigLogSwitchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetCloudGtmInstanceConfigLogSwitchResponseBody) SetRequestId(v string) *SetCloudGtmInstanceConfigLogSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCloudGtmInstanceConfigLogSwitchResponseBody) SetSuccess(v bool) *SetCloudGtmInstanceConfigLogSwitchResponseBody {
	s.Success = &v
	return s
}

func (s *SetCloudGtmInstanceConfigLogSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
