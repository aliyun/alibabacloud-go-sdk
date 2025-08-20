// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkCodeOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DescribeSparkCodeOutputResponseBody
	GetMessage() *string
	SetOutput(v string) *DescribeSparkCodeOutputResponseBody
	GetOutput() *string
	SetRequestId(v string) *DescribeSparkCodeOutputResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSparkCodeOutputResponseBody
	GetSuccess() *bool
}

type DescribeSparkCodeOutputResponseBody struct {
	// The returned message.
	//
	// 	- If the request was successful, **Success*	- is returned.
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution result, which is in the format of JSON objects.
	//
	// example:
	//
	// "{\\"schema\\":[\\"id\\",\\"name\\",\\"age\\"],\\"data\\":[\\"{\\\\\\"id\\\\\\":10,\\\\\\"name\\\\\\":\\\\\\"z\\\\\\",\\\\\\"age\\\\\\":123}\\",\\"{\\\\\\"id\\\\\\":2,\\\\\\"name\\\\\\":\\\\\\"b\\\\\\",\\\\\\"age\\\\\\":17}\\",\\"{\\\\\\"id\\\\\\":1,\\\\\\"name\\\\\\":\\\\\\"a\\\\\\",\\\\\\"age\\\\\\":15}\\",\\"{\\\\\\"id\\\\\\":3,\\\\\\"name\\\\\\":\\\\\\"c\\\\\\",\\\\\\"age\\\\\\":222}\\",\\"{\\\\\\"id\\\\\\":10,\\\\\\"name\\\\\\":\\\\\\"z\\\\\\",\\\\\\"age\\\\\\":123}\\"],\\"haveRows\\":true,\\"rowNumber\\":6}"
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
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
}

func (s DescribeSparkCodeOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkCodeOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeOutputResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSparkCodeOutputResponseBody) GetOutput() *string {
	return s.Output
}

func (s *DescribeSparkCodeOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSparkCodeOutputResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSparkCodeOutputResponseBody) SetMessage(v string) *DescribeSparkCodeOutputResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSparkCodeOutputResponseBody) SetOutput(v string) *DescribeSparkCodeOutputResponseBody {
	s.Output = &v
	return s
}

func (s *DescribeSparkCodeOutputResponseBody) SetRequestId(v string) *DescribeSparkCodeOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSparkCodeOutputResponseBody) SetSuccess(v bool) *DescribeSparkCodeOutputResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSparkCodeOutputResponseBody) Validate() error {
	return dara.Validate(s)
}
