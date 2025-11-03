// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPhysicalConnectionFeaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPhysicalConnectionFeatures(v []*ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures) *ListPhysicalConnectionFeaturesResponseBody
	GetPhysicalConnectionFeatures() []*ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures
	SetRequestId(v string) *ListPhysicalConnectionFeaturesResponseBody
	GetRequestId() *string
}

type ListPhysicalConnectionFeaturesResponseBody struct {
	// The list of Express Connect circuit features.
	PhysicalConnectionFeatures []*ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures `json:"PhysicalConnectionFeatures,omitempty" xml:"PhysicalConnectionFeatures,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A599D38F-3618-18FD-9427-108FB9B5BD26
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPhysicalConnectionFeaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPhysicalConnectionFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPhysicalConnectionFeaturesResponseBody) GetPhysicalConnectionFeatures() []*ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures {
	return s.PhysicalConnectionFeatures
}

func (s *ListPhysicalConnectionFeaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPhysicalConnectionFeaturesResponseBody) SetPhysicalConnectionFeatures(v []*ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures) *ListPhysicalConnectionFeaturesResponseBody {
	s.PhysicalConnectionFeatures = v
	return s
}

func (s *ListPhysicalConnectionFeaturesResponseBody) SetRequestId(v string) *ListPhysicalConnectionFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPhysicalConnectionFeaturesResponseBody) Validate() error {
	if s.PhysicalConnectionFeatures != nil {
		for _, item := range s.PhysicalConnectionFeatures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures struct {
	// The feature key of the Express Connect circuit. Valid values:
	//
	// 	- **SubifRateLimit**: subinterface throttling
	//
	// 	- **BFD Capability**: Bidirectional Forwarding Detection (BFD)
	//
	// 	- **DualStack**: Dual stack
	//
	// 	- **CEN**: When a virtual border router (VBR) is attached to a Cloud Enterprise Network (CEN) instance and BGP routes are advertised on the user side, attributes such as **as-path*	- and **community*	- are carried.
	//
	// 	- **CENv6**: When a VBR is attached to an IPv6 CEN instance and BGP routes are advertised on the user side, attributes such as **as-path*	- and **community*	- are carried.
	//
	// 	- **QOS**: The device supports configuring QOS policies on physical ports.
	//
	// 	- **MSHA**: The device supports fast switching groups between two VBRs.
	//
	// 	- **MULTI_MS_HA**: The device supports a maximum of eight VBRs that can be added to the same ECR.
	//
	// example:
	//
	// SubifRateLimit
	FeatureKey *string `json:"FeatureKey,omitempty" xml:"FeatureKey,omitempty"`
	// The feature value of the Express Connect circuit. Valid values:
	//
	// 	- **OK**: Supported
	//
	// 	- **NOK**: Not supported
	//
	// example:
	//
	// OK
	FeatureValue *string `json:"FeatureValue,omitempty" xml:"FeatureValue,omitempty"`
}

func (s ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures) String() string {
	return dara.Prettify(s)
}

func (s ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures) GoString() string {
	return s.String()
}

func (s *ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures) GetFeatureKey() *string {
	return s.FeatureKey
}

func (s *ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures) GetFeatureValue() *string {
	return s.FeatureValue
}

func (s *ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures) SetFeatureKey(v string) *ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures {
	s.FeatureKey = &v
	return s
}

func (s *ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures) SetFeatureValue(v string) *ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures {
	s.FeatureValue = &v
	return s
}

func (s *ListPhysicalConnectionFeaturesResponseBodyPhysicalConnectionFeatures) Validate() error {
	return dara.Validate(s)
}
