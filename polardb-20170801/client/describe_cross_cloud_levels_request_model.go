// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossCloudLevelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBType(v string) *DescribeCrossCloudLevelsRequest
	GetDBType() *string
	SetDBVersion(v string) *DescribeCrossCloudLevelsRequest
	GetDBVersion() *string
	SetProjectId(v string) *DescribeCrossCloudLevelsRequest
	GetProjectId() *string
	SetStorageType(v string) *DescribeCrossCloudLevelsRequest
	GetStorageType() *string
}

type DescribeCrossCloudLevelsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pj-87681rbcef6******
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SharedStorage
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeCrossCloudLevelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossCloudLevelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrossCloudLevelsRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeCrossCloudLevelsRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeCrossCloudLevelsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeCrossCloudLevelsRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeCrossCloudLevelsRequest) SetDBType(v string) *DescribeCrossCloudLevelsRequest {
	s.DBType = &v
	return s
}

func (s *DescribeCrossCloudLevelsRequest) SetDBVersion(v string) *DescribeCrossCloudLevelsRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeCrossCloudLevelsRequest) SetProjectId(v string) *DescribeCrossCloudLevelsRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeCrossCloudLevelsRequest) SetStorageType(v string) *DescribeCrossCloudLevelsRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeCrossCloudLevelsRequest) Validate() error {
	return dara.Validate(s)
}
