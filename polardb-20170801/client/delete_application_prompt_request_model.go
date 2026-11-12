// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationPromptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeleteApplicationPromptRequest
	GetApplicationId() *string
	SetPromptId(v string) *DeleteApplicationPromptRequest
	GetPromptId() *string
}

type DeleteApplicationPromptRequest struct {
	// The application ID. The ID must point to an application with mem0 capabilities enabled. You can determine this by checking whether the MemApplicationAttribute returned by DescribeApplicationAttribute is not empty. Applications such as polarclaw, supabase, and polarclawsaas that do not have mem0 enabled do not support prompt operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The prompt ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// papt-f9lajgw765f4fnrzn1
	PromptId *string `json:"PromptId,omitempty" xml:"PromptId,omitempty"`
}

func (s DeleteApplicationPromptRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationPromptRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationPromptRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeleteApplicationPromptRequest) GetPromptId() *string {
	return s.PromptId
}

func (s *DeleteApplicationPromptRequest) SetApplicationId(v string) *DeleteApplicationPromptRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteApplicationPromptRequest) SetPromptId(v string) *DeleteApplicationPromptRequest {
	s.PromptId = &v
	return s
}

func (s *DeleteApplicationPromptRequest) Validate() error {
	return dara.Validate(s)
}
