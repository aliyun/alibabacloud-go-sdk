// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTagResourcesResponseBody
	GetCode() *string
	SetData(v *ListTagResourcesResponseBodyData) *ListTagResourcesResponseBody
	GetData() *ListTagResourcesResponseBodyData
	SetErrorCode(v string) *ListTagResourcesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListTagResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTagResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTagResourcesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListTagResourcesResponseBody
	GetTraceId() *string
}

type ListTagResourcesResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirect.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListTagResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidParameter.Obviously
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information.
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If an exception occurs, a specific error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7414187F-4F59-4585-9BCF-5F0804E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the mapping between applications and tags was successfully queried. Valid values:
	//
	// - **true**: The query was successful.
	//
	// - **false**: The query failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID. You can use this ID to precisely query the call information.
	//
	// example:
	//
	// 0bc5f84e15916043198032146d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTagResourcesResponseBody) GetData() *ListTagResourcesResponseBodyData {
	return s.Data
}

func (s *ListTagResourcesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTagResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTagResourcesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListTagResourcesResponseBody) SetCode(v string) *ListTagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetData(v *ListTagResourcesResponseBodyData) *ListTagResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListTagResourcesResponseBody) SetErrorCode(v string) *ListTagResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetMessage(v string) *ListTagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetSuccess(v bool) *ListTagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTraceId(v string) *ListTagResourcesResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListTagResourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTagResourcesResponseBodyData struct {
	// The pagination token. A single query returns up to 50 results. If the total number of results exceeds 50, a NextToken is returned. You can use this token in the next query to retrieve the remaining results.
	//
	// example:
	//
	// ""
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The mapping between applications and tags.
	TagResources []*ListTagResourcesResponseBodyDataTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesResponseBodyData) GetTagResources() []*ListTagResourcesResponseBodyDataTagResources {
	return s.TagResources
}

func (s *ListTagResourcesResponseBodyData) SetNextToken(v string) *ListTagResourcesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBodyData) SetTagResources(v []*ListTagResourcesResponseBodyDataTagResources) *ListTagResourcesResponseBodyData {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBodyData) Validate() error {
	if s.TagResources != nil {
		for _, item := range s.TagResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagResourcesResponseBodyDataTagResources struct {
	// The application ID.
	//
	// example:
	//
	// d42921c4-5433-4abd-8075-0e536f8b****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Only `application` is supported.
	//
	// example:
	//
	// ALIYUN::SAE::APPLICATION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// k1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyDataTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyDataTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyDataTagResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagResourcesResponseBodyDataTagResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesResponseBodyDataTagResources) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagResourcesResponseBodyDataTagResources) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) Validate() error {
	return dara.Validate(s)
}
