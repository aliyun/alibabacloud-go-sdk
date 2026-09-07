// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceHttpApiProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrafanaWorkspaceHttpApiProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrafanaWorkspaceHttpApiProxyResponse
	GetStatusCode() *int32
	SetBody(v *GrafanaWorkspaceHttpApiProxyResponseBody) *GrafanaWorkspaceHttpApiProxyResponse
	GetBody() *GrafanaWorkspaceHttpApiProxyResponseBody
}

type GrafanaWorkspaceHttpApiProxyResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrafanaWorkspaceHttpApiProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrafanaWorkspaceHttpApiProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceHttpApiProxyResponse) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceHttpApiProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrafanaWorkspaceHttpApiProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrafanaWorkspaceHttpApiProxyResponse) GetBody() *GrafanaWorkspaceHttpApiProxyResponseBody {
	return s.Body
}

func (s *GrafanaWorkspaceHttpApiProxyResponse) SetHeaders(v map[string]*string) *GrafanaWorkspaceHttpApiProxyResponse {
	s.Headers = v
	return s
}

func (s *GrafanaWorkspaceHttpApiProxyResponse) SetStatusCode(v int32) *GrafanaWorkspaceHttpApiProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *GrafanaWorkspaceHttpApiProxyResponse) SetBody(v *GrafanaWorkspaceHttpApiProxyResponseBody) *GrafanaWorkspaceHttpApiProxyResponse {
	s.Body = v
	return s
}

func (s *GrafanaWorkspaceHttpApiProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
