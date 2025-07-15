// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedCount(v int32) *DeleteRouteEntriesResponseBody
	GetFailedCount() *int32
	SetFailedRouteEntries(v []*DeleteRouteEntriesResponseBodyFailedRouteEntries) *DeleteRouteEntriesResponseBody
	GetFailedRouteEntries() []*DeleteRouteEntriesResponseBodyFailedRouteEntries
	SetRequestId(v string) *DeleteRouteEntriesResponseBody
	GetRequestId() *string
	SetSuccessCount(v int32) *DeleteRouteEntriesResponseBody
	GetSuccessCount() *int32
}

type DeleteRouteEntriesResponseBody struct {
	// The number of route entries that failed to be deleted.
	//
	// example:
	//
	// 2
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The information about the route entry that failed to be deleted.
	FailedRouteEntries []*DeleteRouteEntriesResponseBodyFailedRouteEntries `json:"FailedRouteEntries,omitempty" xml:"FailedRouteEntries,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of route entries that were deleted.
	//
	// example:
	//
	// 2
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s DeleteRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteEntriesResponseBody) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *DeleteRouteEntriesResponseBody) GetFailedRouteEntries() []*DeleteRouteEntriesResponseBodyFailedRouteEntries {
	return s.FailedRouteEntries
}

func (s *DeleteRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRouteEntriesResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *DeleteRouteEntriesResponseBody) SetFailedCount(v int32) *DeleteRouteEntriesResponseBody {
	s.FailedCount = &v
	return s
}

func (s *DeleteRouteEntriesResponseBody) SetFailedRouteEntries(v []*DeleteRouteEntriesResponseBodyFailedRouteEntries) *DeleteRouteEntriesResponseBody {
	s.FailedRouteEntries = v
	return s
}

func (s *DeleteRouteEntriesResponseBody) SetRequestId(v string) *DeleteRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteEntriesResponseBody) SetSuccessCount(v int32) *DeleteRouteEntriesResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *DeleteRouteEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteRouteEntriesResponseBodyFailedRouteEntries struct {
	// The destination CIDR block of the route entry that failed to be deleted. IPv4 and IPv6 CIDR blocks are supported.
	//
	// example:
	//
	// 47.100.XX.XX/24
	DstCidrBlock *string `json:"DstCidrBlock,omitempty" xml:"DstCidrBlock,omitempty"`
	// The error code.
	//
	// example:
	//
	// VPC_ROUTER_ENTRY_NOT_EXIST
	FailedCode *string `json:"FailedCode,omitempty" xml:"FailedCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// vRouterEntry not exists
	FailedMessage *string `json:"FailedMessage,omitempty" xml:"FailedMessage,omitempty"`
	// The ID of the next hop that failed to be deleted.
	//
	// example:
	//
	// i-j6c2fp57q8rr4jlu****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The ID of the route entry that failed to be deleted.
	//
	// example:
	//
	// rte-bp1mnnr2al0naomnpv****
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
}

func (s DeleteRouteEntriesResponseBodyFailedRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteEntriesResponseBodyFailedRouteEntries) GoString() string {
	return s.String()
}

func (s *DeleteRouteEntriesResponseBodyFailedRouteEntries) GetDstCidrBlock() *string {
	return s.DstCidrBlock
}

func (s *DeleteRouteEntriesResponseBodyFailedRouteEntries) GetFailedCode() *string {
	return s.FailedCode
}

func (s *DeleteRouteEntriesResponseBodyFailedRouteEntries) GetFailedMessage() *string {
	return s.FailedMessage
}

func (s *DeleteRouteEntriesResponseBodyFailedRouteEntries) GetNextHop() *string {
	return s.NextHop
}

func (s *DeleteRouteEntriesResponseBodyFailedRouteEntries) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *DeleteRouteEntriesResponseBodyFailedRouteEntries) SetDstCidrBlock(v string) *DeleteRouteEntriesResponseBodyFailedRouteEntries {
	s.DstCidrBlock = &v
	return s
}

func (s *DeleteRouteEntriesResponseBodyFailedRouteEntries) SetFailedCode(v string) *DeleteRouteEntriesResponseBodyFailedRouteEntries {
	s.FailedCode = &v
	return s
}

func (s *DeleteRouteEntriesResponseBodyFailedRouteEntries) SetFailedMessage(v string) *DeleteRouteEntriesResponseBodyFailedRouteEntries {
	s.FailedMessage = &v
	return s
}

func (s *DeleteRouteEntriesResponseBodyFailedRouteEntries) SetNextHop(v string) *DeleteRouteEntriesResponseBodyFailedRouteEntries {
	s.NextHop = &v
	return s
}

func (s *DeleteRouteEntriesResponseBodyFailedRouteEntries) SetRouteEntryId(v string) *DeleteRouteEntriesResponseBodyFailedRouteEntries {
	s.RouteEntryId = &v
	return s
}

func (s *DeleteRouteEntriesResponseBodyFailedRouteEntries) Validate() error {
	return dara.Validate(s)
}
