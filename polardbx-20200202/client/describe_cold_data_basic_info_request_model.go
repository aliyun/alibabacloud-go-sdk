// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColdDataBasicInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeColdDataBasicInfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeColdDataBasicInfoRequest
	GetRegionId() *string
}

type DescribeColdDataBasicInfoRequest struct {
	// This parameter is required.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeColdDataBasicInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdDataBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeColdDataBasicInfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeColdDataBasicInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeColdDataBasicInfoRequest) SetDBInstanceName(v string) *DescribeColdDataBasicInfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeColdDataBasicInfoRequest) SetRegionId(v string) *DescribeColdDataBasicInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeColdDataBasicInfoRequest) Validate() error {
	return dara.Validate(s)
}
