// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadInstanceCACertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DownloadInstanceCACertificateRequest
	GetDBClusterId() *string
	SetEngine(v string) *DownloadInstanceCACertificateRequest
	GetEngine() *string
	SetOwnerId(v string) *DownloadInstanceCACertificateRequest
	GetOwnerId() *string
	SetRegionId(v string) *DownloadInstanceCACertificateRequest
	GetRegionId() *string
}

type DownloadInstanceCACertificateRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9842849v6****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine of the cluster. Valid values:
	//
	// 	- **AnalyticDB*	- (default): the AnalyticDB for MySQL engine.
	//
	// 	- **Clickhouse**: the wide table engine.
	//
	// example:
	//
	// Clickhouse
	Engine  *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DownloadInstanceCACertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadInstanceCACertificateRequest) GoString() string {
	return s.String()
}

func (s *DownloadInstanceCACertificateRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DownloadInstanceCACertificateRequest) GetEngine() *string {
	return s.Engine
}

func (s *DownloadInstanceCACertificateRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DownloadInstanceCACertificateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DownloadInstanceCACertificateRequest) SetDBClusterId(v string) *DownloadInstanceCACertificateRequest {
	s.DBClusterId = &v
	return s
}

func (s *DownloadInstanceCACertificateRequest) SetEngine(v string) *DownloadInstanceCACertificateRequest {
	s.Engine = &v
	return s
}

func (s *DownloadInstanceCACertificateRequest) SetOwnerId(v string) *DownloadInstanceCACertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *DownloadInstanceCACertificateRequest) SetRegionId(v string) *DownloadInstanceCACertificateRequest {
	s.RegionId = &v
	return s
}

func (s *DownloadInstanceCACertificateRequest) Validate() error {
	return dara.Validate(s)
}
