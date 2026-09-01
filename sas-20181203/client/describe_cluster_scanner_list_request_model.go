// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterScannerListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIdList(v []*string) *DescribeClusterScannerListRequest
	GetClusterIdList() []*string
	SetLang(v string) *DescribeClusterScannerListRequest
	GetLang() *string
	SetStatusList(v []*string) *DescribeClusterScannerListRequest
	GetStatusList() []*string
}

type DescribeClusterScannerListRequest struct {
	// The list of cluster IDs.
	ClusterIdList []*string `json:"ClusterIdList,omitempty" xml:"ClusterIdList,omitempty" type:"Repeated"`
	// The language type for the request and response messages.
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The list of scanner statuses. Valid values:
	//
	// - **online**: running
	//
	// - **offline**: offline
	//
	// - **not_installed**: not installed
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s DescribeClusterScannerListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterScannerListRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterScannerListRequest) GetClusterIdList() []*string {
	return s.ClusterIdList
}

func (s *DescribeClusterScannerListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeClusterScannerListRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *DescribeClusterScannerListRequest) SetClusterIdList(v []*string) *DescribeClusterScannerListRequest {
	s.ClusterIdList = v
	return s
}

func (s *DescribeClusterScannerListRequest) SetLang(v string) *DescribeClusterScannerListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeClusterScannerListRequest) SetStatusList(v []*string) *DescribeClusterScannerListRequest {
	s.StatusList = v
	return s
}

func (s *DescribeClusterScannerListRequest) Validate() error {
	return dara.Validate(s)
}
