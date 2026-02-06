// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetVersionResponseBody) *GetVersionResponse
	GetBody() *GetVersionResponseBody
}

type GetVersionResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVersionResponse) GoString() string {
	return s.String()
}

func (s *GetVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVersionResponse) GetBody() *GetVersionResponseBody {
	return s.Body
}

func (s *GetVersionResponse) SetHeaders(v map[string]*string) *GetVersionResponse {
	s.Headers = v
	return s
}

func (s *GetVersionResponse) SetStatusCode(v int32) *GetVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVersionResponse) SetBody(v *GetVersionResponseBody) *GetVersionResponse {
	s.Body = v
	return s
}

func (s *GetVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
