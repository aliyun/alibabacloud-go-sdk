// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeLogTypeResponseBody
	GetCode() *int32
	SetData(v []*DescribeLogTypeResponseBodyData) *DescribeLogTypeResponseBody
	GetData() []*DescribeLogTypeResponseBodyData
	SetMessage(v string) *DescribeLogTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeLogTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeLogTypeResponseBody
	GetSuccess() *bool
}

type DescribeLogTypeResponseBody struct {
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
	Data []*DescribeLogTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DescribeLogTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogTypeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeLogTypeResponseBody) GetData() []*DescribeLogTypeResponseBodyData {
	return s.Data
}

func (s *DescribeLogTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeLogTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeLogTypeResponseBody) SetCode(v int32) *DescribeLogTypeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogTypeResponseBody) SetData(v []*DescribeLogTypeResponseBodyData) *DescribeLogTypeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLogTypeResponseBody) SetMessage(v string) *DescribeLogTypeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogTypeResponseBody) SetRequestId(v string) *DescribeLogTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogTypeResponseBody) SetSuccess(v bool) *DescribeLogTypeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeLogTypeResponseBody) Validate() error {
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

type DescribeLogTypeResponseBodyData struct {
	// The log type of the rule.
	//
	// example:
	//
	// HTTP_ACTIVITY
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The internal code of the log type.
	//
	// example:
	//
	// sas.cloudsiem.prod.http_activity
	LogTypeName *string `json:"LogTypeName,omitempty" xml:"LogTypeName,omitempty"`
}

func (s DescribeLogTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLogTypeResponseBodyData) GetLogType() *string {
	return s.LogType
}

func (s *DescribeLogTypeResponseBodyData) GetLogTypeName() *string {
	return s.LogTypeName
}

func (s *DescribeLogTypeResponseBodyData) SetLogType(v string) *DescribeLogTypeResponseBodyData {
	s.LogType = &v
	return s
}

func (s *DescribeLogTypeResponseBodyData) SetLogTypeName(v string) *DescribeLogTypeResponseBodyData {
	s.LogTypeName = &v
	return s
}

func (s *DescribeLogTypeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
