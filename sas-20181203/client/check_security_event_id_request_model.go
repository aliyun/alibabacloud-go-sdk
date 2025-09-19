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
	// The IDs of alert events. You can specify up to 100 IDs. If you do not specify this parameter, the value of the response parameter **Data*	- is **false**. The value false indicates that no alert events are generated on the server.
	//
	// > You can call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to query the IDs of alert events.
	//
	// example:
	//
	// ["1234567","98765432"]
	SecurityEventIds []*string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty" type:"Repeated"`
	// The UUID of the server.
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
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
