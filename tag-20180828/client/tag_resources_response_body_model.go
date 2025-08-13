// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedResources(v *TagResourcesResponseBodyFailedResources) *TagResourcesResponseBody
	GetFailedResources() *TagResourcesResponseBodyFailedResources
	SetRequestId(v string) *TagResourcesResponseBody
	GetRequestId() *string
}

type TagResourcesResponseBody struct {
	// The information about the resources to which tags fail to be added.
	//
	// >
	//
	// 	- If tags are added to all resources, the value of `FailedResources` is empty.
	//
	// 	- If tags fail to be added to some or all resources, the value of `FailedResources` contains the detailed information about the resources.
	FailedResources *TagResourcesResponseBodyFailedResources `json:"FailedResources,omitempty" xml:"FailedResources,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 72086426-9F8C-4A60-852B-864048FD1199
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) GetFailedResources() *TagResourcesResponseBodyFailedResources {
	return s.FailedResources
}

func (s *TagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagResourcesResponseBody) SetFailedResources(v *TagResourcesResponseBodyFailedResources) *TagResourcesResponseBody {
	s.FailedResources = v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type TagResourcesResponseBodyFailedResources struct {
	FailedResource []*TagResourcesResponseBodyFailedResourcesFailedResource `json:"FailedResource,omitempty" xml:"FailedResource,omitempty" type:"Repeated"`
}

func (s TagResourcesResponseBodyFailedResources) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesResponseBodyFailedResources) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBodyFailedResources) GetFailedResource() []*TagResourcesResponseBodyFailedResourcesFailedResource {
	return s.FailedResource
}

func (s *TagResourcesResponseBodyFailedResources) SetFailedResource(v []*TagResourcesResponseBodyFailedResourcesFailedResource) *TagResourcesResponseBodyFailedResources {
	s.FailedResource = v
	return s
}

func (s *TagResourcesResponseBodyFailedResources) Validate() error {
	return dara.Validate(s)
}

type TagResourcesResponseBodyFailedResourcesFailedResource struct {
	// The ARN of the resource.
	//
	// example:
	//
	// arn:acs:vpc:cn-hangzhou:123456789****:vpc/vpc-bp19dd90tkt6tz7wu****
	ResourceARN *string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	// The information about the error.
	Result *TagResourcesResponseBodyFailedResourcesFailedResourceResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s TagResourcesResponseBodyFailedResourcesFailedResource) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesResponseBodyFailedResourcesFailedResource) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResource) GetResourceARN() *string {
	return s.ResourceARN
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResource) GetResult() *TagResourcesResponseBodyFailedResourcesFailedResourceResult {
	return s.Result
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResource) SetResourceARN(v string) *TagResourcesResponseBodyFailedResourcesFailedResource {
	s.ResourceARN = &v
	return s
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResource) SetResult(v *TagResourcesResponseBodyFailedResourcesFailedResourceResult) *TagResourcesResponseBodyFailedResourcesFailedResource {
	s.Result = v
	return s
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResource) Validate() error {
	return dara.Validate(s)
}

type TagResourcesResponseBodyFailedResourcesFailedResourceResult struct {
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

func (s TagResourcesResponseBodyFailedResourcesFailedResourceResult) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesResponseBodyFailedResourcesFailedResourceResult) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResourceResult) GetCode() *string {
	return s.Code
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResourceResult) GetMessage() *string {
	return s.Message
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResourceResult) SetCode(v string) *TagResourcesResponseBodyFailedResourcesFailedResourceResult {
	s.Code = &v
	return s
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResourceResult) SetMessage(v string) *TagResourcesResponseBodyFailedResourcesFailedResourceResult {
	s.Message = &v
	return s
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResourceResult) Validate() error {
	return dara.Validate(s)
}
