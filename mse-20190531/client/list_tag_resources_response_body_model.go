// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListTagResourcesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListTagResourcesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTagResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTagResourcesResponseBody
	GetSuccess() *bool
	SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody
	GetTagResources() *ListTagResourcesResponseBodyTagResources
}

type ListTagResourcesResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token that is required for the next query. If the NextToken parameter is empty, no subsequent query is required.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6ng3BeUyo2zbY5wgL1DBa2T
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 92ED7C9C-84A2-595F-87FE-354F84C0F739
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The list of resources.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTagResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTagResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTagResourcesResponseBody) GetTagResources() *ListTagResourcesResponseBodyTagResources {
	return s.TagResources
}

func (s *ListTagResourcesResponseBody) SetErrorCode(v string) *ListTagResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetMessage(v string) *ListTagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
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

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBody) Validate() error {
	if s.TagResources != nil {
		if err := s.TagResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) GetTagResource() []*ListTagResourcesResponseBodyTagResourcesTagResource {
	return s.TagResource
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) Validate() error {
	if s.TagResource != nil {
		for _, item := range s.TagResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// The ID of the resource.
	//
	// example:
	//
	// mse-cn-7pp2st29t16
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// env
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// prd
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) Validate() error {
	return dara.Validate(s)
}
