// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseS3AccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPfsInstanceId(v string) *DescribeLakebaseS3AccountRequest
	GetPfsInstanceId() *string
	SetRegionId(v string) *DescribeLakebaseS3AccountRequest
	GetRegionId() *string
	SetUserAccAk(v string) *DescribeLakebaseS3AccountRequest
	GetUserAccAk() *string
}

type DescribeLakebaseS3AccountRequest struct {
	// The PolarFS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-xxx
	PfsInstanceId *string `json:"PfsInstanceId,omitempty" xml:"PfsInstanceId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query all available regions for your account, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Access Key of the S3 account.
	//
	// > The account name supports only uppercase letters, lowercase letters, and digits, with a maximum length of 32 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// accname
	UserAccAk *string `json:"UserAccAk,omitempty" xml:"UserAccAk,omitempty"`
}

func (s DescribeLakebaseS3AccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseS3AccountRequest) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseS3AccountRequest) GetPfsInstanceId() *string {
	return s.PfsInstanceId
}

func (s *DescribeLakebaseS3AccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLakebaseS3AccountRequest) GetUserAccAk() *string {
	return s.UserAccAk
}

func (s *DescribeLakebaseS3AccountRequest) SetPfsInstanceId(v string) *DescribeLakebaseS3AccountRequest {
	s.PfsInstanceId = &v
	return s
}

func (s *DescribeLakebaseS3AccountRequest) SetRegionId(v string) *DescribeLakebaseS3AccountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLakebaseS3AccountRequest) SetUserAccAk(v string) *DescribeLakebaseS3AccountRequest {
	s.UserAccAk = &v
	return s
}

func (s *DescribeLakebaseS3AccountRequest) Validate() error {
	return dara.Validate(s)
}
