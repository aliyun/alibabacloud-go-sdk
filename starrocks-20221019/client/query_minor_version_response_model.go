// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinorVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMinorVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMinorVersionResponse
	GetStatusCode() *int32
	SetBody(v *QueryMinorVersionResponseBody) *QueryMinorVersionResponse
	GetBody() *QueryMinorVersionResponseBody
}

type QueryMinorVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinorVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinorVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMinorVersionResponse) GoString() string {
	return s.String()
}

func (s *QueryMinorVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMinorVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMinorVersionResponse) GetBody() *QueryMinorVersionResponseBody {
	return s.Body
}

func (s *QueryMinorVersionResponse) SetHeaders(v map[string]*string) *QueryMinorVersionResponse {
	s.Headers = v
	return s
}

func (s *QueryMinorVersionResponse) SetStatusCode(v int32) *QueryMinorVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinorVersionResponse) SetBody(v *QueryMinorVersionResponseBody) *QueryMinorVersionResponse {
	s.Body = v
	return s
}

func (s *QueryMinorVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
