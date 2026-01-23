// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStandardSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStandardSetResponse
	GetStatusCode() *int32
	SetBody(v *GetStandardSetResponseBody) *GetStandardSetResponse
	GetBody() *GetStandardSetResponseBody
}

type GetStandardSetResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStandardSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStandardSetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStandardSetResponse) GoString() string {
	return s.String()
}

func (s *GetStandardSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStandardSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStandardSetResponse) GetBody() *GetStandardSetResponseBody {
	return s.Body
}

func (s *GetStandardSetResponse) SetHeaders(v map[string]*string) *GetStandardSetResponse {
	s.Headers = v
	return s
}

func (s *GetStandardSetResponse) SetStatusCode(v int32) *GetStandardSetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStandardSetResponse) SetBody(v *GetStandardSetResponseBody) *GetStandardSetResponse {
	s.Body = v
	return s
}

func (s *GetStandardSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
