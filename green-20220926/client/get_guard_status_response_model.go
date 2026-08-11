// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGuardStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGuardStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGuardStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetGuardStatusResponseBody) *GetGuardStatusResponse
	GetBody() *GetGuardStatusResponseBody
}

type GetGuardStatusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGuardStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGuardStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGuardStatusResponse) GoString() string {
	return s.String()
}

func (s *GetGuardStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGuardStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGuardStatusResponse) GetBody() *GetGuardStatusResponseBody {
	return s.Body
}

func (s *GetGuardStatusResponse) SetHeaders(v map[string]*string) *GetGuardStatusResponse {
	s.Headers = v
	return s
}

func (s *GetGuardStatusResponse) SetStatusCode(v int32) *GetGuardStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGuardStatusResponse) SetBody(v *GetGuardStatusResponseBody) *GetGuardStatusResponse {
	s.Body = v
	return s
}

func (s *GetGuardStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
