// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebFilterRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWebFilterRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWebFilterRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateWebFilterRulesResponseBody) *CreateWebFilterRulesResponse
	GetBody() *CreateWebFilterRulesResponseBody
}

type CreateWebFilterRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWebFilterRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWebFilterRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWebFilterRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateWebFilterRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWebFilterRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWebFilterRulesResponse) GetBody() *CreateWebFilterRulesResponseBody {
	return s.Body
}

func (s *CreateWebFilterRulesResponse) SetHeaders(v map[string]*string) *CreateWebFilterRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateWebFilterRulesResponse) SetStatusCode(v int32) *CreateWebFilterRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWebFilterRulesResponse) SetBody(v *CreateWebFilterRulesResponseBody) *CreateWebFilterRulesResponse {
	s.Body = v
	return s
}

func (s *CreateWebFilterRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
