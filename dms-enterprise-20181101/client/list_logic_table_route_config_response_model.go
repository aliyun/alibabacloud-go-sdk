// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogicTableRouteConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogicTableRouteConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogicTableRouteConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListLogicTableRouteConfigResponseBody) *ListLogicTableRouteConfigResponse
	GetBody() *ListLogicTableRouteConfigResponseBody
}

type ListLogicTableRouteConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogicTableRouteConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogicTableRouteConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogicTableRouteConfigResponse) GoString() string {
	return s.String()
}

func (s *ListLogicTableRouteConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogicTableRouteConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogicTableRouteConfigResponse) GetBody() *ListLogicTableRouteConfigResponseBody {
	return s.Body
}

func (s *ListLogicTableRouteConfigResponse) SetHeaders(v map[string]*string) *ListLogicTableRouteConfigResponse {
	s.Headers = v
	return s
}

func (s *ListLogicTableRouteConfigResponse) SetStatusCode(v int32) *ListLogicTableRouteConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogicTableRouteConfigResponse) SetBody(v *ListLogicTableRouteConfigResponseBody) *ListLogicTableRouteConfigResponse {
	s.Body = v
	return s
}

func (s *ListLogicTableRouteConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
