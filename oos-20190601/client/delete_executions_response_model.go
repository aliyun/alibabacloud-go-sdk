// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExecutionsResponseBody) *DeleteExecutionsResponse
	GetBody() *DeleteExecutionsResponseBody
}

type DeleteExecutionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExecutionsResponse) GoString() string {
	return s.String()
}

func (s *DeleteExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExecutionsResponse) GetBody() *DeleteExecutionsResponseBody {
	return s.Body
}

func (s *DeleteExecutionsResponse) SetHeaders(v map[string]*string) *DeleteExecutionsResponse {
	s.Headers = v
	return s
}

func (s *DeleteExecutionsResponse) SetStatusCode(v int32) *DeleteExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExecutionsResponse) SetBody(v *DeleteExecutionsResponseBody) *DeleteExecutionsResponse {
	s.Body = v
	return s
}

func (s *DeleteExecutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
