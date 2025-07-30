// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMfaDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdDomain(v string) *DescribeMfaDevicesRequest
	GetAdDomain() *string
	SetEndUserIds(v []*string) *DescribeMfaDevicesRequest
	GetEndUserIds() []*string
	SetMaxResults(v int64) *DescribeMfaDevicesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeMfaDevicesRequest
	GetNextToken() *string
	SetSerialNumbers(v []*string) *DescribeMfaDevicesRequest
	GetSerialNumbers() []*string
}

type DescribeMfaDevicesRequest struct {
	// The domain of the Active Directory (AD) workspace.
	//
	// example:
	//
	// cn.misumi.pri
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The usernames of the convenience users.
	//
	// example:
	//
	// test
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return. Valid values: 1 to 500.\\
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Set the value to the token that is obtained from the previous query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The serial numbers of the virtual MFA devices.
	//
	// example:
	//
	// c2d9ae94-a64b-4a0d-8024-9519ca50****
	SerialNumbers []*string `json:"SerialNumbers,omitempty" xml:"SerialNumbers,omitempty" type:"Repeated"`
}

func (s DescribeMfaDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMfaDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMfaDevicesRequest) GetAdDomain() *string {
	return s.AdDomain
}

func (s *DescribeMfaDevicesRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribeMfaDevicesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeMfaDevicesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMfaDevicesRequest) GetSerialNumbers() []*string {
	return s.SerialNumbers
}

func (s *DescribeMfaDevicesRequest) SetAdDomain(v string) *DescribeMfaDevicesRequest {
	s.AdDomain = &v
	return s
}

func (s *DescribeMfaDevicesRequest) SetEndUserIds(v []*string) *DescribeMfaDevicesRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribeMfaDevicesRequest) SetMaxResults(v int64) *DescribeMfaDevicesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeMfaDevicesRequest) SetNextToken(v string) *DescribeMfaDevicesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMfaDevicesRequest) SetSerialNumbers(v []*string) *DescribeMfaDevicesRequest {
	s.SerialNumbers = v
	return s
}

func (s *DescribeMfaDevicesRequest) Validate() error {
	return dara.Validate(s)
}
