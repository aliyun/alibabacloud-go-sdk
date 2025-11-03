// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedCount(v int32) *CreateRouteEntriesResponseBody
	GetFailedCount() *int32
	SetFailedRouteEntries(v []*CreateRouteEntriesResponseBodyFailedRouteEntries) *CreateRouteEntriesResponseBody
	GetFailedRouteEntries() []*CreateRouteEntriesResponseBodyFailedRouteEntries
	SetRequestId(v string) *CreateRouteEntriesResponseBody
	GetRequestId() *string
	SetRouteEntryIds(v []*string) *CreateRouteEntriesResponseBody
	GetRouteEntryIds() []*string
	SetSuccessCount(v int32) *CreateRouteEntriesResponseBody
	GetSuccessCount() *int32
}

type CreateRouteEntriesResponseBody struct {
	// The number of custom route entries that failed to be added.
	//
	// example:
	//
	// 2
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The details about the custom route entry that failed to be added.
	FailedRouteEntries []*CreateRouteEntriesResponseBodyFailedRouteEntries `json:"FailedRouteEntries,omitempty" xml:"FailedRouteEntries,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the ID of the custom route entry that was successfully added.
	RouteEntryIds []*string `json:"RouteEntryIds,omitempty" xml:"RouteEntryIds,omitempty" type:"Repeated"`
	// The number of custom route entries that were successfully added.
	//
	// example:
	//
	// 2
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s CreateRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRouteEntriesResponseBody) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *CreateRouteEntriesResponseBody) GetFailedRouteEntries() []*CreateRouteEntriesResponseBodyFailedRouteEntries {
	return s.FailedRouteEntries
}

func (s *CreateRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRouteEntriesResponseBody) GetRouteEntryIds() []*string {
	return s.RouteEntryIds
}

func (s *CreateRouteEntriesResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *CreateRouteEntriesResponseBody) SetFailedCount(v int32) *CreateRouteEntriesResponseBody {
	s.FailedCount = &v
	return s
}

func (s *CreateRouteEntriesResponseBody) SetFailedRouteEntries(v []*CreateRouteEntriesResponseBodyFailedRouteEntries) *CreateRouteEntriesResponseBody {
	s.FailedRouteEntries = v
	return s
}

func (s *CreateRouteEntriesResponseBody) SetRequestId(v string) *CreateRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRouteEntriesResponseBody) SetRouteEntryIds(v []*string) *CreateRouteEntriesResponseBody {
	s.RouteEntryIds = v
	return s
}

func (s *CreateRouteEntriesResponseBody) SetSuccessCount(v int32) *CreateRouteEntriesResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *CreateRouteEntriesResponseBody) Validate() error {
	if s.FailedRouteEntries != nil {
		for _, item := range s.FailedRouteEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRouteEntriesResponseBodyFailedRouteEntries struct {
	// The destination CIDR block of the custom route entry that failed to be added.
	//
	// example:
	//
	// 192.168.0.0/24
	DstCidrBlock *string `json:"DstCidrBlock,omitempty" xml:"DstCidrBlock,omitempty"`
	// The error code.
	//
	// example:
	//
	// VPC_ROUTE_ENTRY_CIDR_BLOCK_DUPLICATE
	FailedCode *string `json:"FailedCode,omitempty" xml:"FailedCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Specified CIDR block is already exists, entry.cidrBlock=xxxx
	FailedMessage *string `json:"FailedMessage,omitempty" xml:"FailedMessage,omitempty"`
	// The ID of the next hop of the custom route entry that failed to be added.
	//
	// example:
	//
	// i-j6c2fp57q8rr4jlu****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
}

func (s CreateRouteEntriesResponseBodyFailedRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteEntriesResponseBodyFailedRouteEntries) GoString() string {
	return s.String()
}

func (s *CreateRouteEntriesResponseBodyFailedRouteEntries) GetDstCidrBlock() *string {
	return s.DstCidrBlock
}

func (s *CreateRouteEntriesResponseBodyFailedRouteEntries) GetFailedCode() *string {
	return s.FailedCode
}

func (s *CreateRouteEntriesResponseBodyFailedRouteEntries) GetFailedMessage() *string {
	return s.FailedMessage
}

func (s *CreateRouteEntriesResponseBodyFailedRouteEntries) GetNextHop() *string {
	return s.NextHop
}

func (s *CreateRouteEntriesResponseBodyFailedRouteEntries) SetDstCidrBlock(v string) *CreateRouteEntriesResponseBodyFailedRouteEntries {
	s.DstCidrBlock = &v
	return s
}

func (s *CreateRouteEntriesResponseBodyFailedRouteEntries) SetFailedCode(v string) *CreateRouteEntriesResponseBodyFailedRouteEntries {
	s.FailedCode = &v
	return s
}

func (s *CreateRouteEntriesResponseBodyFailedRouteEntries) SetFailedMessage(v string) *CreateRouteEntriesResponseBodyFailedRouteEntries {
	s.FailedMessage = &v
	return s
}

func (s *CreateRouteEntriesResponseBodyFailedRouteEntries) SetNextHop(v string) *CreateRouteEntriesResponseBodyFailedRouteEntries {
	s.NextHop = &v
	return s
}

func (s *CreateRouteEntriesResponseBodyFailedRouteEntries) Validate() error {
	return dara.Validate(s)
}
