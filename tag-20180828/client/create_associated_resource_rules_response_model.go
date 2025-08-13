// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAssociatedResourceRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAssociatedResourceRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAssociatedResourceRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateAssociatedResourceRulesResponseBody) *CreateAssociatedResourceRulesResponse
	GetBody() *CreateAssociatedResourceRulesResponseBody
}

type CreateAssociatedResourceRulesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAssociatedResourceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAssociatedResourceRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAssociatedResourceRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateAssociatedResourceRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAssociatedResourceRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAssociatedResourceRulesResponse) GetBody() *CreateAssociatedResourceRulesResponseBody {
	return s.Body
}

func (s *CreateAssociatedResourceRulesResponse) SetHeaders(v map[string]*string) *CreateAssociatedResourceRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateAssociatedResourceRulesResponse) SetStatusCode(v int32) *CreateAssociatedResourceRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAssociatedResourceRulesResponse) SetBody(v *CreateAssociatedResourceRulesResponseBody) *CreateAssociatedResourceRulesResponse {
	s.Body = v
	return s
}

func (s *CreateAssociatedResourceRulesResponse) Validate() error {
	return dara.Validate(s)
}
