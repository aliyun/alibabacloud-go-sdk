// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApisecEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApisecEventsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApisecEventsResponseBody) *ModifyApisecEventsResponse
	GetBody() *ModifyApisecEventsResponseBody
}

type ModifyApisecEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApisecEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApisecEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecEventsResponse) GoString() string {
	return s.String()
}

func (s *ModifyApisecEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApisecEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApisecEventsResponse) GetBody() *ModifyApisecEventsResponseBody {
	return s.Body
}

func (s *ModifyApisecEventsResponse) SetHeaders(v map[string]*string) *ModifyApisecEventsResponse {
	s.Headers = v
	return s
}

func (s *ModifyApisecEventsResponse) SetStatusCode(v int32) *ModifyApisecEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApisecEventsResponse) SetBody(v *ModifyApisecEventsResponseBody) *ModifyApisecEventsResponse {
	s.Body = v
	return s
}

func (s *ModifyApisecEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
