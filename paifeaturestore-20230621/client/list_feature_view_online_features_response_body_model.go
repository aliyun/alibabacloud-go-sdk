// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureViewOnlineFeaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOnlineFeatures(v []*string) *ListFeatureViewOnlineFeaturesResponseBody
	GetOnlineFeatures() []*string
	SetRequestId(v string) *ListFeatureViewOnlineFeaturesResponseBody
	GetRequestId() *string
}

type ListFeatureViewOnlineFeaturesResponseBody struct {
	OnlineFeatures []*string `json:"OnlineFeatures,omitempty" xml:"OnlineFeatures,omitempty" type:"Repeated"`
	// example:
	//
	// BF349686-C932-55B5-9B31-DAFA395C0E06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFeatureViewOnlineFeaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewOnlineFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureViewOnlineFeaturesResponseBody) GetOnlineFeatures() []*string {
	return s.OnlineFeatures
}

func (s *ListFeatureViewOnlineFeaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFeatureViewOnlineFeaturesResponseBody) SetOnlineFeatures(v []*string) *ListFeatureViewOnlineFeaturesResponseBody {
	s.OnlineFeatures = v
	return s
}

func (s *ListFeatureViewOnlineFeaturesResponseBody) SetRequestId(v string) *ListFeatureViewOnlineFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeatureViewOnlineFeaturesResponseBody) Validate() error {
	return dara.Validate(s)
}
