// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSeoBypassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSeoBypassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSeoBypassResponse
	GetStatusCode() *int32
	SetBody(v *GetSeoBypassResponseBody) *GetSeoBypassResponse
	GetBody() *GetSeoBypassResponseBody
}

type GetSeoBypassResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSeoBypassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSeoBypassResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSeoBypassResponse) GoString() string {
	return s.String()
}

func (s *GetSeoBypassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSeoBypassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSeoBypassResponse) GetBody() *GetSeoBypassResponseBody {
	return s.Body
}

func (s *GetSeoBypassResponse) SetHeaders(v map[string]*string) *GetSeoBypassResponse {
	s.Headers = v
	return s
}

func (s *GetSeoBypassResponse) SetStatusCode(v int32) *GetSeoBypassResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSeoBypassResponse) SetBody(v *GetSeoBypassResponseBody) *GetSeoBypassResponse {
	s.Body = v
	return s
}

func (s *GetSeoBypassResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
