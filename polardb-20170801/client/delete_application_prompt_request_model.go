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
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
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
