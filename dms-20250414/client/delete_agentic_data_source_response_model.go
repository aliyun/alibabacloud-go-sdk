// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgenticDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgenticDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgenticDataSourceResponseBody) *DeleteAgenticDataSourceResponse
	GetBody() *DeleteAgenticDataSourceResponseBody
}

type DeleteAgenticDataSourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgenticDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgenticDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgenticDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgenticDataSourceResponse) GetBody() *DeleteAgenticDataSourceResponseBody {
	return s.Body
}

func (s *DeleteAgenticDataSourceResponse) SetHeaders(v map[string]*string) *DeleteAgenticDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgenticDataSourceResponse) SetStatusCode(v int32) *DeleteAgenticDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgenticDataSourceResponse) SetBody(v *DeleteAgenticDataSourceResponseBody) *DeleteAgenticDataSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteAgenticDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
