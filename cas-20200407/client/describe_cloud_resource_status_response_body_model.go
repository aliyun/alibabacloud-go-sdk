// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeCloudResourceStatusResponseBodyData) *DescribeCloudResourceStatusResponseBody
	GetData() []*DescribeCloudResourceStatusResponseBodyData
	SetRequestId(v string) *DescribeCloudResourceStatusResponseBody
	GetRequestId() *string
}

type DescribeCloudResourceStatusResponseBody struct {
	// The response parameters.
	Data []*DescribeCloudResourceStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 09470F19-CEE8-5C63-BF2C-02B5E3F07A17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudResourceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceStatusResponseBody) GetData() []*DescribeCloudResourceStatusResponseBodyData {
	return s.Data
}

func (s *DescribeCloudResourceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudResourceStatusResponseBody) SetData(v []*DescribeCloudResourceStatusResponseBodyData) *DescribeCloudResourceStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudResourceStatusResponseBody) SetRequestId(v string) *DescribeCloudResourceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudResourceStatusResponseBody) Validate() error {
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

type DescribeCloudResourceStatusResponseBodyData struct {
	// The cloud service provider.
	//
	// example:
	//
	// aliyun
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The cloud service.
	//
	// example:
	//
	// OSS
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The total number of cloud resources on which certificates are deployed.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudResourceStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceStatusResponseBodyData) GetCloudName() *string {
	return s.CloudName
}

func (s *DescribeCloudResourceStatusResponseBodyData) GetCloudProduct() *string {
	return s.CloudProduct
}

func (s *DescribeCloudResourceStatusResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCloudResourceStatusResponseBodyData) SetCloudName(v string) *DescribeCloudResourceStatusResponseBodyData {
	s.CloudName = &v
	return s
}

func (s *DescribeCloudResourceStatusResponseBodyData) SetCloudProduct(v string) *DescribeCloudResourceStatusResponseBodyData {
	s.CloudProduct = &v
	return s
}

func (s *DescribeCloudResourceStatusResponseBodyData) SetTotalCount(v int64) *DescribeCloudResourceStatusResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudResourceStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
