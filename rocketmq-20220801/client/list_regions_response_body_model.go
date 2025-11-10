// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRegionsResponseBody
	GetCode() *string
	SetData(v []*ListRegionsResponseBodyData) *ListRegionsResponseBody
	GetData() []*ListRegionsResponseBodyData
	SetDynamicCode(v string) *ListRegionsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListRegionsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListRegionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListRegionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRegionsResponseBody
	GetSuccess() *bool
}

type ListRegionsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingPageNumber
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data []*ListRegionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The dynamic error code.
	//
	// example:
	//
	// ConsumerGroupId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B962390-D84B-5D44-8C11-79DF40299D41
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRegionsResponseBody) GetData() []*ListRegionsResponseBodyData {
	return s.Data
}

func (s *ListRegionsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListRegionsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListRegionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRegionsResponseBody) SetCode(v string) *ListRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRegionsResponseBody) SetData(v []*ListRegionsResponseBodyData) *ListRegionsResponseBody {
	s.Data = v
	return s
}

func (s *ListRegionsResponseBody) SetDynamicCode(v string) *ListRegionsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListRegionsResponseBody) SetDynamicMessage(v string) *ListRegionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListRegionsResponseBody) SetHttpStatusCode(v int32) *ListRegionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRegionsResponseBody) SetMessage(v string) *ListRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) SetSuccess(v bool) *ListRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRegionsResponseBody) Validate() error {
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

type ListRegionsResponseBodyData struct {
	// The time when the ApsaraMQ for RocketMQ instance was created.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The region name.
	//
	// example:
	//
	// hangzhou
	RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	// Indicates whether ApsaraMQ for RocketMQ V4 is activated.
	//
	// example:
	//
	// true
	SupportRocketmqV4 *bool `json:"supportRocketmqV4,omitempty" xml:"supportRocketmqV4,omitempty"`
	// Indicates whether ApsaraMQ for RocketMQ V5 is activated.
	//
	// example:
	//
	// true
	SupportRocketmqV5 *bool `json:"supportRocketmqV5,omitempty" xml:"supportRocketmqV5,omitempty"`
	// The region tags.
	Tags []*ListRegionsResponseBodyDataTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The time when the ApsaraMQ for RocketMQ instance was last modified.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListRegionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRegionsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRegionsResponseBodyData) GetRegionName() *string {
	return s.RegionName
}

func (s *ListRegionsResponseBodyData) GetSupportRocketmqV4() *bool {
	return s.SupportRocketmqV4
}

func (s *ListRegionsResponseBodyData) GetSupportRocketmqV5() *bool {
	return s.SupportRocketmqV5
}

func (s *ListRegionsResponseBodyData) GetTags() []*ListRegionsResponseBodyDataTags {
	return s.Tags
}

func (s *ListRegionsResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListRegionsResponseBodyData) SetCreateTime(v string) *ListRegionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListRegionsResponseBodyData) SetRegionId(v string) *ListRegionsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListRegionsResponseBodyData) SetRegionName(v string) *ListRegionsResponseBodyData {
	s.RegionName = &v
	return s
}

func (s *ListRegionsResponseBodyData) SetSupportRocketmqV4(v bool) *ListRegionsResponseBodyData {
	s.SupportRocketmqV4 = &v
	return s
}

func (s *ListRegionsResponseBodyData) SetSupportRocketmqV5(v bool) *ListRegionsResponseBodyData {
	s.SupportRocketmqV5 = &v
	return s
}

func (s *ListRegionsResponseBodyData) SetTags(v []*ListRegionsResponseBodyDataTags) *ListRegionsResponseBodyData {
	s.Tags = v
	return s
}

func (s *ListRegionsResponseBodyData) SetUpdateTime(v string) *ListRegionsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListRegionsResponseBodyData) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRegionsResponseBodyDataTags struct {
	// The tag code.
	//
	// example:
	//
	// xx
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	// The tag value.
	//
	// example:
	//
	// xx
	TagValue interface{} `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListRegionsResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyDataTags) GetTagCode() *string {
	return s.TagCode
}

func (s *ListRegionsResponseBodyDataTags) GetTagValue() interface{} {
	return s.TagValue
}

func (s *ListRegionsResponseBodyDataTags) SetTagCode(v string) *ListRegionsResponseBodyDataTags {
	s.TagCode = &v
	return s
}

func (s *ListRegionsResponseBodyDataTags) SetTagValue(v interface{}) *ListRegionsResponseBodyDataTags {
	s.TagValue = v
	return s
}

func (s *ListRegionsResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
