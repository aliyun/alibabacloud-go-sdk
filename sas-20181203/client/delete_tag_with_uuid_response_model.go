// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagWithUuidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTagWithUuidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTagWithUuidResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTagWithUuidResponseBody) *DeleteTagWithUuidResponse
	GetBody() *DeleteTagWithUuidResponseBody
}

type DeleteTagWithUuidResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTagWithUuidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTagWithUuidResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagWithUuidResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagWithUuidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTagWithUuidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTagWithUuidResponse) GetBody() *DeleteTagWithUuidResponseBody {
	return s.Body
}

func (s *DeleteTagWithUuidResponse) SetHeaders(v map[string]*string) *DeleteTagWithUuidResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagWithUuidResponse) SetStatusCode(v int32) *DeleteTagWithUuidResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTagWithUuidResponse) SetBody(v *DeleteTagWithUuidResponseBody) *DeleteTagWithUuidResponse {
	s.Body = v
	return s
}

func (s *DeleteTagWithUuidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
