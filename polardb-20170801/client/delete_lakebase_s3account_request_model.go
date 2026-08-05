// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLakebaseS3AccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPfsInstanceId(v string) *DeleteLakebaseS3AccountRequest
	GetPfsInstanceId() *string
	SetRegionId(v string) *DeleteLakebaseS3AccountRequest
	GetRegionId() *string
	SetUserAccAk(v string) *DeleteLakebaseS3AccountRequest
	GetUserAccAk() *string
}

type DeleteLakebaseS3AccountRequest struct {
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
	// >You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The access key of the S3 account to delete.
	//
	// > The default account cannot be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// accname
	UserAccAk *string `json:"UserAccAk,omitempty" xml:"UserAccAk,omitempty"`
}

func (s DeleteLakebaseS3AccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLakebaseS3AccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteLakebaseS3AccountRequest) GetPfsInstanceId() *string {
	return s.PfsInstanceId
}

func (s *DeleteLakebaseS3AccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLakebaseS3AccountRequest) GetUserAccAk() *string {
	return s.UserAccAk
}

func (s *DeleteLakebaseS3AccountRequest) SetPfsInstanceId(v string) *DeleteLakebaseS3AccountRequest {
	s.PfsInstanceId = &v
	return s
}

func (s *DeleteLakebaseS3AccountRequest) SetRegionId(v string) *DeleteLakebaseS3AccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLakebaseS3AccountRequest) SetUserAccAk(v string) *DeleteLakebaseS3AccountRequest {
	s.UserAccAk = &v
	return s
}

func (s *DeleteLakebaseS3AccountRequest) Validate() error {
	return dara.Validate(s)
}
