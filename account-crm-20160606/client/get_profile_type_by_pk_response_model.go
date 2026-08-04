// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProfileTypeByPkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProfileTypeByPkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProfileTypeByPkResponse
	GetStatusCode() *int32
	SetBody(v *GetProfileTypeByPkResponseBody) *GetProfileTypeByPkResponse
	GetBody() *GetProfileTypeByPkResponseBody
}

type GetProfileTypeByPkResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProfileTypeByPkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProfileTypeByPkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProfileTypeByPkResponse) GoString() string {
	return s.String()
}

func (s *GetProfileTypeByPkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProfileTypeByPkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProfileTypeByPkResponse) GetBody() *GetProfileTypeByPkResponseBody {
	return s.Body
}

func (s *GetProfileTypeByPkResponse) SetHeaders(v map[string]*string) *GetProfileTypeByPkResponse {
	s.Headers = v
	return s
}

func (s *GetProfileTypeByPkResponse) SetStatusCode(v int32) *GetProfileTypeByPkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProfileTypeByPkResponse) SetBody(v *GetProfileTypeByPkResponseBody) *GetProfileTypeByPkResponse {
	s.Body = v
	return s
}

func (s *GetProfileTypeByPkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
