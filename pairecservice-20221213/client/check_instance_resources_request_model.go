// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *CheckInstanceResourcesRequest
	GetResourceId() *string
	SetType(v string) *CheckInstanceResourcesRequest
	GetType() *string
	SetUri(v string) *CheckInstanceResourcesRequest
	GetUri() *string
}

type CheckInstanceResourcesRequest struct {
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// bucket-test-123
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CheckInstanceResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *CheckInstanceResourcesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *CheckInstanceResourcesRequest) GetType() *string {
	return s.Type
}

func (s *CheckInstanceResourcesRequest) GetUri() *string {
	return s.Uri
}

func (s *CheckInstanceResourcesRequest) SetResourceId(v string) *CheckInstanceResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *CheckInstanceResourcesRequest) SetType(v string) *CheckInstanceResourcesRequest {
	s.Type = &v
	return s
}

func (s *CheckInstanceResourcesRequest) SetUri(v string) *CheckInstanceResourcesRequest {
	s.Uri = &v
	return s
}

func (s *CheckInstanceResourcesRequest) Validate() error {
	return dara.Validate(s)
}
