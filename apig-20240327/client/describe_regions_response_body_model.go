// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeRegionsResponseBody
	GetCode() *string
	SetData(v []*DescribeRegionsResponseBodyData) *DescribeRegionsResponseBody
	GetData() []*DescribeRegionsResponseBodyData
	SetMessage(v string) *DescribeRegionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
}

type DescribeRegionsResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data []*DescribeRegionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// E6BD6C79-32B1-5D7E-A89A-93C5A6B7xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeRegionsResponseBody) GetData() []*DescribeRegionsResponseBodyData {
	return s.Data
}

func (s *DescribeRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) SetCode(v string) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetData(v []*DescribeRegionsResponseBodyData) *DescribeRegionsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRegionsResponseBody) SetMessage(v string) *DescribeRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyData struct {
	LocalName *string `json:"localName,omitempty" xml:"localName,omitempty"`
	RegionId  *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescribeRegionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyData) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyData) SetLocalName(v string) *DescribeRegionsResponseBodyData {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyData) SetRegionId(v string) *DescribeRegionsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
