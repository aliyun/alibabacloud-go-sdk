// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendLocationServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *SuspendLocationServiceRequest
	GetId() *int64
	SetJwtToken(v string) *SuspendLocationServiceRequest
	GetJwtToken() *string
}

type SuspendLocationServiceRequest struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s SuspendLocationServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendLocationServiceRequest) GoString() string {
	return s.String()
}

func (s *SuspendLocationServiceRequest) GetId() *int64 {
	return s.Id
}

func (s *SuspendLocationServiceRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *SuspendLocationServiceRequest) SetId(v int64) *SuspendLocationServiceRequest {
	s.Id = &v
	return s
}

func (s *SuspendLocationServiceRequest) SetJwtToken(v string) *SuspendLocationServiceRequest {
	s.JwtToken = &v
	return s
}

func (s *SuspendLocationServiceRequest) Validate() error {
	return dara.Validate(s)
}
