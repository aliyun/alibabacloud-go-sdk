// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *ResumeSessionRequest
	GetQualifier() *string
}

type ResumeSessionRequest struct {
	// example:
	//
	// aliasName1
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s ResumeSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeSessionRequest) GoString() string {
	return s.String()
}

func (s *ResumeSessionRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *ResumeSessionRequest) SetQualifier(v string) *ResumeSessionRequest {
	s.Qualifier = &v
	return s
}

func (s *ResumeSessionRequest) Validate() error {
	return dara.Validate(s)
}
