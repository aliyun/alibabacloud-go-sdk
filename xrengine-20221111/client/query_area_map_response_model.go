// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAreaMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAreaMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAreaMapResponse
	GetStatusCode() *int32
	SetBody(v *QueryAreaMapResponseBody) *QueryAreaMapResponse
	GetBody() *QueryAreaMapResponseBody
}

type QueryAreaMapResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAreaMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAreaMapResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAreaMapResponse) GoString() string {
	return s.String()
}

func (s *QueryAreaMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAreaMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAreaMapResponse) GetBody() *QueryAreaMapResponseBody {
	return s.Body
}

func (s *QueryAreaMapResponse) SetHeaders(v map[string]*string) *QueryAreaMapResponse {
	s.Headers = v
	return s
}

func (s *QueryAreaMapResponse) SetStatusCode(v int32) *QueryAreaMapResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAreaMapResponse) SetBody(v *QueryAreaMapResponseBody) *QueryAreaMapResponse {
	s.Body = v
	return s
}

func (s *QueryAreaMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
