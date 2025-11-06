// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVhostRateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVhostRateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVhostRateResponse
	GetStatusCode() *int32
	SetBody(v *GetVhostRateResponseBody) *GetVhostRateResponse
	GetBody() *GetVhostRateResponseBody
}

type GetVhostRateResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVhostRateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVhostRateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVhostRateResponse) GoString() string {
	return s.String()
}

func (s *GetVhostRateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVhostRateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVhostRateResponse) GetBody() *GetVhostRateResponseBody {
	return s.Body
}

func (s *GetVhostRateResponse) SetHeaders(v map[string]*string) *GetVhostRateResponse {
	s.Headers = v
	return s
}

func (s *GetVhostRateResponse) SetStatusCode(v int32) *GetVhostRateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVhostRateResponse) SetBody(v *GetVhostRateResponseBody) *GetVhostRateResponse {
	s.Body = v
	return s
}

func (s *GetVhostRateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
