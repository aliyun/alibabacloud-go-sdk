// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeLocationServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ResumeLocationServiceRequest
	GetId() *int64
	SetJwtToken(v string) *ResumeLocationServiceRequest
	GetJwtToken() *string
}

type ResumeLocationServiceRequest struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s ResumeLocationServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeLocationServiceRequest) GoString() string {
	return s.String()
}

func (s *ResumeLocationServiceRequest) GetId() *int64 {
	return s.Id
}

func (s *ResumeLocationServiceRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *ResumeLocationServiceRequest) SetId(v int64) *ResumeLocationServiceRequest {
	s.Id = &v
	return s
}

func (s *ResumeLocationServiceRequest) SetJwtToken(v string) *ResumeLocationServiceRequest {
	s.JwtToken = &v
	return s
}

func (s *ResumeLocationServiceRequest) Validate() error {
	return dara.Validate(s)
}
