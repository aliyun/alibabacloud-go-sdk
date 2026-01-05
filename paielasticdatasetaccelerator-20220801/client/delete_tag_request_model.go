// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *DeleteTagRequest
	GetKey() *string
	SetResourceId(v string) *DeleteTagRequest
	GetResourceId() *string
	SetResourceType(v string) *DeleteTagRequest
	GetResourceType() *string
}

type DeleteTagRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dataset_version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// inst-my1tk3jggooi5uwks
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DeleteTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagRequest) GetKey() *string {
	return s.Key
}

func (s *DeleteTagRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DeleteTagRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteTagRequest) SetKey(v string) *DeleteTagRequest {
	s.Key = &v
	return s
}

func (s *DeleteTagRequest) SetResourceId(v string) *DeleteTagRequest {
	s.ResourceId = &v
	return s
}

func (s *DeleteTagRequest) SetResourceType(v string) *DeleteTagRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteTagRequest) Validate() error {
	return dara.Validate(s)
}
