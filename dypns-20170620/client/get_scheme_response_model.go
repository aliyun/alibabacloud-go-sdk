// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSchemeResponse
	GetStatusCode() *int32
	SetBody(v *GetSchemeResponseBody) *GetSchemeResponse
	GetBody() *GetSchemeResponseBody
}

type GetSchemeResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSchemeResponse) GoString() string {
	return s.String()
}

func (s *GetSchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSchemeResponse) GetBody() *GetSchemeResponseBody {
	return s.Body
}

func (s *GetSchemeResponse) SetHeaders(v map[string]*string) *GetSchemeResponse {
	s.Headers = v
	return s
}

func (s *GetSchemeResponse) SetStatusCode(v int32) *GetSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSchemeResponse) SetBody(v *GetSchemeResponseBody) *GetSchemeResponse {
	s.Body = v
	return s
}

func (s *GetSchemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
