// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFormationInstsByTaskIDResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFormationInstsByTaskIDResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFormationInstsByTaskIDResponse
	GetStatusCode() *int32
	SetBody(v *QueryFormationInstsByTaskIDResponseBody) *QueryFormationInstsByTaskIDResponse
	GetBody() *QueryFormationInstsByTaskIDResponseBody
}

type QueryFormationInstsByTaskIDResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFormationInstsByTaskIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFormationInstsByTaskIDResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFormationInstsByTaskIDResponse) GoString() string {
	return s.String()
}

func (s *QueryFormationInstsByTaskIDResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFormationInstsByTaskIDResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFormationInstsByTaskIDResponse) GetBody() *QueryFormationInstsByTaskIDResponseBody {
	return s.Body
}

func (s *QueryFormationInstsByTaskIDResponse) SetHeaders(v map[string]*string) *QueryFormationInstsByTaskIDResponse {
	s.Headers = v
	return s
}

func (s *QueryFormationInstsByTaskIDResponse) SetStatusCode(v int32) *QueryFormationInstsByTaskIDResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFormationInstsByTaskIDResponse) SetBody(v *QueryFormationInstsByTaskIDResponseBody) *QueryFormationInstsByTaskIDResponse {
	s.Body = v
	return s
}

func (s *QueryFormationInstsByTaskIDResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
