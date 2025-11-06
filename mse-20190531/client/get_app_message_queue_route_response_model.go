// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppMessageQueueRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppMessageQueueRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppMessageQueueRouteResponse
	GetStatusCode() *int32
	SetBody(v *GetAppMessageQueueRouteResponseBody) *GetAppMessageQueueRouteResponse
	GetBody() *GetAppMessageQueueRouteResponseBody
}

type GetAppMessageQueueRouteResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppMessageQueueRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppMessageQueueRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppMessageQueueRouteResponse) GoString() string {
	return s.String()
}

func (s *GetAppMessageQueueRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppMessageQueueRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppMessageQueueRouteResponse) GetBody() *GetAppMessageQueueRouteResponseBody {
	return s.Body
}

func (s *GetAppMessageQueueRouteResponse) SetHeaders(v map[string]*string) *GetAppMessageQueueRouteResponse {
	s.Headers = v
	return s
}

func (s *GetAppMessageQueueRouteResponse) SetStatusCode(v int32) *GetAppMessageQueueRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppMessageQueueRouteResponse) SetBody(v *GetAppMessageQueueRouteResponseBody) *GetAppMessageQueueRouteResponse {
	s.Body = v
	return s
}

func (s *GetAppMessageQueueRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
