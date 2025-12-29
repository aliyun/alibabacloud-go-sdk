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
	// The HTTP status code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **InvalidPhoneNumber.Check**: The phone number is invalid.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *DescribeEmptyNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot issues.
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
	// The specified phone number.
	//
	// example:
	//
	// 189****1234
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The returned status for the queried phone number. Valid values:
	//
	// 	- **EMPTY**: The queried phone number is a nonexistent number.
	//
	// 	- **NORMAL**: The queried phone number is valid.
	//
	// 	- **SUSPECT_EMPTY**: The queried phone number is suspected to be a nonexistent number.
	//
	// 	- **UNKNOWN**: The queried phone number is unknown.
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
