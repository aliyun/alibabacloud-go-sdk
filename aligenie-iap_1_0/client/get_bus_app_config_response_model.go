// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusAppConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBusAppConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBusAppConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetBusAppConfigResponseBody) *GetBusAppConfigResponse
	GetBody() *GetBusAppConfigResponseBody
}

type GetBusAppConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBusAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBusAppConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBusAppConfigResponse) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBusAppConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBusAppConfigResponse) GetBody() *GetBusAppConfigResponseBody {
	return s.Body
}

func (s *GetBusAppConfigResponse) SetHeaders(v map[string]*string) *GetBusAppConfigResponse {
	s.Headers = v
	return s
}

func (s *GetBusAppConfigResponse) SetStatusCode(v int32) *GetBusAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBusAppConfigResponse) SetBody(v *GetBusAppConfigResponseBody) *GetBusAppConfigResponse {
	s.Body = v
	return s
}

func (s *GetBusAppConfigResponse) Validate() error {
	return dara.Validate(s)
}
