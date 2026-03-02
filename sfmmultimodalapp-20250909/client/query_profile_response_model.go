// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryProfileResponse
	GetStatusCode() *int32
	SetBody(v *QueryProfileResponseBody) *QueryProfileResponse
	GetBody() *QueryProfileResponseBody
}

type QueryProfileResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryProfileResponse) GoString() string {
	return s.String()
}

func (s *QueryProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryProfileResponse) GetBody() *QueryProfileResponseBody {
	return s.Body
}

func (s *QueryProfileResponse) SetHeaders(v map[string]*string) *QueryProfileResponse {
	s.Headers = v
	return s
}

func (s *QueryProfileResponse) SetStatusCode(v int32) *QueryProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProfileResponse) SetBody(v *QueryProfileResponseBody) *QueryProfileResponse {
	s.Body = v
	return s
}

func (s *QueryProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
