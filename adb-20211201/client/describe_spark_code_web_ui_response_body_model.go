// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkCodeWebUiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DescribeSparkCodeWebUiResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSparkCodeWebUiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSparkCodeWebUiResponseBody
	GetSuccess() *bool
	SetUrl(v string) *DescribeSparkCodeWebUiResponseBody
	GetUrl() *string
}

type DescribeSparkCodeWebUiResponseBody struct {
	// The returned message.
	//
	// 	- If the request was successful, **SUCCESS*	- is returned.
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 774DDC37-1908-58F6-B9CA-99E3E45965A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The URL of the web UI for the Spark application.
	//
	// example:
	//
	// https://adb-subuser-cn-hangzhou-1358535755648527-100000648.oss-cn-hangzhou.aliyuncs.com/%3Facl?Expires=1681295967&OSSAccessKeyId=LTAI5tB7NAkm25oiGASu****&Signature=hKAZ1vgvhJ%2FD8hNHTuX%2FOOBWht****
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeSparkCodeWebUiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkCodeWebUiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeWebUiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSparkCodeWebUiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSparkCodeWebUiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSparkCodeWebUiResponseBody) GetUrl() *string {
	return s.Url
}

func (s *DescribeSparkCodeWebUiResponseBody) SetMessage(v string) *DescribeSparkCodeWebUiResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSparkCodeWebUiResponseBody) SetRequestId(v string) *DescribeSparkCodeWebUiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSparkCodeWebUiResponseBody) SetSuccess(v bool) *DescribeSparkCodeWebUiResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSparkCodeWebUiResponseBody) SetUrl(v string) *DescribeSparkCodeWebUiResponseBody {
	s.Url = &v
	return s
}

func (s *DescribeSparkCodeWebUiResponseBody) Validate() error {
	return dara.Validate(s)
}
