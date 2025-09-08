// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAlertSourceResponseBody
	GetCode() *int32
	SetData(v []*DescribeAlertSourceResponseBodyData) *DescribeAlertSourceResponseBody
	GetData() []*DescribeAlertSourceResponseBodyData
	SetMessage(v string) *DescribeAlertSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertSourceResponseBody
	GetSuccess() *bool
}

type DescribeAlertSourceResponseBody struct {
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
	Data []*DescribeAlertSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DescribeAlertSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAlertSourceResponseBody) GetData() []*DescribeAlertSourceResponseBodyData {
	return s.Data
}

func (s *DescribeAlertSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertSourceResponseBody) SetCode(v int32) *DescribeAlertSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertSourceResponseBody) SetData(v []*DescribeAlertSourceResponseBodyData) *DescribeAlertSourceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertSourceResponseBody) SetMessage(v string) *DescribeAlertSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertSourceResponseBody) SetRequestId(v string) *DescribeAlertSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertSourceResponseBody) SetSuccess(v bool) *DescribeAlertSourceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertSourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertSourceResponseBodyData struct {
	// The internal code of the alert data source.
	//
	// example:
	//
	// aliyun.siem.alert_datasource.sas
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The name of the alert data source.
	//
	// example:
	//
	// sas
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
}

func (s DescribeAlertSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *DescribeAlertSourceResponseBodyData) GetSourceName() *string {
	return s.SourceName
}

func (s *DescribeAlertSourceResponseBodyData) SetSource(v string) *DescribeAlertSourceResponseBodyData {
	s.Source = &v
	return s
}

func (s *DescribeAlertSourceResponseBodyData) SetSourceName(v string) *DescribeAlertSourceResponseBodyData {
	s.SourceName = &v
	return s
}

func (s *DescribeAlertSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
