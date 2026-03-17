// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartAGByAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSmartAGByAccessPointResponseBody
	GetRequestId() *string
	SetSmartAccessGateways(v []*ListSmartAGByAccessPointResponseBodySmartAccessGateways) *ListSmartAGByAccessPointResponseBody
	GetSmartAccessGateways() []*ListSmartAGByAccessPointResponseBodySmartAccessGateways
	SetTotalCount(v int32) *ListSmartAGByAccessPointResponseBody
	GetTotalCount() *int32
}

type ListSmartAGByAccessPointResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AE203140-5D0C-4B4D-88D1-D008206B3A01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the SAG instance.
	SmartAccessGateways []*ListSmartAGByAccessPointResponseBodySmartAccessGateways `json:"SmartAccessGateways,omitempty" xml:"SmartAccessGateways,omitempty" type:"Repeated"`
	// The number of SAG instances within the access point.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSmartAGByAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSmartAGByAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmartAGByAccessPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSmartAGByAccessPointResponseBody) GetSmartAccessGateways() []*ListSmartAGByAccessPointResponseBodySmartAccessGateways {
	return s.SmartAccessGateways
}

func (s *ListSmartAGByAccessPointResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSmartAGByAccessPointResponseBody) SetRequestId(v string) *ListSmartAGByAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSmartAGByAccessPointResponseBody) SetSmartAccessGateways(v []*ListSmartAGByAccessPointResponseBodySmartAccessGateways) *ListSmartAGByAccessPointResponseBody {
	s.SmartAccessGateways = v
	return s
}

func (s *ListSmartAGByAccessPointResponseBody) SetTotalCount(v int32) *ListSmartAGByAccessPointResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSmartAGByAccessPointResponseBody) Validate() error {
	if s.SmartAccessGateways != nil {
		for _, item := range s.SmartAccessGateways {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSmartAGByAccessPointResponseBodySmartAccessGateways struct {
	// The ID of the Cloud Connect Network (CCN) instance with which the SAG instance is associated.
	//
	// example:
	//
	// ccn-l42qf3vpvb****
	AssociatedCcnId *string `json:"AssociatedCcnId,omitempty" xml:"AssociatedCcnId,omitempty"`
	// The model of the SAG device with which the SAG instance is associated. Valid values:
	//
	// 	- **sag-1000**.
	//
	// 	- **sag-100WM**.
	//
	// example:
	//
	// sag-1000
	HardwareVersion *string `json:"HardwareVersion,omitempty" xml:"HardwareVersion,omitempty"`
	// The method that the SAG instance uses to synchronize Alibaba Cloud-facing routes. Valid values:
	//
	// 	- **static**: static routing.
	//
	// 	- **dynamic**: dynamic routing.
	//
	// example:
	//
	// static
	RoutingStrategy *string `json:"RoutingStrategy,omitempty" xml:"RoutingStrategy,omitempty"`
	// The description of the SAG instance.
	//
	// example:
	//
	// test
	SmartAGDescription *string `json:"SmartAGDescription,omitempty" xml:"SmartAGDescription,omitempty"`
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-p86e06z4geaji1****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The name of the SAG instance.
	//
	// example:
	//
	// test
	SmartAGName *string `json:"SmartAGName,omitempty" xml:"SmartAGName,omitempty"`
	// The status of the SAG instance. Valid values:
	//
	// 	- **Active**: The SAG device is connected to Alibaba Cloud.
	//
	// 	- **offline**: The SAG device is disconnected from Alibaba Cloud.
	//
	// example:
	//
	// Active
	SmartAGStatus *string `json:"SmartAGStatus,omitempty" xml:"SmartAGStatus,omitempty"`
}

func (s ListSmartAGByAccessPointResponseBodySmartAccessGateways) String() string {
	return dara.Prettify(s)
}

func (s ListSmartAGByAccessPointResponseBodySmartAccessGateways) GoString() string {
	return s.String()
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) GetAssociatedCcnId() *string {
	return s.AssociatedCcnId
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) GetHardwareVersion() *string {
	return s.HardwareVersion
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) GetRoutingStrategy() *string {
	return s.RoutingStrategy
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) GetSmartAGDescription() *string {
	return s.SmartAGDescription
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) GetSmartAGName() *string {
	return s.SmartAGName
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) GetSmartAGStatus() *string {
	return s.SmartAGStatus
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) SetAssociatedCcnId(v string) *ListSmartAGByAccessPointResponseBodySmartAccessGateways {
	s.AssociatedCcnId = &v
	return s
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) SetHardwareVersion(v string) *ListSmartAGByAccessPointResponseBodySmartAccessGateways {
	s.HardwareVersion = &v
	return s
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) SetRoutingStrategy(v string) *ListSmartAGByAccessPointResponseBodySmartAccessGateways {
	s.RoutingStrategy = &v
	return s
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) SetSmartAGDescription(v string) *ListSmartAGByAccessPointResponseBodySmartAccessGateways {
	s.SmartAGDescription = &v
	return s
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) SetSmartAGId(v string) *ListSmartAGByAccessPointResponseBodySmartAccessGateways {
	s.SmartAGId = &v
	return s
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) SetSmartAGName(v string) *ListSmartAGByAccessPointResponseBodySmartAccessGateways {
	s.SmartAGName = &v
	return s
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) SetSmartAGStatus(v string) *ListSmartAGByAccessPointResponseBodySmartAccessGateways {
	s.SmartAGStatus = &v
	return s
}

func (s *ListSmartAGByAccessPointResponseBodySmartAccessGateways) Validate() error {
	return dara.Validate(s)
}
