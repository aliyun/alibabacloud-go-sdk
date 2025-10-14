// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafPhasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWafPhasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWafPhasesResponse
	GetStatusCode() *int32
	SetBody(v *ListWafPhasesResponseBody) *ListWafPhasesResponse
	GetBody() *ListWafPhasesResponseBody
}

type ListWafPhasesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWafPhasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWafPhasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWafPhasesResponse) GoString() string {
	return s.String()
}

func (s *ListWafPhasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWafPhasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWafPhasesResponse) GetBody() *ListWafPhasesResponseBody {
	return s.Body
}

func (s *ListWafPhasesResponse) SetHeaders(v map[string]*string) *ListWafPhasesResponse {
	s.Headers = v
	return s
}

func (s *ListWafPhasesResponse) SetStatusCode(v int32) *ListWafPhasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWafPhasesResponse) SetBody(v *ListWafPhasesResponseBody) *ListWafPhasesResponse {
	s.Body = v
	return s
}

func (s *ListWafPhasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
