// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSecurityEventIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityEventIds(v []*string) *CheckSecurityEventIdRequest
	GetSecurityEventIds() []*string
	SetUuid(v string) *CheckSecurityEventIdRequest
	GetUuid() *string
}

type CheckSecurityEventIdRequest struct {
	// The IDs of the alert events to be verified. You can specify up to 100 alert event IDs. If this parameter is left empty, the verification result **Data*	- defaults to **false**, indicating that the events do not belong to the specified server.
	//
	// > Call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to retrieve this parameter.
	//
	// example:
	//
	// ["1234567","98765432"]
	SecurityEventIds []*string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty" type:"Repeated"`
	// The UUID of the server to be verified.
	//
	// > Call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to retrieve this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 96ad2473-bc60-45ba-ad1c-932e2866****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CheckSecurityEventIdRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckSecurityEventIdRequest) GoString() string {
	return s.String()
}

func (s *CheckSecurityEventIdRequest) GetSecurityEventIds() []*string {
	return s.SecurityEventIds
}

func (s *CheckSecurityEventIdRequest) GetUuid() *string {
	return s.Uuid
}

func (s *CheckSecurityEventIdRequest) SetSecurityEventIds(v []*string) *CheckSecurityEventIdRequest {
	s.SecurityEventIds = v
	return s
}

func (s *CheckSecurityEventIdRequest) SetUuid(v string) *CheckSecurityEventIdRequest {
	s.Uuid = &v
	return s
}

func (s *CheckSecurityEventIdRequest) Validate() error {
	return dara.Validate(s)
}
