// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTagResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTagResponseBody) *DeleteTagResponse
	GetBody() *DeleteTagResponseBody
}

type DeleteTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTagResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTagResponse) GetBody() *DeleteTagResponseBody {
	return s.Body
}

func (s *DeleteTagResponse) SetHeaders(v map[string]*string) *DeleteTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagResponse) SetStatusCode(v int32) *DeleteTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTagResponse) SetBody(v *DeleteTagResponseBody) *DeleteTagResponse {
	s.Body = v
	return s
}

func (s *DeleteTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
