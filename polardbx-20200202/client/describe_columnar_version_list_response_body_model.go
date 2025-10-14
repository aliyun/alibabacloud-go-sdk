// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnarVersionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeColumnarVersionListResponseBodyData) *DescribeColumnarVersionListResponseBody
	GetData() *DescribeColumnarVersionListResponseBodyData
	SetRequestId(v string) *DescribeColumnarVersionListResponseBody
	GetRequestId() *string
}

type DescribeColumnarVersionListResponseBody struct {
	Data *DescribeColumnarVersionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-****-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeColumnarVersionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarVersionListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColumnarVersionListResponseBody) GetData() *DescribeColumnarVersionListResponseBodyData {
	return s.Data
}

func (s *DescribeColumnarVersionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeColumnarVersionListResponseBody) SetData(v *DescribeColumnarVersionListResponseBodyData) *DescribeColumnarVersionListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeColumnarVersionListResponseBody) SetRequestId(v string) *DescribeColumnarVersionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColumnarVersionListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeColumnarVersionListResponseBodyData struct {
	VersionList []*string `json:"VersionList,omitempty" xml:"VersionList,omitempty" type:"Repeated"`
}

func (s DescribeColumnarVersionListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarVersionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeColumnarVersionListResponseBodyData) GetVersionList() []*string {
	return s.VersionList
}

func (s *DescribeColumnarVersionListResponseBodyData) SetVersionList(v []*string) *DescribeColumnarVersionListResponseBodyData {
	s.VersionList = v
	return s
}

func (s *DescribeColumnarVersionListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
