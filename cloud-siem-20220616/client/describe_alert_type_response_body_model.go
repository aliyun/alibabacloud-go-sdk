// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAlertTypeResponseBody
	GetCode() *int32
	SetData(v []*DescribeAlertTypeResponseBodyData) *DescribeAlertTypeResponseBody
	GetData() []*DescribeAlertTypeResponseBodyData
	SetMessage(v string) *DescribeAlertTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertTypeResponseBody
	GetSuccess() *bool
}

type DescribeAlertTypeResponseBody struct {
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
	Data []*DescribeAlertTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DescribeAlertTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAlertTypeResponseBody) GetData() []*DescribeAlertTypeResponseBodyData {
	return s.Data
}

func (s *DescribeAlertTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertTypeResponseBody) SetCode(v int32) *DescribeAlertTypeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetData(v []*DescribeAlertTypeResponseBodyData) *DescribeAlertTypeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetMessage(v string) *DescribeAlertTypeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetRequestId(v string) *DescribeAlertTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetSuccess(v bool) *DescribeAlertTypeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) Validate() error {
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

type DescribeAlertTypeResponseBodyData struct {
	// The type of the risk.
	//
	// example:
	//
	// WEBSHELL
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The internal code of the risk type.
	//
	// example:
	//
	// siem_rule_type_process_abnormal_command
	AlertTypeMds *string `json:"AlertTypeMds,omitempty" xml:"AlertTypeMds,omitempty"`
}

func (s DescribeAlertTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeResponseBodyData) GetAlertType() *string {
	return s.AlertType
}

func (s *DescribeAlertTypeResponseBodyData) GetAlertTypeMds() *string {
	return s.AlertTypeMds
}

func (s *DescribeAlertTypeResponseBodyData) SetAlertType(v string) *DescribeAlertTypeResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertTypeResponseBodyData) SetAlertTypeMds(v string) *DescribeAlertTypeResponseBodyData {
	s.AlertTypeMds = &v
	return s
}

func (s *DescribeAlertTypeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
