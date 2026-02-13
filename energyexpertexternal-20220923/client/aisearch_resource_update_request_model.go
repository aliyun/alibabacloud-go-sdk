// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchResourceUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *AISearchResourceUpdateRequest
	GetData() interface{}
	SetResourceId(v string) *AISearchResourceUpdateRequest
	GetResourceId() *string
	SetType(v string) *AISearchResourceUpdateRequest
	GetType() *string
}

type AISearchResourceUpdateRequest struct {
	// example:
	//
	// {
	//
	//       "title": "update title"
	//
	//     }
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WzMGQZwB7nQEs3Qk3ajH
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// miniapp
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AISearchResourceUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceUpdateRequest) GoString() string {
	return s.String()
}

func (s *AISearchResourceUpdateRequest) GetData() interface{} {
	return s.Data
}

func (s *AISearchResourceUpdateRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *AISearchResourceUpdateRequest) GetType() *string {
	return s.Type
}

func (s *AISearchResourceUpdateRequest) SetData(v interface{}) *AISearchResourceUpdateRequest {
	s.Data = v
	return s
}

func (s *AISearchResourceUpdateRequest) SetResourceId(v string) *AISearchResourceUpdateRequest {
	s.ResourceId = &v
	return s
}

func (s *AISearchResourceUpdateRequest) SetType(v string) *AISearchResourceUpdateRequest {
	s.Type = &v
	return s
}

func (s *AISearchResourceUpdateRequest) Validate() error {
	return dara.Validate(s)
}
