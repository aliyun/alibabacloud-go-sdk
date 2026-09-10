// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostInnerReaderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostInnerReaderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostInnerReaderResponse
	GetStatusCode() *int32
	SetBody(v *PostInnerReaderResponseBody) *PostInnerReaderResponse
	GetBody() *PostInnerReaderResponseBody
}

type PostInnerReaderResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostInnerReaderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostInnerReaderResponse) String() string {
	return dara.Prettify(s)
}

func (s PostInnerReaderResponse) GoString() string {
	return s.String()
}

func (s *PostInnerReaderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostInnerReaderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostInnerReaderResponse) GetBody() *PostInnerReaderResponseBody {
	return s.Body
}

func (s *PostInnerReaderResponse) SetHeaders(v map[string]*string) *PostInnerReaderResponse {
	s.Headers = v
	return s
}

func (s *PostInnerReaderResponse) SetStatusCode(v int32) *PostInnerReaderResponse {
	s.StatusCode = &v
	return s
}

func (s *PostInnerReaderResponse) SetBody(v *PostInnerReaderResponseBody) *PostInnerReaderResponse {
	s.Body = v
	return s
}

func (s *PostInnerReaderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
