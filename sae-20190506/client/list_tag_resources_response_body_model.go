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
	// The HTTP status code.
	//
	// - **2xx*	- indicates that the request was successful.
	//
	// - **3xx*	- indicates that the request was redirected.
	//
	// - **4xx*	- indicates that a client-side error occurred.
	//
	// - **5xx*	- indicates that a server-side error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ListTagResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - This parameter is returned only if the request fails.
	//
	// - For more information, see the **Error codes*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned for the request.
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, an error message is returned.
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
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of a request.
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
	// The token that is used to retrieve the next page of results. A query returns a maximum of 50 results. If the results are truncated, you can use this token in a subsequent request to retrieve the next page of results.
	//
	// example:
	//
	// ""
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of tags and their associated resources.
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
	// The ID of the application.
	//
	// example:
	//
	// d42921c4-5433-4abd-8075-0e536f8b****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. The value is fixed as `application`.
	//
	// example:
	//
	// ALIYUN::SAE::APPLICATION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// k1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
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
