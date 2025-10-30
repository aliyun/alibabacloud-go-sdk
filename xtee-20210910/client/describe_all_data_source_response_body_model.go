// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAllDataSourceResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeAllDataSourceResponseBodyResultObject) *DescribeAllDataSourceResponseBody
	GetResultObject() []*DescribeAllDataSourceResponseBodyResultObject
}

type DescribeAllDataSourceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeAllDataSourceResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAllDataSourceResponseBody) GetResultObject() []*DescribeAllDataSourceResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeAllDataSourceResponseBody) SetRequestId(v string) *DescribeAllDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetResultObject(v []*DescribeAllDataSourceResponseBodyResultObject) *DescribeAllDataSourceResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeAllDataSourceResponseBody) Validate() error {
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

type DescribeAllDataSourceResponseBodyResultObject struct {
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
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Policy primary key ID
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DescribeAllDataSourceResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeAllDataSourceResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeAllDataSourceResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeAllDataSourceResponseBodyResultObject) SetEventCode(v string) *DescribeAllDataSourceResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyResultObject) SetEventName(v string) *DescribeAllDataSourceResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyResultObject) SetId(v int64) *DescribeAllDataSourceResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
