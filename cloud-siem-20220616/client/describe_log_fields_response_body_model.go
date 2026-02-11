// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeLogFieldsResponseBody
	GetCode() *int32
	SetData(v []*DescribeLogFieldsResponseBodyData) *DescribeLogFieldsResponseBody
	GetData() []*DescribeLogFieldsResponseBodyData
	SetMessage(v string) *DescribeLogFieldsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeLogFieldsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeLogFieldsResponseBody
	GetSuccess() *bool
}

type DescribeLogFieldsResponseBody struct {
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
	Data []*DescribeLogFieldsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DescribeLogFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogFieldsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeLogFieldsResponseBody) GetData() []*DescribeLogFieldsResponseBodyData {
	return s.Data
}

func (s *DescribeLogFieldsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeLogFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogFieldsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeLogFieldsResponseBody) SetCode(v int32) *DescribeLogFieldsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogFieldsResponseBody) SetData(v []*DescribeLogFieldsResponseBodyData) *DescribeLogFieldsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLogFieldsResponseBody) SetMessage(v string) *DescribeLogFieldsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogFieldsResponseBody) SetRequestId(v string) *DescribeLogFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogFieldsResponseBody) SetSuccess(v bool) *DescribeLogFieldsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeLogFieldsResponseBody) Validate() error {
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

type DescribeLogFieldsResponseBodyData struct {
	// The type of the log to which the field belongs.
	//
	// example:
	//
	// HTTP_ACTIVITY
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The internal code of the field description.
	//
	// example:
	//
	// sas.cloudsiem.prod.activity_name
	FieldDesc *string `json:"FieldDesc,omitempty" xml:"FieldDesc,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// activity_name
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The data type of the field. Valid values:
	//
	// 	- varchar
	//
	// 	- bigint
	//
	// example:
	//
	// varchar
	FieldType *string `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	// The log source to which the field belongs.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
}

func (s DescribeLogFieldsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogFieldsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLogFieldsResponseBodyData) GetActivityName() *string {
	return s.ActivityName
}

func (s *DescribeLogFieldsResponseBodyData) GetFieldDesc() *string {
	return s.FieldDesc
}

func (s *DescribeLogFieldsResponseBodyData) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeLogFieldsResponseBodyData) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeLogFieldsResponseBodyData) GetLogCode() *string {
	return s.LogCode
}

func (s *DescribeLogFieldsResponseBodyData) SetActivityName(v string) *DescribeLogFieldsResponseBodyData {
	s.ActivityName = &v
	return s
}

func (s *DescribeLogFieldsResponseBodyData) SetFieldDesc(v string) *DescribeLogFieldsResponseBodyData {
	s.FieldDesc = &v
	return s
}

func (s *DescribeLogFieldsResponseBodyData) SetFieldName(v string) *DescribeLogFieldsResponseBodyData {
	s.FieldName = &v
	return s
}

func (s *DescribeLogFieldsResponseBodyData) SetFieldType(v string) *DescribeLogFieldsResponseBodyData {
	s.FieldType = &v
	return s
}

func (s *DescribeLogFieldsResponseBodyData) SetLogCode(v string) *DescribeLogFieldsResponseBodyData {
	s.LogCode = &v
	return s
}

func (s *DescribeLogFieldsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
