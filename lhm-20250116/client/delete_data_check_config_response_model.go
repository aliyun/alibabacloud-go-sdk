// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataCheckConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataCheckConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataCheckConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataCheckConfigResponseBody) *DeleteDataCheckConfigResponse
	GetBody() *DeleteDataCheckConfigResponseBody
}

type DeleteDataCheckConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataCheckConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataCheckConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataCheckConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataCheckConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataCheckConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataCheckConfigResponse) GetBody() *DeleteDataCheckConfigResponseBody {
	return s.Body
}

func (s *DeleteDataCheckConfigResponse) SetHeaders(v map[string]*string) *DeleteDataCheckConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataCheckConfigResponse) SetStatusCode(v int32) *DeleteDataCheckConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataCheckConfigResponse) SetBody(v *DeleteDataCheckConfigResponseBody) *DeleteDataCheckConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteDataCheckConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
