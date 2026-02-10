// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDigitalSignByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDigitalSignByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDigitalSignByNameResponse
	GetStatusCode() *int32
	SetBody(v *QueryDigitalSignByNameResponseBody) *QueryDigitalSignByNameResponse
	GetBody() *QueryDigitalSignByNameResponseBody
}

type QueryDigitalSignByNameResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDigitalSignByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDigitalSignByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDigitalSignByNameResponse) GoString() string {
	return s.String()
}

func (s *QueryDigitalSignByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDigitalSignByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDigitalSignByNameResponse) GetBody() *QueryDigitalSignByNameResponseBody {
	return s.Body
}

func (s *QueryDigitalSignByNameResponse) SetHeaders(v map[string]*string) *QueryDigitalSignByNameResponse {
	s.Headers = v
	return s
}

func (s *QueryDigitalSignByNameResponse) SetStatusCode(v int32) *QueryDigitalSignByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDigitalSignByNameResponse) SetBody(v *QueryDigitalSignByNameResponseBody) *QueryDigitalSignByNameResponse {
	s.Body = v
	return s
}

func (s *QueryDigitalSignByNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
