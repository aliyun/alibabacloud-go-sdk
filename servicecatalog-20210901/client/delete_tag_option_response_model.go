// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagOptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTagOptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTagOptionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTagOptionResponseBody) *DeleteTagOptionResponse
	GetBody() *DeleteTagOptionResponseBody
}

type DeleteTagOptionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTagOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTagOptionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagOptionResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagOptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTagOptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTagOptionResponse) GetBody() *DeleteTagOptionResponseBody {
	return s.Body
}

func (s *DeleteTagOptionResponse) SetHeaders(v map[string]*string) *DeleteTagOptionResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagOptionResponse) SetStatusCode(v int32) *DeleteTagOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTagOptionResponse) SetBody(v *DeleteTagOptionResponseBody) *DeleteTagOptionResponse {
	s.Body = v
	return s
}

func (s *DeleteTagOptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
