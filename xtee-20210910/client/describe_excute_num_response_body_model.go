// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExcuteNumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v []*string) *DescribeExcuteNumResponseBody
	GetCategories() []*string
	SetData(v []*string) *DescribeExcuteNumResponseBody
	GetData() []*string
	SetRequestId(v string) *DescribeExcuteNumResponseBody
	GetRequestId() *string
}

type DescribeExcuteNumResponseBody struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	Data       []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExcuteNumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcuteNumResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExcuteNumResponseBody) GetCategories() []*string {
	return s.Categories
}

func (s *DescribeExcuteNumResponseBody) GetData() []*string {
	return s.Data
}

func (s *DescribeExcuteNumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExcuteNumResponseBody) SetCategories(v []*string) *DescribeExcuteNumResponseBody {
	s.Categories = v
	return s
}

func (s *DescribeExcuteNumResponseBody) SetData(v []*string) *DescribeExcuteNumResponseBody {
	s.Data = v
	return s
}

func (s *DescribeExcuteNumResponseBody) SetRequestId(v string) *DescribeExcuteNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExcuteNumResponseBody) Validate() error {
	return dara.Validate(s)
}
