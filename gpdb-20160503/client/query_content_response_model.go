// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryContentResponse
	GetStatusCode() *int32
	SetBody(v *QueryContentResponseBody) *QueryContentResponse
	GetBody() *QueryContentResponseBody
}

type QueryContentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryContentResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponse) GoString() string {
	return s.String()
}

func (s *QueryContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryContentResponse) GetBody() *QueryContentResponseBody {
	return s.Body
}

func (s *QueryContentResponse) SetHeaders(v map[string]*string) *QueryContentResponse {
	s.Headers = v
	return s
}

func (s *QueryContentResponse) SetStatusCode(v int32) *QueryContentResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryContentResponse) SetBody(v *QueryContentResponseBody) *QueryContentResponse {
	s.Body = v
	return s
}

func (s *QueryContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
