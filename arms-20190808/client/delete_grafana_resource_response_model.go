// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGrafanaResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGrafanaResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGrafanaResourceResponseBody) *DeleteGrafanaResourceResponse
	GetBody() *DeleteGrafanaResourceResponseBody
}

type DeleteGrafanaResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGrafanaResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGrafanaResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGrafanaResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGrafanaResourceResponse) GetBody() *DeleteGrafanaResourceResponseBody {
	return s.Body
}

func (s *DeleteGrafanaResourceResponse) SetHeaders(v map[string]*string) *DeleteGrafanaResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGrafanaResourceResponse) SetStatusCode(v int32) *DeleteGrafanaResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGrafanaResourceResponse) SetBody(v *DeleteGrafanaResourceResponseBody) *DeleteGrafanaResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteGrafanaResourceResponse) Validate() error {
	return dara.Validate(s)
}
