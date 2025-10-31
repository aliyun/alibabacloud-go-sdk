// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAnycastEipAddressAssociationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnycastId(v string) *UpdateAnycastEipAddressAssociationsRequest
	GetAnycastId() *string
	SetAssociationMode(v string) *UpdateAnycastEipAddressAssociationsRequest
	GetAssociationMode() *string
	SetBindInstanceId(v string) *UpdateAnycastEipAddressAssociationsRequest
	GetBindInstanceId() *string
	SetClientToken(v string) *UpdateAnycastEipAddressAssociationsRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateAnycastEipAddressAssociationsRequest
	GetDryRun() *bool
	SetPopLocationAddList(v []*UpdateAnycastEipAddressAssociationsRequestPopLocationAddList) *UpdateAnycastEipAddressAssociationsRequest
	GetPopLocationAddList() []*UpdateAnycastEipAddressAssociationsRequestPopLocationAddList
	SetPopLocationDeleteList(v []*UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList) *UpdateAnycastEipAddressAssociationsRequest
	GetPopLocationDeleteList() []*UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList
}

type UpdateAnycastEipAddressAssociationsRequest struct {
	// The ID of the Anycast EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// aeip-bp1ix34fralt4ykf3****
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The association mode. Valid values:
	//
	// 	- **Default**: the default mode. In this mode, cloud resources to be associated are set as default origin servers.
	//
	// 	- **Normal**: the standard mode. In this mode, cloud resources to be associated are set as standard origin servers.
	//
	// example:
	//
	// Default
	AssociationMode *string `json:"AssociationMode,omitempty" xml:"AssociationMode,omitempty"`
	// The ID of the cloud resource with which you want to associate the Anycast EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-d7oxbixhxv1uupnon****
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not set this parameter, the system automatically uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// 	- **true**: prechecks the request without updating the association information. The system checks the required parameters, request syntax, and limits. If the request fails to pass the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the API request. If the request passes the precheck, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The access areas and access points to be added.
	PopLocationAddList []*UpdateAnycastEipAddressAssociationsRequestPopLocationAddList `json:"PopLocationAddList,omitempty" xml:"PopLocationAddList,omitempty" type:"Repeated"`
	// The access areas and access points to be deleted.
	PopLocationDeleteList []*UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList `json:"PopLocationDeleteList,omitempty" xml:"PopLocationDeleteList,omitempty" type:"Repeated"`
}

func (s UpdateAnycastEipAddressAssociationsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAnycastEipAddressAssociationsRequest) GoString() string {
	return s.String()
}

func (s *UpdateAnycastEipAddressAssociationsRequest) GetAnycastId() *string {
	return s.AnycastId
}

func (s *UpdateAnycastEipAddressAssociationsRequest) GetAssociationMode() *string {
	return s.AssociationMode
}

func (s *UpdateAnycastEipAddressAssociationsRequest) GetBindInstanceId() *string {
	return s.BindInstanceId
}

func (s *UpdateAnycastEipAddressAssociationsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAnycastEipAddressAssociationsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateAnycastEipAddressAssociationsRequest) GetPopLocationAddList() []*UpdateAnycastEipAddressAssociationsRequestPopLocationAddList {
	return s.PopLocationAddList
}

func (s *UpdateAnycastEipAddressAssociationsRequest) GetPopLocationDeleteList() []*UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList {
	return s.PopLocationDeleteList
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetAnycastId(v string) *UpdateAnycastEipAddressAssociationsRequest {
	s.AnycastId = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetAssociationMode(v string) *UpdateAnycastEipAddressAssociationsRequest {
	s.AssociationMode = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetBindInstanceId(v string) *UpdateAnycastEipAddressAssociationsRequest {
	s.BindInstanceId = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetClientToken(v string) *UpdateAnycastEipAddressAssociationsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetDryRun(v bool) *UpdateAnycastEipAddressAssociationsRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetPopLocationAddList(v []*UpdateAnycastEipAddressAssociationsRequestPopLocationAddList) *UpdateAnycastEipAddressAssociationsRequest {
	s.PopLocationAddList = v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequest) SetPopLocationDeleteList(v []*UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList) *UpdateAnycastEipAddressAssociationsRequest {
	s.PopLocationDeleteList = v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequest) Validate() error {
	if s.PopLocationAddList != nil {
		for _, item := range s.PopLocationAddList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PopLocationDeleteList != nil {
		for _, item := range s.PopLocationDeleteList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAnycastEipAddressAssociationsRequestPopLocationAddList struct {
	// The access points in the access areas to be added.
	//
	// You can call the [DescribeAnycastPopLocations](https://help.aliyun.com/document_detail/171938.html) operation to query the access points in supported access areas.
	//
	// example:
	//
	// us-west-1-pop
	PopLocation *string `json:"PopLocation,omitempty" xml:"PopLocation,omitempty"`
}

func (s UpdateAnycastEipAddressAssociationsRequestPopLocationAddList) String() string {
	return dara.Prettify(s)
}

func (s UpdateAnycastEipAddressAssociationsRequestPopLocationAddList) GoString() string {
	return s.String()
}

func (s *UpdateAnycastEipAddressAssociationsRequestPopLocationAddList) GetPopLocation() *string {
	return s.PopLocation
}

func (s *UpdateAnycastEipAddressAssociationsRequestPopLocationAddList) SetPopLocation(v string) *UpdateAnycastEipAddressAssociationsRequestPopLocationAddList {
	s.PopLocation = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequestPopLocationAddList) Validate() error {
	return dara.Validate(s)
}

type UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList struct {
	// The access points in the access areas to be deleted.
	//
	// >  If the access point in the access area is associated with a default origin server, you cannot delete the access point in the access area.
	//
	// example:
	//
	// eu-west-1-pop
	PopLocation *string `json:"PopLocation,omitempty" xml:"PopLocation,omitempty"`
}

func (s UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList) String() string {
	return dara.Prettify(s)
}

func (s UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList) GoString() string {
	return s.String()
}

func (s *UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList) GetPopLocation() *string {
	return s.PopLocation
}

func (s *UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList) SetPopLocation(v string) *UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList {
	s.PopLocation = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsRequestPopLocationDeleteList) Validate() error {
	return dara.Validate(s)
}
