// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRuntimeModelTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelTemplateId(v string) *RemoveRuntimeModelTemplateRequest
	GetModelTemplateId() *string
	SetRuntimeIds(v []*string) *RemoveRuntimeModelTemplateRequest
	GetRuntimeIds() []*string
	SetRuntimeType(v string) *RemoveRuntimeModelTemplateRequest
	GetRuntimeType() *string
}

type RemoveRuntimeModelTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mt-xxxxxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// This parameter is required.
	RuntimeIds []*string `json:"RuntimeIds,omitempty" xml:"RuntimeIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// CloudDesktop
	RuntimeType *string `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
}

func (s RemoveRuntimeModelTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveRuntimeModelTemplateRequest) GoString() string {
	return s.String()
}

func (s *RemoveRuntimeModelTemplateRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *RemoveRuntimeModelTemplateRequest) GetRuntimeIds() []*string {
	return s.RuntimeIds
}

func (s *RemoveRuntimeModelTemplateRequest) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *RemoveRuntimeModelTemplateRequest) SetModelTemplateId(v string) *RemoveRuntimeModelTemplateRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *RemoveRuntimeModelTemplateRequest) SetRuntimeIds(v []*string) *RemoveRuntimeModelTemplateRequest {
	s.RuntimeIds = v
	return s
}

func (s *RemoveRuntimeModelTemplateRequest) SetRuntimeType(v string) *RemoveRuntimeModelTemplateRequest {
	s.RuntimeType = &v
	return s
}

func (s *RemoveRuntimeModelTemplateRequest) Validate() error {
	return dara.Validate(s)
}
