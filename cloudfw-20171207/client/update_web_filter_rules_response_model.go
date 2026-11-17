// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebFilterRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWebFilterRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWebFilterRulesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWebFilterRulesResponseBody) *UpdateWebFilterRulesResponse
	GetBody() *UpdateWebFilterRulesResponseBody
}

type UpdateWebFilterRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWebFilterRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebFilterRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebFilterRulesResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebFilterRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWebFilterRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWebFilterRulesResponse) GetBody() *UpdateWebFilterRulesResponseBody {
	return s.Body
}

func (s *UpdateWebFilterRulesResponse) SetHeaders(v map[string]*string) *UpdateWebFilterRulesResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebFilterRulesResponse) SetStatusCode(v int32) *UpdateWebFilterRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebFilterRulesResponse) SetBody(v *UpdateWebFilterRulesResponseBody) *UpdateWebFilterRulesResponse {
	s.Body = v
	return s
}

func (s *UpdateWebFilterRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
