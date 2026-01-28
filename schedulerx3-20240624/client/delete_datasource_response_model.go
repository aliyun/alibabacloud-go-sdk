// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDatasourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDatasourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDatasourceResponseBody) *DeleteDatasourceResponse
	GetBody() *DeleteDatasourceResponseBody
}

type DeleteDatasourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDatasourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDatasourceResponse) GetBody() *DeleteDatasourceResponseBody {
	return s.Body
}

func (s *DeleteDatasourceResponse) SetHeaders(v map[string]*string) *DeleteDatasourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasourceResponse) SetStatusCode(v int32) *DeleteDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasourceResponse) SetBody(v *DeleteDatasourceResponseBody) *DeleteDatasourceResponse {
	s.Body = v
	return s
}

func (s *DeleteDatasourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
