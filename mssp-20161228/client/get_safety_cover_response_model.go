// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSafetyCoverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSafetyCoverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSafetyCoverResponse
	GetStatusCode() *int32
	SetBody(v *GetSafetyCoverResponseBody) *GetSafetyCoverResponse
	GetBody() *GetSafetyCoverResponseBody
}

type GetSafetyCoverResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSafetyCoverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSafetyCoverResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSafetyCoverResponse) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSafetyCoverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSafetyCoverResponse) GetBody() *GetSafetyCoverResponseBody {
	return s.Body
}

func (s *GetSafetyCoverResponse) SetHeaders(v map[string]*string) *GetSafetyCoverResponse {
	s.Headers = v
	return s
}

func (s *GetSafetyCoverResponse) SetStatusCode(v int32) *GetSafetyCoverResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSafetyCoverResponse) SetBody(v *GetSafetyCoverResponseBody) *GetSafetyCoverResponse {
	s.Body = v
	return s
}

func (s *GetSafetyCoverResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
