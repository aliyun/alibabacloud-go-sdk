// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossCloudRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudProvider(v string) *DescribeCrossCloudRegionRequest
	GetCloudProvider() *string
	SetCrossCloudRegionId(v string) *DescribeCrossCloudRegionRequest
	GetCrossCloudRegionId() *string
	SetDBType(v string) *DescribeCrossCloudRegionRequest
	GetDBType() *string
}

type DescribeCrossCloudRegionRequest struct {
	// example:
	//
	// HuaweiCloud
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// example:
	//
	// cn-east-3
	CrossCloudRegionId *string `json:"CrossCloudRegionId,omitempty" xml:"CrossCloudRegionId,omitempty"`
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
}

func (s DescribeCrossCloudRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossCloudRegionRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrossCloudRegionRequest) GetCloudProvider() *string {
	return s.CloudProvider
}

func (s *DescribeCrossCloudRegionRequest) GetCrossCloudRegionId() *string {
	return s.CrossCloudRegionId
}

func (s *DescribeCrossCloudRegionRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeCrossCloudRegionRequest) SetCloudProvider(v string) *DescribeCrossCloudRegionRequest {
	s.CloudProvider = &v
	return s
}

func (s *DescribeCrossCloudRegionRequest) SetCrossCloudRegionId(v string) *DescribeCrossCloudRegionRequest {
	s.CrossCloudRegionId = &v
	return s
}

func (s *DescribeCrossCloudRegionRequest) SetDBType(v string) *DescribeCrossCloudRegionRequest {
	s.DBType = &v
	return s
}

func (s *DescribeCrossCloudRegionRequest) Validate() error {
	return dara.Validate(s)
}
