// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDasProResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeInstanceDasProResponseBody
	GetCode() *string
	SetData(v bool) *DescribeInstanceDasProResponseBody
	GetData() *bool
	SetMessage(v string) *DescribeInstanceDasProResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeInstanceDasProResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeInstanceDasProResponseBody
	GetSuccess() *string
}

type DescribeInstanceDasProResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether DAS Enterprise Edition is enabled for the database instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9CB97BC4-6479-55D0-B9D0-EA925AFE****
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstanceDasProResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDasProResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDasProResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceDasProResponseBody) GetData() *bool {
	return s.Data
}

func (s *DescribeInstanceDasProResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeInstanceDasProResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceDasProResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeInstanceDasProResponseBody) SetCode(v string) *DescribeInstanceDasProResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceDasProResponseBody) SetData(v bool) *DescribeInstanceDasProResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeInstanceDasProResponseBody) SetMessage(v string) *DescribeInstanceDasProResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceDasProResponseBody) SetRequestId(v string) *DescribeInstanceDasProResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceDasProResponseBody) SetSuccess(v string) *DescribeInstanceDasProResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceDasProResponseBody) Validate() error {
	return dara.Validate(s)
}
