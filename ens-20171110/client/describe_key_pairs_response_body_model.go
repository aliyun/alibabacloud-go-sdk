// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKeyPairsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairs(v *DescribeKeyPairsResponseBodyKeyPairs) *DescribeKeyPairsResponseBody
	GetKeyPairs() *DescribeKeyPairsResponseBodyKeyPairs
	SetPageNumber(v int32) *DescribeKeyPairsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeKeyPairsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeKeyPairsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeKeyPairsResponseBody
	GetTotalCount() *int32
}

type DescribeKeyPairsResponseBody struct {
	// Details about the key pairs.
	KeyPairs *DescribeKeyPairsResponseBodyKeyPairs `json:"KeyPairs,omitempty" xml:"KeyPairs,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 37B52F33-6879-49D0-A39B-22966B01449E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of key pairs.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeKeyPairsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBody) GetKeyPairs() *DescribeKeyPairsResponseBodyKeyPairs {
	return s.KeyPairs
}

func (s *DescribeKeyPairsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeKeyPairsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeKeyPairsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKeyPairsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeKeyPairsResponseBody) SetKeyPairs(v *DescribeKeyPairsResponseBodyKeyPairs) *DescribeKeyPairsResponseBody {
	s.KeyPairs = v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetPageNumber(v int32) *DescribeKeyPairsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetPageSize(v int32) *DescribeKeyPairsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetRequestId(v string) *DescribeKeyPairsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) SetTotalCount(v int32) *DescribeKeyPairsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeKeyPairsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeKeyPairsResponseBodyKeyPairs struct {
	KeyPair []*DescribeKeyPairsResponseBodyKeyPairsKeyPair `json:"KeyPair,omitempty" xml:"KeyPair,omitempty" type:"Repeated"`
}

func (s DescribeKeyPairsResponseBodyKeyPairs) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsResponseBodyKeyPairs) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBodyKeyPairs) GetKeyPair() []*DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	return s.KeyPair
}

func (s *DescribeKeyPairsResponseBodyKeyPairs) SetKeyPair(v []*DescribeKeyPairsResponseBodyKeyPairsKeyPair) *DescribeKeyPairsResponseBodyKeyPairs {
	s.KeyPair = v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairs) Validate() error {
	return dara.Validate(s)
}

type DescribeKeyPairsResponseBodyKeyPairsKeyPair struct {
	// The time when the key pair was created. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-04-26T15:38:27Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The fingerprint of the key pair.
	//
	// example:
	//
	// fdaf8ff7a756ef843814fc****
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty"`
	// The ID of the key pair.
	//
	// example:
	//
	// ssh-50cynkq42sgj4ej1tn78t4***
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the SSH key pair.
	//
	// example:
	//
	// TestKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
}

func (s DescribeKeyPairsResponseBodyKeyPairsKeyPair) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsResponseBodyKeyPairsKeyPair) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) GetKeyPairFingerPrint() *string {
	return s.KeyPairFingerPrint
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) SetCreationTime(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	s.CreationTime = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) SetKeyPairFingerPrint(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	s.KeyPairFingerPrint = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) SetKeyPairId(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	s.KeyPairId = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) SetKeyPairName(v string) *DescribeKeyPairsResponseBodyKeyPairsKeyPair {
	s.KeyPairName = &v
	return s
}

func (s *DescribeKeyPairsResponseBodyKeyPairsKeyPair) Validate() error {
	return dara.Validate(s)
}
