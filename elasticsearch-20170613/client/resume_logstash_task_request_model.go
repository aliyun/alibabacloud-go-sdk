// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeLogstashTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ResumeLogstashTaskRequest
	GetClientToken() *string
}

type ResumeLogstashTaskRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s ResumeLogstashTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeLogstashTaskRequest) GoString() string {
	return s.String()
}

func (s *ResumeLogstashTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ResumeLogstashTaskRequest) SetClientToken(v string) *ResumeLogstashTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *ResumeLogstashTaskRequest) Validate() error {
	return dara.Validate(s)
}
