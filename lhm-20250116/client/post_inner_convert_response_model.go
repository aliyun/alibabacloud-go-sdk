// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostInnerConvertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostInnerConvertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostInnerConvertResponse
	GetStatusCode() *int32
	SetBody(v *PostInnerConvertResponseBody) *PostInnerConvertResponse
	GetBody() *PostInnerConvertResponseBody
}

type PostInnerConvertResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostInnerConvertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostInnerConvertResponse) String() string {
	return dara.Prettify(s)
}

func (s PostInnerConvertResponse) GoString() string {
	return s.String()
}

func (s *PostInnerConvertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostInnerConvertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostInnerConvertResponse) GetBody() *PostInnerConvertResponseBody {
	return s.Body
}

func (s *PostInnerConvertResponse) SetHeaders(v map[string]*string) *PostInnerConvertResponse {
	s.Headers = v
	return s
}

func (s *PostInnerConvertResponse) SetStatusCode(v int32) *PostInnerConvertResponse {
	s.StatusCode = &v
	return s
}

func (s *PostInnerConvertResponse) SetBody(v *PostInnerConvertResponseBody) *PostInnerConvertResponse {
	s.Body = v
	return s
}

func (s *PostInnerConvertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
