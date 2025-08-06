// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodAppNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeVodAppNameResponseBodyData) *DescribeVodAppNameResponseBody
	GetData() []*DescribeVodAppNameResponseBodyData
	SetRequestId(v string) *DescribeVodAppNameResponseBody
	GetRequestId() *string
}

type DescribeVodAppNameResponseBody struct {
	Data      []*DescribeVodAppNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodAppNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodAppNameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodAppNameResponseBody) GetData() []*DescribeVodAppNameResponseBodyData {
	return s.Data
}

func (s *DescribeVodAppNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodAppNameResponseBody) SetData(v []*DescribeVodAppNameResponseBodyData) *DescribeVodAppNameResponseBody {
	s.Data = v
	return s
}

func (s *DescribeVodAppNameResponseBody) SetRequestId(v string) *DescribeVodAppNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodAppNameResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodAppNameResponseBodyData struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s DescribeVodAppNameResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodAppNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeVodAppNameResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeVodAppNameResponseBodyData) SetAppName(v string) *DescribeVodAppNameResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeVodAppNameResponseBodyData) Validate() error {
	return dara.Validate(s)
}
