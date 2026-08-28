// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAgentSpecVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SubmitAgentSpecVersionRequestBody) *SubmitAgentSpecVersionRequest
	GetBody() *SubmitAgentSpecVersionRequestBody
}

type SubmitAgentSpecVersionRequest struct {
	// The request body.
	Body *SubmitAgentSpecVersionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s SubmitAgentSpecVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAgentSpecVersionRequest) GoString() string {
	return s.String()
}

func (s *SubmitAgentSpecVersionRequest) GetBody() *SubmitAgentSpecVersionRequestBody {
	return s.Body
}

func (s *SubmitAgentSpecVersionRequest) SetBody(v *SubmitAgentSpecVersionRequestBody) *SubmitAgentSpecVersionRequest {
	s.Body = v
	return s
}

func (s *SubmitAgentSpecVersionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAgentSpecVersionRequestBody struct {
}

func (s SubmitAgentSpecVersionRequestBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAgentSpecVersionRequestBody) GoString() string {
	return s.String()
}

func (s *SubmitAgentSpecVersionRequestBody) Validate() error {
	return dara.Validate(s)
}
