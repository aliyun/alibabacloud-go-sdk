// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneAllEventNameCodeListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSceneAllEventNameCodeListResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeSceneAllEventNameCodeListResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeSceneAllEventNameCodeListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSceneAllEventNameCodeListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeSceneAllEventNameCodeListResponseBodyResultObject) *DescribeSceneAllEventNameCodeListResponseBody
	GetResultObject() []*DescribeSceneAllEventNameCodeListResponseBodyResultObject
	SetSuccess(v bool) *DescribeSceneAllEventNameCodeListResponseBody
	GetSuccess() *bool
}

type DescribeSceneAllEventNameCodeListResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error details
	//
	// example:
	//
	// The input parameter data is not valid. order_storage_company_num component not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, which is unique for each request, facilitating subsequent troubleshooting
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject []*DescribeSceneAllEventNameCodeListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Indicates whether the operation was successful, where true means success.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeSceneAllEventNameCodeListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneAllEventNameCodeListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSceneAllEventNameCodeListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSceneAllEventNameCodeListResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeSceneAllEventNameCodeListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSceneAllEventNameCodeListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSceneAllEventNameCodeListResponseBody) GetResultObject() []*DescribeSceneAllEventNameCodeListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSceneAllEventNameCodeListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSceneAllEventNameCodeListResponseBody) SetCode(v string) *DescribeSceneAllEventNameCodeListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBody) SetHttpStatusCode(v string) *DescribeSceneAllEventNameCodeListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBody) SetMessage(v string) *DescribeSceneAllEventNameCodeListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBody) SetRequestId(v string) *DescribeSceneAllEventNameCodeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBody) SetResultObject(v []*DescribeSceneAllEventNameCodeListResponseBodyResultObject) *DescribeSceneAllEventNameCodeListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBody) SetSuccess(v bool) *DescribeSceneAllEventNameCodeListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSceneAllEventNameCodeListResponseBodyResultObject struct {
	// Child objects
	Children []*DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// Creation type
	//
	// example:
	//
	// MORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Event code
	//
	// example:
	//
	// de_aszbjb7236
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 营销风险识别_增强版
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Event type
	//
	// example:
	//
	// MAIN
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
}

func (s DescribeSceneAllEventNameCodeListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneAllEventNameCodeListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObject) GetChildren() []*DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren {
	return s.Children
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObject) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObject) GetEventType() *string {
	return s.EventType
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObject) SetChildren(v []*DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren) *DescribeSceneAllEventNameCodeListResponseBodyResultObject {
	s.Children = v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObject) SetCreateType(v string) *DescribeSceneAllEventNameCodeListResponseBodyResultObject {
	s.CreateType = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObject) SetEventCode(v string) *DescribeSceneAllEventNameCodeListResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObject) SetEventName(v string) *DescribeSceneAllEventNameCodeListResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObject) SetEventType(v string) *DescribeSceneAllEventNameCodeListResponseBodyResultObject {
	s.EventType = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObject) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren struct {
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Event code
	//
	// example:
	//
	// de_aamexg3015
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 测试
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Event type
	//
	// example:
	//
	// BYPASS
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
}

func (s DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren) GoString() string {
	return s.String()
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren) GetEventName() *string {
	return s.EventName
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren) GetEventType() *string {
	return s.EventType
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren) SetCreateType(v string) *DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren {
	s.CreateType = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren) SetEventCode(v string) *DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren {
	s.EventCode = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren) SetEventName(v string) *DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren {
	s.EventName = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren) SetEventType(v string) *DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren {
	s.EventType = &v
	return s
}

func (s *DescribeSceneAllEventNameCodeListResponseBodyResultObjectChildren) Validate() error {
	return dara.Validate(s)
}
