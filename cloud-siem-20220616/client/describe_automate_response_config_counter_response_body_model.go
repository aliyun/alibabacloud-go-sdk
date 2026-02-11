// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutomateResponseConfigCounterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAutomateResponseConfigCounterResponseBody
	GetCode() *int32
	SetData(v *DescribeAutomateResponseConfigCounterResponseBodyData) *DescribeAutomateResponseConfigCounterResponseBody
	GetData() *DescribeAutomateResponseConfigCounterResponseBodyData
	SetMessage(v string) *DescribeAutomateResponseConfigCounterResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAutomateResponseConfigCounterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAutomateResponseConfigCounterResponseBody
	GetSuccess() *bool
}

type DescribeAutomateResponseConfigCounterResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeAutomateResponseConfigCounterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAutomateResponseConfigCounterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigCounterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) GetData() *DescribeAutomateResponseConfigCounterResponseBodyData {
	return s.Data
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetCode(v int32) *DescribeAutomateResponseConfigCounterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetData(v *DescribeAutomateResponseConfigCounterResponseBodyData) *DescribeAutomateResponseConfigCounterResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetMessage(v string) *DescribeAutomateResponseConfigCounterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetRequestId(v string) *DescribeAutomateResponseConfigCounterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetSuccess(v bool) *DescribeAutomateResponseConfigCounterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutomateResponseConfigCounterResponseBodyData struct {
	// The total number of rules.
	//
	// example:
	//
	// 20
	All *int64 `json:"All,omitempty" xml:"All,omitempty"`
	// The number of enabled rules.
	//
	// example:
	//
	// 10
	Online *int64 `json:"Online,omitempty" xml:"Online,omitempty"`
}

func (s DescribeAutomateResponseConfigCounterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigCounterResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigCounterResponseBodyData) GetAll() *int64 {
	return s.All
}

func (s *DescribeAutomateResponseConfigCounterResponseBodyData) GetOnline() *int64 {
	return s.Online
}

func (s *DescribeAutomateResponseConfigCounterResponseBodyData) SetAll(v int64) *DescribeAutomateResponseConfigCounterResponseBodyData {
	s.All = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBodyData) SetOnline(v int64) *DescribeAutomateResponseConfigCounterResponseBodyData {
	s.Online = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
