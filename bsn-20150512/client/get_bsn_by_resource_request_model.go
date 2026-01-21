// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBsnByResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliuid(v int64) *GetBsnByResourceRequest
	GetAliuid() *int64
	SetResourceId(v string) *GetBsnByResourceRequest
	GetResourceId() *string
	SetResourceType(v int32) *GetBsnByResourceRequest
	GetResourceType() *int32
}

type GetBsnByResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	Aliuid *int64 `json:"aliuid,omitempty" xml:"aliuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3097938
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alicloud_edas_application_scale
	ResourceType *int32 `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetBsnByResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBsnByResourceRequest) GoString() string {
	return s.String()
}

func (s *GetBsnByResourceRequest) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *GetBsnByResourceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetBsnByResourceRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *GetBsnByResourceRequest) SetAliuid(v int64) *GetBsnByResourceRequest {
	s.Aliuid = &v
	return s
}

func (s *GetBsnByResourceRequest) SetResourceId(v string) *GetBsnByResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *GetBsnByResourceRequest) SetResourceType(v int32) *GetBsnByResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *GetBsnByResourceRequest) Validate() error {
	return dara.Validate(s)
}
