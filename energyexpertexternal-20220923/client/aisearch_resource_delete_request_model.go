// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchResourceDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *AISearchResourceDeleteRequest
	GetResourceId() *string
	SetType(v string) *AISearchResourceDeleteRequest
	GetType() *string
}

type AISearchResourceDeleteRequest struct {
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

func (s AISearchResourceDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceDeleteRequest) GoString() string {
	return s.String()
}

func (s *AISearchResourceDeleteRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *AISearchResourceDeleteRequest) GetType() *string {
	return s.Type
}

func (s *AISearchResourceDeleteRequest) SetResourceId(v string) *AISearchResourceDeleteRequest {
	s.ResourceId = &v
	return s
}

func (s *AISearchResourceDeleteRequest) SetType(v string) *AISearchResourceDeleteRequest {
	s.Type = &v
	return s
}

func (s *AISearchResourceDeleteRequest) Validate() error {
	return dara.Validate(s)
}
