// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigLbStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudGtmInstanceConfigLbStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmInstanceConfigLbStrategyResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmInstanceConfigLbStrategyResponseBody struct {
	// example:
	//
	// 0F32959D-417B-4D66-8463-68606605E3E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigLbStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigLbStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyResponseBody) SetRequestId(v string) *UpdateCloudGtmInstanceConfigLbStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyResponseBody) SetSuccess(v bool) *UpdateCloudGtmInstanceConfigLbStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
