// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerDimensionDataDemoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*string) *DescribeVodPlayerDimensionDataDemoResponseBody
	GetDataList() []*string
	SetRequestId(v string) *DescribeVodPlayerDimensionDataDemoResponseBody
	GetRequestId() *string
}

type DescribeVodPlayerDimensionDataDemoResponseBody struct {
	DataList []*string `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodPlayerDimensionDataDemoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerDimensionDataDemoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerDimensionDataDemoResponseBody) GetDataList() []*string {
	return s.DataList
}

func (s *DescribeVodPlayerDimensionDataDemoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodPlayerDimensionDataDemoResponseBody) SetDataList(v []*string) *DescribeVodPlayerDimensionDataDemoResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeVodPlayerDimensionDataDemoResponseBody) SetRequestId(v string) *DescribeVodPlayerDimensionDataDemoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodPlayerDimensionDataDemoResponseBody) Validate() error {
	return dara.Validate(s)
}
