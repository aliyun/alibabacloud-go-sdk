// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceBelongIdentifiersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifiers(v []*ListDeviceBelongIdentifiersResponseBodyIdentifiers) *ListDeviceBelongIdentifiersResponseBody
	GetIdentifiers() []*ListDeviceBelongIdentifiersResponseBodyIdentifiers
	SetMaxResults(v int32) *ListDeviceBelongIdentifiersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDeviceBelongIdentifiersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDeviceBelongIdentifiersResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDeviceBelongIdentifiersResponseBody
	GetTotalCount() *int64
}

type ListDeviceBelongIdentifiersResponseBody struct {
	// The device ownership identifier records on the current page. An empty list is returned if no matching records are found.
	Identifiers []*ListDeviceBelongIdentifiersResponseBodyIdentifiers `json:"Identifiers,omitempty" xml:"Identifiers,omitempty" type:"Repeated"`
	// A reserved pagination field. No valid value is returned in the current version.
	//
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A reserved pagination cursor. No valid value is returned in the current version.
	//
	// example:
	//
	// d9fbd0d1ac24be90
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7A92B549-EC82-5F79-8301-63711AFC9DCE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that match the filter conditions for the current tenant.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeviceBelongIdentifiersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceBelongIdentifiersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceBelongIdentifiersResponseBody) GetIdentifiers() []*ListDeviceBelongIdentifiersResponseBodyIdentifiers {
	return s.Identifiers
}

func (s *ListDeviceBelongIdentifiersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDeviceBelongIdentifiersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDeviceBelongIdentifiersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceBelongIdentifiersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDeviceBelongIdentifiersResponseBody) SetIdentifiers(v []*ListDeviceBelongIdentifiersResponseBodyIdentifiers) *ListDeviceBelongIdentifiersResponseBody {
	s.Identifiers = v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBody) SetMaxResults(v int32) *ListDeviceBelongIdentifiersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBody) SetNextToken(v string) *ListDeviceBelongIdentifiersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBody) SetRequestId(v string) *ListDeviceBelongIdentifiersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBody) SetTotalCount(v int64) *ListDeviceBelongIdentifiersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBody) Validate() error {
	if s.Identifiers != nil {
		for _, item := range s.Identifiers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeviceBelongIdentifiersResponseBodyIdentifiers struct {
	// The time when the record was created. The time is in RFC 3339 format and includes time zone information.
	//
	// example:
	//
	// 2026-11-24T10:30:00+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the identifier record, returned as a decimal string. You can use the ID to query or remove the record.
	//
	// example:
	//
	// 482731
	IdentifierId *string `json:"IdentifierId,omitempty" xml:"IdentifierId,omitempty"`
	// The identifier type. MAC indicates the MAC address of the network interface controller (NIC), SystemSN indicates the system serial number, and BIOSSN indicates the BIOS serial number.
	//
	// example:
	//
	// MAC
	IdentifierType *string `json:"IdentifierType,omitempty" xml:"IdentifierType,omitempty"`
	// The identifier value in the record. MAC addresses are in the lowercase colon-separated format. System serial numbers and BIOS serial numbers are returned as recorded.
	//
	// example:
	//
	// 02:00:00:00:00:01
	IdentifierValue *string `json:"IdentifierValue,omitempty" xml:"IdentifierValue,omitempty"`
}

func (s ListDeviceBelongIdentifiersResponseBodyIdentifiers) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceBelongIdentifiersResponseBodyIdentifiers) GoString() string {
	return s.String()
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) GetIdentifierId() *string {
	return s.IdentifierId
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) GetIdentifierType() *string {
	return s.IdentifierType
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) GetIdentifierValue() *string {
	return s.IdentifierValue
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) SetCreateTime(v string) *ListDeviceBelongIdentifiersResponseBodyIdentifiers {
	s.CreateTime = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) SetIdentifierId(v string) *ListDeviceBelongIdentifiersResponseBodyIdentifiers {
	s.IdentifierId = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) SetIdentifierType(v string) *ListDeviceBelongIdentifiersResponseBodyIdentifiers {
	s.IdentifierType = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) SetIdentifierValue(v string) *ListDeviceBelongIdentifiersResponseBodyIdentifiers {
	s.IdentifierValue = &v
	return s
}

func (s *ListDeviceBelongIdentifiersResponseBodyIdentifiers) Validate() error {
	return dara.Validate(s)
}
