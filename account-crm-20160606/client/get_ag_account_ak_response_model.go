// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgAccountAkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgAccountAkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgAccountAkResponse
	GetStatusCode() *int32
	SetBody(v *GetAgAccountAkResponseBody) *GetAgAccountAkResponse
	GetBody() *GetAgAccountAkResponseBody
}

type GetAgAccountAkResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgAccountAkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgAccountAkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgAccountAkResponse) GoString() string {
	return s.String()
}

func (s *GetAgAccountAkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgAccountAkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgAccountAkResponse) GetBody() *GetAgAccountAkResponseBody {
	return s.Body
}

func (s *GetAgAccountAkResponse) SetHeaders(v map[string]*string) *GetAgAccountAkResponse {
	s.Headers = v
	return s
}

func (s *GetAgAccountAkResponse) SetStatusCode(v int32) *GetAgAccountAkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgAccountAkResponse) SetBody(v *GetAgAccountAkResponseBody) *GetAgAccountAkResponse {
	s.Body = v
	return s
}

func (s *GetAgAccountAkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
