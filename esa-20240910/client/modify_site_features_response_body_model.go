// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySiteFeaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedFeatures(v string) *ModifySiteFeaturesResponseBody
	GetFailedFeatures() *string
	SetRequestId(v string) *ModifySiteFeaturesResponseBody
	GetRequestId() *string
}

type ModifySiteFeaturesResponseBody struct {
	// The site feature information that failed to be cleared.
	//
	// example:
	//
	// loadbalance
	FailedFeatures *string `json:"FailedFeatures,omitempty" xml:"FailedFeatures,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySiteFeaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySiteFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySiteFeaturesResponseBody) GetFailedFeatures() *string {
	return s.FailedFeatures
}

func (s *ModifySiteFeaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySiteFeaturesResponseBody) SetFailedFeatures(v string) *ModifySiteFeaturesResponseBody {
	s.FailedFeatures = &v
	return s
}

func (s *ModifySiteFeaturesResponseBody) SetRequestId(v string) *ModifySiteFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySiteFeaturesResponseBody) Validate() error {
	return dara.Validate(s)
}
