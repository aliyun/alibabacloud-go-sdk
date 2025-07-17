// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLogicTableRouteConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLogicTableRouteConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLogicTableRouteConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddLogicTableRouteConfigResponseBody) *AddLogicTableRouteConfigResponse
	GetBody() *AddLogicTableRouteConfigResponseBody
}

type AddLogicTableRouteConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLogicTableRouteConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLogicTableRouteConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLogicTableRouteConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLogicTableRouteConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLogicTableRouteConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLogicTableRouteConfigResponse) GetBody() *AddLogicTableRouteConfigResponseBody {
	return s.Body
}

func (s *AddLogicTableRouteConfigResponse) SetHeaders(v map[string]*string) *AddLogicTableRouteConfigResponse {
	s.Headers = v
	return s
}

func (s *AddLogicTableRouteConfigResponse) SetStatusCode(v int32) *AddLogicTableRouteConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLogicTableRouteConfigResponse) SetBody(v *AddLogicTableRouteConfigResponseBody) *AddLogicTableRouteConfigResponse {
	s.Body = v
	return s
}

func (s *AddLogicTableRouteConfigResponse) Validate() error {
	return dara.Validate(s)
}
