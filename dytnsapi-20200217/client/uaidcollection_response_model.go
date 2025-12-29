// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUAIDCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UAIDCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UAIDCollectionResponse
	GetStatusCode() *int32
	SetBody(v *UAIDCollectionResponseBody) *UAIDCollectionResponse
	GetBody() *UAIDCollectionResponseBody
}

type UAIDCollectionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UAIDCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UAIDCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UAIDCollectionResponse) GoString() string {
	return s.String()
}

func (s *UAIDCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UAIDCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UAIDCollectionResponse) GetBody() *UAIDCollectionResponseBody {
	return s.Body
}

func (s *UAIDCollectionResponse) SetHeaders(v map[string]*string) *UAIDCollectionResponse {
	s.Headers = v
	return s
}

func (s *UAIDCollectionResponse) SetStatusCode(v int32) *UAIDCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UAIDCollectionResponse) SetBody(v *UAIDCollectionResponseBody) *UAIDCollectionResponse {
	s.Body = v
	return s
}

func (s *UAIDCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
