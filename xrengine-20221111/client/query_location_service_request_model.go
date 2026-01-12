// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLocationServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *QueryLocationServiceRequest
	GetId() *int64
	SetJwtToken(v string) *QueryLocationServiceRequest
	GetJwtToken() *string
}

type QueryLocationServiceRequest struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s QueryLocationServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLocationServiceRequest) GoString() string {
	return s.String()
}

func (s *QueryLocationServiceRequest) GetId() *int64 {
	return s.Id
}

func (s *QueryLocationServiceRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *QueryLocationServiceRequest) SetId(v int64) *QueryLocationServiceRequest {
	s.Id = &v
	return s
}

func (s *QueryLocationServiceRequest) SetJwtToken(v string) *QueryLocationServiceRequest {
	s.JwtToken = &v
	return s
}

func (s *QueryLocationServiceRequest) Validate() error {
	return dara.Validate(s)
}
