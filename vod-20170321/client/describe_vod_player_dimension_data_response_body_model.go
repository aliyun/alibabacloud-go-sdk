// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerDimensionDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*string) *DescribeVodPlayerDimensionDataResponseBody
	GetDataList() []*string
	SetRequestId(v string) *DescribeVodPlayerDimensionDataResponseBody
	GetRequestId() *string
}

type DescribeVodPlayerDimensionDataResponseBody struct {
	DataList []*string `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodPlayerDimensionDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerDimensionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerDimensionDataResponseBody) GetDataList() []*string {
	return s.DataList
}

func (s *DescribeVodPlayerDimensionDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodPlayerDimensionDataResponseBody) SetDataList(v []*string) *DescribeVodPlayerDimensionDataResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeVodPlayerDimensionDataResponseBody) SetRequestId(v string) *DescribeVodPlayerDimensionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodPlayerDimensionDataResponseBody) Validate() error {
	return dara.Validate(s)
}
