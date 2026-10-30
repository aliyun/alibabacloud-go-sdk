// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKeyPairsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairId(v string) *DescribeKeyPairsRequest
	GetKeyPairId() *string
	SetKeyPairName(v string) *DescribeKeyPairsRequest
	GetKeyPairName() *string
	SetPageNumber(v string) *DescribeKeyPairsRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeKeyPairsRequest
	GetPageSize() *string
}

type DescribeKeyPairsRequest struct {
	// The ID of the key pair.
	//
	// example:
	//
	// ssh-50cynkq42sgj4ej1tn78t4***
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the key pair. The name must be 2 to 128 characters in length and can contain letters and Chinese characters. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. It can contain the following characters:
	//
	// - Digits
	//
	// - :
	//
	// - _
	//
	// - .
	//
	// Only a single name is supported. By default, information about all key pairs is queried.
	//
	// example:
	//
	// TestKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The page number. Valid values: integers greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The paging Settings for the number of entries per page in a paged query. Valid values: integers greater than 0. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeKeyPairsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsRequest) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *DescribeKeyPairsRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeKeyPairsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeKeyPairsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeKeyPairsRequest) SetKeyPairId(v string) *DescribeKeyPairsRequest {
	s.KeyPairId = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetKeyPairName(v string) *DescribeKeyPairsRequest {
	s.KeyPairName = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetPageNumber(v string) *DescribeKeyPairsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetPageSize(v string) *DescribeKeyPairsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeKeyPairsRequest) Validate() error {
	return dara.Validate(s)
}
