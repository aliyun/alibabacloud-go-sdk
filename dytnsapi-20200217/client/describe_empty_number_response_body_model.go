// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEmptyNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeEmptyNumberResponseBody
	GetCode() *string
	SetData(v *DescribeEmptyNumberResponseBodyData) *DescribeEmptyNumberResponseBody
	GetData() *DescribeEmptyNumberResponseBodyData
	SetMessage(v string) *DescribeEmptyNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEmptyNumberResponseBody
	GetRequestId() *string
}

type DescribeEmptyNumberResponseBody struct {
	// 返回状态码。取值：
	//
	// - **OK**：成功。
	//
	// - **InvalidPhoneNumber.Check**：手机号非法。
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回结果。
	Data *DescribeEmptyNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 状态码的描述。
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 公共参数，每个请求返回的ID都是唯一的，可用于排查和定位问题。
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEmptyNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmptyNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeEmptyNumberResponseBody) GetData() *DescribeEmptyNumberResponseBodyData {
	return s.Data
}

func (s *DescribeEmptyNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEmptyNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEmptyNumberResponseBody) SetCode(v string) *DescribeEmptyNumberResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEmptyNumberResponseBody) SetData(v *DescribeEmptyNumberResponseBodyData) *DescribeEmptyNumberResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEmptyNumberResponseBody) SetMessage(v string) *DescribeEmptyNumberResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEmptyNumberResponseBody) SetRequestId(v string) *DescribeEmptyNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEmptyNumberResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEmptyNumberResponseBodyData struct {
	// 传入的手机号。
	//
	// example:
	//
	// 189****1234
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// 检测手机号返回状态。取值：
	//
	// - **EMPTY**：空号。
	//
	// - **NORMAL**：正常。
	//
	// - **SUSPECT_EMPTY**：疑似空号。
	//
	// - **UNKNOWN**：未知。
	//
	// example:
	//
	// EMPTY
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEmptyNumberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmptyNumberResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEmptyNumberResponseBodyData) GetNumber() *string {
	return s.Number
}

func (s *DescribeEmptyNumberResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeEmptyNumberResponseBodyData) SetNumber(v string) *DescribeEmptyNumberResponseBodyData {
	s.Number = &v
	return s
}

func (s *DescribeEmptyNumberResponseBodyData) SetStatus(v string) *DescribeEmptyNumberResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeEmptyNumberResponseBodyData) Validate() error {
	return dara.Validate(s)
}
