// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertSourceWithEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAlertSourceWithEventResponseBody
	GetCode() *int32
	SetData(v []*DescribeAlertSourceWithEventResponseBodyData) *DescribeAlertSourceWithEventResponseBody
	GetData() []*DescribeAlertSourceWithEventResponseBodyData
	SetMessage(v string) *DescribeAlertSourceWithEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertSourceWithEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertSourceWithEventResponseBody
	GetSuccess() *bool
}

type DescribeAlertSourceWithEventResponseBody struct {
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
	Data []*DescribeAlertSourceWithEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DescribeAlertSourceWithEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSourceWithEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceWithEventResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAlertSourceWithEventResponseBody) GetData() []*DescribeAlertSourceWithEventResponseBodyData {
	return s.Data
}

func (s *DescribeAlertSourceWithEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertSourceWithEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertSourceWithEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertSourceWithEventResponseBody) SetCode(v int32) *DescribeAlertSourceWithEventResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBody) SetData(v []*DescribeAlertSourceWithEventResponseBodyData) *DescribeAlertSourceWithEventResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBody) SetMessage(v string) *DescribeAlertSourceWithEventResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBody) SetRequestId(v string) *DescribeAlertSourceWithEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBody) SetSuccess(v bool) *DescribeAlertSourceWithEventResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBody) Validate() error {
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

type DescribeAlertSourceWithEventResponseBodyData struct {
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

func (s DescribeAlertSourceWithEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSourceWithEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceWithEventResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *DescribeAlertSourceWithEventResponseBodyData) GetSourceName() *string {
	return s.SourceName
}

func (s *DescribeAlertSourceWithEventResponseBodyData) SetSource(v string) *DescribeAlertSourceWithEventResponseBodyData {
	s.Source = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBodyData) SetSourceName(v string) *DescribeAlertSourceWithEventResponseBodyData {
	s.SourceName = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBodyData) Validate() error {
	return dara.Validate(s)
}
