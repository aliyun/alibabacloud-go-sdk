// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogicTableRouteConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLogicTableRouteConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLogicTableRouteConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLogicTableRouteConfigResponseBody) *DeleteLogicTableRouteConfigResponse
	GetBody() *DeleteLogicTableRouteConfigResponseBody
}

type DeleteLogicTableRouteConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLogicTableRouteConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLogicTableRouteConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogicTableRouteConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogicTableRouteConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLogicTableRouteConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLogicTableRouteConfigResponse) GetBody() *DeleteLogicTableRouteConfigResponseBody {
	return s.Body
}

func (s *DeleteLogicTableRouteConfigResponse) SetHeaders(v map[string]*string) *DeleteLogicTableRouteConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogicTableRouteConfigResponse) SetStatusCode(v int32) *DeleteLogicTableRouteConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLogicTableRouteConfigResponse) SetBody(v *DeleteLogicTableRouteConfigResponseBody) *DeleteLogicTableRouteConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLogicTableRouteConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
