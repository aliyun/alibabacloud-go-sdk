// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLogPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLogPathResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLogPathResponseBody) *DeleteLogPathResponse
	GetBody() *DeleteLogPathResponseBody
}

type DeleteLogPathResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLogPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLogPathResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogPathResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLogPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLogPathResponse) GetBody() *DeleteLogPathResponseBody {
	return s.Body
}

func (s *DeleteLogPathResponse) SetHeaders(v map[string]*string) *DeleteLogPathResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogPathResponse) SetStatusCode(v int32) *DeleteLogPathResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLogPathResponse) SetBody(v *DeleteLogPathResponseBody) *DeleteLogPathResponse {
	s.Body = v
	return s
}

func (s *DeleteLogPathResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
