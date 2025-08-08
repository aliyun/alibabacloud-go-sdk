// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallation interface {
	dara.Model
	String() string
	GoString() string
	SetActionUri(v string) *Installation
	GetActionUri() *string
	SetMessage(v string) *Installation
	GetMessage() *string
	SetStage(v string) *Installation
	GetStage() *string
}

type Installation struct {
	// example:
	//
	// https://github.com/login/oauth/authorize?client_id=86059a1b2bb20d3e5fc3&scope=repo,repo:status,delete_repo
	ActionUri *string `json:"actionUri,omitempty" xml:"actionUri,omitempty"`
	// example:
	//
	// Please click \"actionUri\" to complete the OAuth authorization process
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// finished
	Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
}

func (s Installation) String() string {
	return dara.Prettify(s)
}

func (s Installation) GoString() string {
	return s.String()
}

func (s *Installation) GetActionUri() *string {
	return s.ActionUri
}

func (s *Installation) GetMessage() *string {
	return s.Message
}

func (s *Installation) GetStage() *string {
	return s.Stage
}

func (s *Installation) SetActionUri(v string) *Installation {
	s.ActionUri = &v
	return s
}

func (s *Installation) SetMessage(v string) *Installation {
	s.Message = &v
	return s
}

func (s *Installation) SetStage(v string) *Installation {
	s.Stage = &v
	return s
}

func (s *Installation) Validate() error {
	return dara.Validate(s)
}
