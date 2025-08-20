// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *GetResourceTypeRequest
	GetResourceType() *string
	SetVersionId(v string) *GetResourceTypeRequest
	GetVersionId() *string
}

type GetResourceTypeRequest struct {
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN::ROS::WaitConditionHandle
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The version ID. If you want to query a specific version of the resource type, you must specify this parameter. If you do not specify this parameter, only the resource type is queried.
	//
	// > This parameter is supported only for modules.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetResourceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *GetResourceTypeRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceTypeRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetResourceTypeRequest) SetResourceType(v string) *GetResourceTypeRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceTypeRequest) SetVersionId(v string) *GetResourceTypeRequest {
	s.VersionId = &v
	return s
}

func (s *GetResourceTypeRequest) Validate() error {
	return dara.Validate(s)
}
