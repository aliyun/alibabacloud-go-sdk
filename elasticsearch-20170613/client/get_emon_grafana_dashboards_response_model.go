// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmonGrafanaDashboardsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEmonGrafanaDashboardsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEmonGrafanaDashboardsResponse
	GetStatusCode() *int32
	SetBody(v *GetEmonGrafanaDashboardsResponseBody) *GetEmonGrafanaDashboardsResponse
	GetBody() *GetEmonGrafanaDashboardsResponseBody
}

type GetEmonGrafanaDashboardsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEmonGrafanaDashboardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmonGrafanaDashboardsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEmonGrafanaDashboardsResponse) GoString() string {
	return s.String()
}

func (s *GetEmonGrafanaDashboardsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEmonGrafanaDashboardsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEmonGrafanaDashboardsResponse) GetBody() *GetEmonGrafanaDashboardsResponseBody {
	return s.Body
}

func (s *GetEmonGrafanaDashboardsResponse) SetHeaders(v map[string]*string) *GetEmonGrafanaDashboardsResponse {
	s.Headers = v
	return s
}

func (s *GetEmonGrafanaDashboardsResponse) SetStatusCode(v int32) *GetEmonGrafanaDashboardsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmonGrafanaDashboardsResponse) SetBody(v *GetEmonGrafanaDashboardsResponseBody) *GetEmonGrafanaDashboardsResponse {
	s.Body = v
	return s
}

func (s *GetEmonGrafanaDashboardsResponse) Validate() error {
	return dara.Validate(s)
}
