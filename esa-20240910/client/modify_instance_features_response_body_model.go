// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceFeaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedFeatures(v string) *ModifyInstanceFeaturesResponseBody
	GetFailedFeatures() *string
	SetRequestId(v string) *ModifyInstanceFeaturesResponseBody
	GetRequestId() *string
}

type ModifyInstanceFeaturesResponseBody struct {
	// The site feature configurations that failed to be modified.
	//
	// example:
	//
	// loadbalance
	FailedFeatures *string `json:"FailedFeatures,omitempty" xml:"FailedFeatures,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2430E05E-1340-5773-B5E1-B743929F46F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceFeaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceFeaturesResponseBody) GetFailedFeatures() *string {
	return s.FailedFeatures
}

func (s *ModifyInstanceFeaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceFeaturesResponseBody) SetFailedFeatures(v string) *ModifyInstanceFeaturesResponseBody {
	s.FailedFeatures = &v
	return s
}

func (s *ModifyInstanceFeaturesResponseBody) SetRequestId(v string) *ModifyInstanceFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceFeaturesResponseBody) Validate() error {
	return dara.Validate(s)
}
