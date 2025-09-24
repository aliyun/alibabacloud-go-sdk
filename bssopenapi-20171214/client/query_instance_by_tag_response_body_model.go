// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceByTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryInstanceByTagResponseBody
	GetCode() *string
	SetMessage(v string) *QueryInstanceByTagResponseBody
	GetMessage() *string
	SetNextToken(v string) *QueryInstanceByTagResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryInstanceByTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryInstanceByTagResponseBody
	GetSuccess() *bool
	SetTagResource(v []*QueryInstanceByTagResponseBodyTagResource) *QueryInstanceByTagResponseBody
	GetTagResource() []*QueryInstanceByTagResponseBodyTagResource
}

type QueryInstanceByTagResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// PARAM_ERROR
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// param is null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token that determines the start point of the query. The return value is the value of the NextToken response parameter that was returned last time the QueryInstanceByTag operation was called.
	//
	// example:
	//
	// CAESEgoQCg4KCm
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9EC6C0B7-3397-5FAE-9915-8972CDDB1211
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The instances returned.
	TagResource []*QueryInstanceByTagResponseBodyTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s QueryInstanceByTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceByTagResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceByTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryInstanceByTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryInstanceByTagResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryInstanceByTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInstanceByTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryInstanceByTagResponseBody) GetTagResource() []*QueryInstanceByTagResponseBodyTagResource {
	return s.TagResource
}

func (s *QueryInstanceByTagResponseBody) SetCode(v string) *QueryInstanceByTagResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInstanceByTagResponseBody) SetMessage(v string) *QueryInstanceByTagResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstanceByTagResponseBody) SetNextToken(v string) *QueryInstanceByTagResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryInstanceByTagResponseBody) SetRequestId(v string) *QueryInstanceByTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstanceByTagResponseBody) SetSuccess(v bool) *QueryInstanceByTagResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInstanceByTagResponseBody) SetTagResource(v []*QueryInstanceByTagResponseBodyTagResource) *QueryInstanceByTagResponseBody {
	s.TagResource = v
	return s
}

func (s *QueryInstanceByTagResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryInstanceByTagResponseBodyTagResource struct {
	// The ID of the resource.
	//
	// example:
	//
	// spn-xxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. The returned resource type indicates a savings plan instance.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tag []*QueryInstanceByTagResponseBodyTagResourceTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s QueryInstanceByTagResponseBodyTagResource) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceByTagResponseBodyTagResource) GoString() string {
	return s.String()
}

func (s *QueryInstanceByTagResponseBodyTagResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryInstanceByTagResponseBodyTagResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryInstanceByTagResponseBodyTagResource) GetTag() []*QueryInstanceByTagResponseBodyTagResourceTag {
	return s.Tag
}

func (s *QueryInstanceByTagResponseBodyTagResource) SetResourceId(v string) *QueryInstanceByTagResponseBodyTagResource {
	s.ResourceId = &v
	return s
}

func (s *QueryInstanceByTagResponseBodyTagResource) SetResourceType(v string) *QueryInstanceByTagResponseBodyTagResource {
	s.ResourceType = &v
	return s
}

func (s *QueryInstanceByTagResponseBodyTagResource) SetTag(v []*QueryInstanceByTagResponseBodyTagResourceTag) *QueryInstanceByTagResponseBodyTagResource {
	s.Tag = v
	return s
}

func (s *QueryInstanceByTagResponseBodyTagResource) Validate() error {
	return dara.Validate(s)
}

type QueryInstanceByTagResponseBodyTagResourceTag struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryInstanceByTagResponseBodyTagResourceTag) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceByTagResponseBodyTagResourceTag) GoString() string {
	return s.String()
}

func (s *QueryInstanceByTagResponseBodyTagResourceTag) GetKey() *string {
	return s.Key
}

func (s *QueryInstanceByTagResponseBodyTagResourceTag) GetValue() *string {
	return s.Value
}

func (s *QueryInstanceByTagResponseBodyTagResourceTag) SetKey(v string) *QueryInstanceByTagResponseBodyTagResourceTag {
	s.Key = &v
	return s
}

func (s *QueryInstanceByTagResponseBodyTagResourceTag) SetValue(v string) *QueryInstanceByTagResponseBodyTagResourceTag {
	s.Value = &v
	return s
}

func (s *QueryInstanceByTagResponseBodyTagResourceTag) Validate() error {
	return dara.Validate(s)
}
