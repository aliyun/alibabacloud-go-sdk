// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *PauseSessionRequest
	GetQualifier() *string
}

type PauseSessionRequest struct {
	// example:
	//
	// aliasName1
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s PauseSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseSessionRequest) GoString() string {
	return s.String()
}

func (s *PauseSessionRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *PauseSessionRequest) SetQualifier(v string) *PauseSessionRequest {
	s.Qualifier = &v
	return s
}

func (s *PauseSessionRequest) Validate() error {
	return dara.Validate(s)
}
