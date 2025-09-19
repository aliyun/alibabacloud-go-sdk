// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssBucketScanStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketNameList(v []*string) *GetOssBucketScanStatisticRequest
	GetBucketNameList() []*string
	SetSource(v string) *GetOssBucketScanStatisticRequest
	GetSource() *string
}

type GetOssBucketScanStatisticRequest struct {
	// The names of the buckets.
	BucketNameList []*string `json:"BucketNameList,omitempty" xml:"BucketNameList,omitempty" type:"Repeated"`
	// The data source. Valid values:
	//
	// 	- **API**: API operations.
	//
	// 	- **OSS**: Object Storage Service (OSS) file check.
	//
	// example:
	//
	// API
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s GetOssBucketScanStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssBucketScanStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetOssBucketScanStatisticRequest) GetBucketNameList() []*string {
	return s.BucketNameList
}

func (s *GetOssBucketScanStatisticRequest) GetSource() *string {
	return s.Source
}

func (s *GetOssBucketScanStatisticRequest) SetBucketNameList(v []*string) *GetOssBucketScanStatisticRequest {
	s.BucketNameList = v
	return s
}

func (s *GetOssBucketScanStatisticRequest) SetSource(v string) *GetOssBucketScanStatisticRequest {
	s.Source = &v
	return s
}

func (s *GetOssBucketScanStatisticRequest) Validate() error {
	return dara.Validate(s)
}
