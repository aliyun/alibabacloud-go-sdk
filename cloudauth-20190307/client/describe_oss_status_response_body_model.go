// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeOssStatusResponseBodyData) *DescribeOssStatusResponseBody
	GetData() *DescribeOssStatusResponseBodyData
	SetRequestId(v string) *DescribeOssStatusResponseBody
	GetRequestId() *string
}

type DescribeOssStatusResponseBody struct {
	// Returned data.
	Data *DescribeOssStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of this request.
	//
	// example:
	//
	// F2DB870B-EEB7-51BD-9F0A-B5D8D3C79308
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOssStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssStatusResponseBody) GetData() *DescribeOssStatusResponseBodyData {
	return s.Data
}

func (s *DescribeOssStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOssStatusResponseBody) SetData(v *DescribeOssStatusResponseBodyData) *DescribeOssStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOssStatusResponseBody) SetRequestId(v string) *DescribeOssStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOssStatusResponseBodyData struct {
	// The name of the OSS bucket for delivering authentication information.
	//
	// example:
	//
	// cn-hangzhou-aliyun-cloudauth-20250516xxxxxx
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// User activation status, SUCCESS indicates activated.
	//
	// example:
	//
	// SUCCESS
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s DescribeOssStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOssStatusResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeOssStatusResponseBodyData) GetUserStatus() *string {
	return s.UserStatus
}

func (s *DescribeOssStatusResponseBodyData) SetBucketName(v string) *DescribeOssStatusResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *DescribeOssStatusResponseBodyData) SetUserStatus(v string) *DescribeOssStatusResponseBodyData {
	s.UserStatus = &v
	return s
}

func (s *DescribeOssStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
