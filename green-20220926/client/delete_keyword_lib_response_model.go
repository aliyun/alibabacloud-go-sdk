// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeywordLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKeywordLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKeywordLibResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKeywordLibResponseBody) *DeleteKeywordLibResponse
	GetBody() *DeleteKeywordLibResponseBody
}

type DeleteKeywordLibResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKeywordLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKeywordLibResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeywordLibResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeywordLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKeywordLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKeywordLibResponse) GetBody() *DeleteKeywordLibResponseBody {
	return s.Body
}

func (s *DeleteKeywordLibResponse) SetHeaders(v map[string]*string) *DeleteKeywordLibResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeywordLibResponse) SetStatusCode(v int32) *DeleteKeywordLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeywordLibResponse) SetBody(v *DeleteKeywordLibResponseBody) *DeleteKeywordLibResponse {
	s.Body = v
	return s
}

func (s *DeleteKeywordLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
