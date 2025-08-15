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
	SetDynamicCode(v string) *ListTagResourcesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListTagResourcesResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListTagResourcesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTagResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTagResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTagResourcesResponseBody
	GetSuccess() *bool
}

type ListTagResourcesResponseBody struct {
	// Error code
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Return result
	Data *ListTagResourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Dynamic error code
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// Dynamic error message
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// The topic already exists.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID
	//
	// example:
	//
	// F00C6A70-C782-5DD6-9D11-0CFC710100C7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Whether the operation was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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

func (s *ListTagResourcesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListTagResourcesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListTagResourcesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
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

func (s *ListTagResourcesResponseBody) SetCode(v string) *ListTagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetData(v *ListTagResourcesResponseBodyData) *ListTagResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListTagResourcesResponseBody) SetDynamicCode(v string) *ListTagResourcesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetDynamicMessage(v string) *ListTagResourcesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetHttpStatusCode(v int32) *ListTagResourcesResponseBody {
	s.HttpStatusCode = &v
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

func (s *ListTagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTagResourcesResponseBodyData struct {
	// The position from which the next query starts.
	//
	// example:
	//
	// d09e2b63e1b12d905b7080ff70
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// F00C6A70-C782-5DD6-9D11-0CFC710100C7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Resource tag relationships.
	TagResources []*ListTagResourcesResponseBodyDataTagResources `json:"tagResources,omitempty" xml:"tagResources,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagResourcesResponseBodyData) GetTagResources() []*ListTagResourcesResponseBodyDataTagResources {
	return s.TagResources
}

func (s *ListTagResourcesResponseBodyData) SetNextToken(v string) *ListTagResourcesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBodyData) SetRequestId(v string) *ListTagResourcesResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBodyData) SetTagResources(v []*ListTagResourcesResponseBodyDataTagResources) *ListTagResourcesResponseBodyData {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListTagResourcesResponseBodyDataTagResources struct {
	// UID of the resource owner.
	//
	// example:
	//
	// 1876441048322426
	AliUid *int64 `json:"aliUid,omitempty" xml:"aliUid,omitempty"`
	// Tag category.
	//
	// example:
	//
	// custom
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Resource ID.
	//
	// example:
	//
	// rmq-cn-pe334n08h08
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// Visibility scope.
	//
	// example:
	//
	// public
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// Tag key.
	//
	// example:
	//
	// key
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// Tag value.
	//
	// example:
	//
	// value
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyDataTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyDataTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyDataTagResources) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListTagResourcesResponseBodyDataTagResources) GetCategory() *string {
	return s.Category
}

func (s *ListTagResourcesResponseBodyDataTagResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagResourcesResponseBodyDataTagResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesResponseBodyDataTagResources) GetScope() *string {
	return s.Scope
}

func (s *ListTagResourcesResponseBodyDataTagResources) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagResourcesResponseBodyDataTagResources) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetAliUid(v int64) *ListTagResourcesResponseBodyDataTagResources {
	s.AliUid = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetCategory(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.Category = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetScope(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.Scope = &v
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
