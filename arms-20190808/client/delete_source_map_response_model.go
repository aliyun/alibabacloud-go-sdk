// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSourceMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSourceMapResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSourceMapResponseBody) *DeleteSourceMapResponse
	GetBody() *DeleteSourceMapResponseBody
}

type DeleteSourceMapResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSourceMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSourceMapResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceMapResponse) GoString() string {
	return s.String()
}

func (s *DeleteSourceMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSourceMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSourceMapResponse) GetBody() *DeleteSourceMapResponseBody {
	return s.Body
}

func (s *DeleteSourceMapResponse) SetHeaders(v map[string]*string) *DeleteSourceMapResponse {
	s.Headers = v
	return s
}

func (s *DeleteSourceMapResponse) SetStatusCode(v int32) *DeleteSourceMapResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSourceMapResponse) SetBody(v *DeleteSourceMapResponseBody) *DeleteSourceMapResponse {
	s.Body = v
	return s
}

func (s *DeleteSourceMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
