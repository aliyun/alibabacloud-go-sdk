// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssStatusV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeOssStatusV2ResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *DescribeOssStatusV2ResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *DescribeOssStatusV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeOssStatusV2ResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeOssStatusV2ResponseBodyResultObject) *DescribeOssStatusV2ResponseBody
	GetResultObject() *DescribeOssStatusV2ResponseBodyResultObject
	SetSuccess(v bool) *DescribeOssStatusV2ResponseBody
	GetSuccess() *bool
}

type DescribeOssStatusV2ResponseBody struct {
	// Return code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// B506328A-D84B-4750-82C7-6A207C585CF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information.
	ResultObject *DescribeOssStatusV2ResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	// Whether the response was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeOssStatusV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssStatusV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssStatusV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeOssStatusV2ResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeOssStatusV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeOssStatusV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOssStatusV2ResponseBody) GetResultObject() *DescribeOssStatusV2ResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeOssStatusV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeOssStatusV2ResponseBody) SetCode(v string) *DescribeOssStatusV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeOssStatusV2ResponseBody) SetHttpStatusCode(v int64) *DescribeOssStatusV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeOssStatusV2ResponseBody) SetMessage(v string) *DescribeOssStatusV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeOssStatusV2ResponseBody) SetRequestId(v string) *DescribeOssStatusV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssStatusV2ResponseBody) SetResultObject(v *DescribeOssStatusV2ResponseBodyResultObject) *DescribeOssStatusV2ResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeOssStatusV2ResponseBody) SetSuccess(v bool) *DescribeOssStatusV2ResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeOssStatusV2ResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOssStatusV2ResponseBodyResultObject struct {
	// Bucket name.
	//
	// example:
	//
	// cn-hangzhou-aliyun-cloudauth-20250516xxxxxx
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// Region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// User activation status, **SUCCESS*	- indicates activated.
	//
	// example:
	//
	// SUCCESS
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s DescribeOssStatusV2ResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssStatusV2ResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeOssStatusV2ResponseBodyResultObject) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeOssStatusV2ResponseBodyResultObject) GetRegion() *string {
	return s.Region
}

func (s *DescribeOssStatusV2ResponseBodyResultObject) GetUserStatus() *string {
	return s.UserStatus
}

func (s *DescribeOssStatusV2ResponseBodyResultObject) SetBucketName(v string) *DescribeOssStatusV2ResponseBodyResultObject {
	s.BucketName = &v
	return s
}

func (s *DescribeOssStatusV2ResponseBodyResultObject) SetRegion(v string) *DescribeOssStatusV2ResponseBodyResultObject {
	s.Region = &v
	return s
}

func (s *DescribeOssStatusV2ResponseBodyResultObject) SetUserStatus(v string) *DescribeOssStatusV2ResponseBodyResultObject {
	s.UserStatus = &v
	return s
}

func (s *DescribeOssStatusV2ResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
