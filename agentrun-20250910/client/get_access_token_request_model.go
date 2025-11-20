// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *GetAccessTokenRequest
	GetResourceId() *string
	SetResourceName(v string) *GetAccessTokenRequest
	GetResourceName() *string
	SetResourceType(v string) *GetAccessTokenRequest
	GetResourceType() *string
}

type GetAccessTokenRequest struct {
	// example:
	//
	// mod-kw1750tjqs94n9k37o5hjk
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// COUPON_OPERATION_CALLBACK
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// example:
	//
	// SceneTestingTask
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetAccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAccessTokenRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetAccessTokenRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetAccessTokenRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAccessTokenRequest) SetResourceId(v string) *GetAccessTokenRequest {
	s.ResourceId = &v
	return s
}

func (s *GetAccessTokenRequest) SetResourceName(v string) *GetAccessTokenRequest {
	s.ResourceName = &v
	return s
}

func (s *GetAccessTokenRequest) SetResourceType(v string) *GetAccessTokenRequest {
	s.ResourceType = &v
	return s
}

func (s *GetAccessTokenRequest) Validate() error {
	return dara.Validate(s)
}
