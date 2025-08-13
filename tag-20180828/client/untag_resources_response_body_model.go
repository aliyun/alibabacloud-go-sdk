// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedResources(v *UntagResourcesResponseBodyFailedResources) *UntagResourcesResponseBody
	GetFailedResources() *UntagResourcesResponseBodyFailedResources
	SetRequestId(v string) *UntagResourcesResponseBody
	GetRequestId() *string
}

type UntagResourcesResponseBody struct {
	// The information about the resources from which tags fail to be removed.
	//
	// >
	//
	// 	- If tags are removed from all resources, the value of FailedResources is empty.
	//
	// 	- If tags fail to be removed from some or all resources, the value of FailedResources contains the detailed information about the resources.
	FailedResources *UntagResourcesResponseBodyFailedResources `json:"FailedResources,omitempty" xml:"FailedResources,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 43D12436-B10F-4469-8136-FD1C5D2B2083
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) GetFailedResources() *UntagResourcesResponseBodyFailedResources {
	return s.FailedResources
}

func (s *UntagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagResourcesResponseBody) SetFailedResources(v *UntagResourcesResponseBodyFailedResources) *UntagResourcesResponseBody {
	s.FailedResources = v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type UntagResourcesResponseBodyFailedResources struct {
	FailedResource []*UntagResourcesResponseBodyFailedResourcesFailedResource `json:"FailedResource,omitempty" xml:"FailedResource,omitempty" type:"Repeated"`
}

func (s UntagResourcesResponseBodyFailedResources) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesResponseBodyFailedResources) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBodyFailedResources) GetFailedResource() []*UntagResourcesResponseBodyFailedResourcesFailedResource {
	return s.FailedResource
}

func (s *UntagResourcesResponseBodyFailedResources) SetFailedResource(v []*UntagResourcesResponseBodyFailedResourcesFailedResource) *UntagResourcesResponseBodyFailedResources {
	s.FailedResource = v
	return s
}

func (s *UntagResourcesResponseBodyFailedResources) Validate() error {
	return dara.Validate(s)
}

type UntagResourcesResponseBodyFailedResourcesFailedResource struct {
	// The ARN of the resource.
	//
	// example:
	//
	// arn:acs:ecs:cn-hangzhou:123456789****:instance/i-xxxxxxxxxx1
	ResourceARN *string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	// The information about the error.
	Result *UntagResourcesResponseBodyFailedResourcesFailedResourceResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UntagResourcesResponseBodyFailedResourcesFailedResource) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesResponseBodyFailedResourcesFailedResource) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResource) GetResourceARN() *string {
	return s.ResourceARN
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResource) GetResult() *UntagResourcesResponseBodyFailedResourcesFailedResourceResult {
	return s.Result
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResource) SetResourceARN(v string) *UntagResourcesResponseBodyFailedResourcesFailedResource {
	s.ResourceARN = &v
	return s
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResource) SetResult(v *UntagResourcesResponseBodyFailedResourcesFailedResourceResult) *UntagResourcesResponseBodyFailedResourcesFailedResource {
	s.Result = v
	return s
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResource) Validate() error {
	return dara.Validate(s)
}

type UntagResourcesResponseBodyFailedResourcesFailedResourceResult struct {
	// The error code.
	//
	// example:
	//
	// InvalidResourceId.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified ResourceIds are not found in our records.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UntagResourcesResponseBodyFailedResourcesFailedResourceResult) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesResponseBodyFailedResourcesFailedResourceResult) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResourceResult) GetCode() *string {
	return s.Code
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResourceResult) GetMessage() *string {
	return s.Message
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResourceResult) SetCode(v string) *UntagResourcesResponseBodyFailedResourcesFailedResourceResult {
	s.Code = &v
	return s
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResourceResult) SetMessage(v string) *UntagResourcesResponseBodyFailedResourcesFailedResourceResult {
	s.Message = &v
	return s
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResourceResult) Validate() error {
	return dara.Validate(s)
}
