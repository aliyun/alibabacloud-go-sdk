// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceBelongIdentifiersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListDeviceBelongIdentifiersRequest
	GetCurrentPage() *int32
	SetIdentifierIds(v []*string) *ListDeviceBelongIdentifiersRequest
	GetIdentifierIds() []*string
	SetIdentifierType(v string) *ListDeviceBelongIdentifiersRequest
	GetIdentifierType() *string
	SetIdentifierValues(v []*string) *ListDeviceBelongIdentifiersRequest
	GetIdentifierValues() []*string
	SetMaxResults(v int32) *ListDeviceBelongIdentifiersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDeviceBelongIdentifiersRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListDeviceBelongIdentifiersRequest
	GetPageSize() *int32
}

type ListDeviceBelongIdentifiersRequest struct {
	// The current page number. Valid values: 1 to 2147483647.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of record IDs to query. You can specify 1 to 100 positive integer decimal strings at a time. If this parameter is specified along with other filter parameters, the records that meet all the conditions are returned.
	IdentifierIds []*string `json:"IdentifierIds,omitempty" xml:"IdentifierIds,omitempty" type:"Repeated"`
	// The identifier type. MAC indicates the MAC address of the network interface controller (NIC), SystemSN indicates the system serial number, and BIOSSN indicates the BIOS serial number. This parameter is required when IdentifierValues is specified.
	//
	// example:
	//
	// MAC
	IdentifierType *string `json:"IdentifierType,omitempty" xml:"IdentifierType,omitempty"`
	// The list of identifier values for exact match. You can specify 1 to 100 items at a time, and each item can contain 1 to 128 characters. You must specify IdentifierType at the same time. Records that match any value in the array are returned. MAC addresses are uniformly converted to the lowercase colon-separated format. System serial numbers and BIOS serial numbers are case-sensitive and space-sensitive.
	IdentifierValues []*string `json:"IdentifierValues,omitempty" xml:"IdentifierValues,omitempty" type:"Repeated"`
	// A reserved pagination parameter. This value is not used in the current version. Use PageSize to specify the number of entries per page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A reserved pagination cursor. This value is not used in the current version. Use CurrentPage to specify the page number.
	//
	// example:
	//
	// d9fbd0d1ac24be90
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The maximum number of records to return per page. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDeviceBelongIdentifiersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceBelongIdentifiersRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceBelongIdentifiersRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDeviceBelongIdentifiersRequest) GetIdentifierIds() []*string {
	return s.IdentifierIds
}

func (s *ListDeviceBelongIdentifiersRequest) GetIdentifierType() *string {
	return s.IdentifierType
}

func (s *ListDeviceBelongIdentifiersRequest) GetIdentifierValues() []*string {
	return s.IdentifierValues
}

func (s *ListDeviceBelongIdentifiersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDeviceBelongIdentifiersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDeviceBelongIdentifiersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeviceBelongIdentifiersRequest) SetCurrentPage(v int32) *ListDeviceBelongIdentifiersRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDeviceBelongIdentifiersRequest) SetIdentifierIds(v []*string) *ListDeviceBelongIdentifiersRequest {
	s.IdentifierIds = v
	return s
}

func (s *ListDeviceBelongIdentifiersRequest) SetIdentifierType(v string) *ListDeviceBelongIdentifiersRequest {
	s.IdentifierType = &v
	return s
}

func (s *ListDeviceBelongIdentifiersRequest) SetIdentifierValues(v []*string) *ListDeviceBelongIdentifiersRequest {
	s.IdentifierValues = v
	return s
}

func (s *ListDeviceBelongIdentifiersRequest) SetMaxResults(v int32) *ListDeviceBelongIdentifiersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDeviceBelongIdentifiersRequest) SetNextToken(v string) *ListDeviceBelongIdentifiersRequest {
	s.NextToken = &v
	return s
}

func (s *ListDeviceBelongIdentifiersRequest) SetPageSize(v int32) *ListDeviceBelongIdentifiersRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeviceBelongIdentifiersRequest) Validate() error {
	return dara.Validate(s)
}
