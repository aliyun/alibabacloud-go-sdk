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
}

type GetOssBucketScanStatisticRequest struct {
	BucketNameList []*string `json:"BucketNameList,omitempty" xml:"BucketNameList,omitempty" type:"Repeated"`
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

func (s *GetOssBucketScanStatisticRequest) SetBucketNameList(v []*string) *GetOssBucketScanStatisticRequest {
	s.BucketNameList = v
	return s
}

func (s *GetOssBucketScanStatisticRequest) Validate() error {
	return dara.Validate(s)
}
