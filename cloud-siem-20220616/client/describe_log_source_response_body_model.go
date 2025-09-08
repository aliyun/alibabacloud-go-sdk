// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeLogSourceResponseBody
	GetCode() *int32
	SetData(v []*DescribeLogSourceResponseBodyData) *DescribeLogSourceResponseBody
	GetData() []*DescribeLogSourceResponseBodyData
	SetMessage(v string) *DescribeLogSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeLogSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeLogSourceResponseBody
	GetSuccess() *bool
}

type DescribeLogSourceResponseBody struct {
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
	Data []*DescribeLogSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DescribeLogSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogSourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeLogSourceResponseBody) GetData() []*DescribeLogSourceResponseBodyData {
	return s.Data
}

func (s *DescribeLogSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeLogSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeLogSourceResponseBody) SetCode(v int32) *DescribeLogSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogSourceResponseBody) SetData(v []*DescribeLogSourceResponseBodyData) *DescribeLogSourceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLogSourceResponseBody) SetMessage(v string) *DescribeLogSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogSourceResponseBody) SetRequestId(v string) *DescribeLogSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogSourceResponseBody) SetSuccess(v bool) *DescribeLogSourceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeLogSourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLogSourceResponseBodyData struct {
	// The log source of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The internal code of the log source.
	//
	// example:
	//
	// sas.cloudsiem.prod.cloud_siem_aegis_sas_alert
	LogSourceName *string `json:"LogSourceName,omitempty" xml:"LogSourceName,omitempty"`
}

func (s DescribeLogSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLogSourceResponseBodyData) GetLogSource() *string {
	return s.LogSource
}

func (s *DescribeLogSourceResponseBodyData) GetLogSourceName() *string {
	return s.LogSourceName
}

func (s *DescribeLogSourceResponseBodyData) SetLogSource(v string) *DescribeLogSourceResponseBodyData {
	s.LogSource = &v
	return s
}

func (s *DescribeLogSourceResponseBodyData) SetLogSourceName(v string) *DescribeLogSourceResponseBodyData {
	s.LogSourceName = &v
	return s
}

func (s *DescribeLogSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
