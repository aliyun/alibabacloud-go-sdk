// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerbose(v string) *GetChatRequest
	GetVerbose() *string
}

type GetChatRequest struct {
	// example:
	//
	// QuestionAndAnswer
	Verbose *string `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetChatRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatRequest) GoString() string {
	return s.String()
}

func (s *GetChatRequest) GetVerbose() *string {
	return s.Verbose
}

func (s *GetChatRequest) SetVerbose(v string) *GetChatRequest {
	s.Verbose = &v
	return s
}

func (s *GetChatRequest) Validate() error {
	return dara.Validate(s)
}
