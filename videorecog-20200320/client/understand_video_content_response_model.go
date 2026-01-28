// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnderstandVideoContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnderstandVideoContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnderstandVideoContentResponse
	GetStatusCode() *int32
	SetBody(v *UnderstandVideoContentResponseBody) *UnderstandVideoContentResponse
	GetBody() *UnderstandVideoContentResponseBody
}

type UnderstandVideoContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnderstandVideoContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnderstandVideoContentResponse) String() string {
	return dara.Prettify(s)
}

func (s UnderstandVideoContentResponse) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnderstandVideoContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnderstandVideoContentResponse) GetBody() *UnderstandVideoContentResponseBody {
	return s.Body
}

func (s *UnderstandVideoContentResponse) SetHeaders(v map[string]*string) *UnderstandVideoContentResponse {
	s.Headers = v
	return s
}

func (s *UnderstandVideoContentResponse) SetStatusCode(v int32) *UnderstandVideoContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UnderstandVideoContentResponse) SetBody(v *UnderstandVideoContentResponseBody) *UnderstandVideoContentResponse {
	s.Body = v
	return s
}

func (s *UnderstandVideoContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
