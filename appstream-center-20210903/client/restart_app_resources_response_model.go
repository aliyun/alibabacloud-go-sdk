// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartAppResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartAppResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartAppResourcesResponse
	GetStatusCode() *int32
	SetBody(v *RestartAppResourcesResponseBody) *RestartAppResourcesResponse
	GetBody() *RestartAppResourcesResponseBody
}

type RestartAppResourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartAppResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartAppResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartAppResourcesResponse) GoString() string {
	return s.String()
}

func (s *RestartAppResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartAppResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartAppResourcesResponse) GetBody() *RestartAppResourcesResponseBody {
	return s.Body
}

func (s *RestartAppResourcesResponse) SetHeaders(v map[string]*string) *RestartAppResourcesResponse {
	s.Headers = v
	return s
}

func (s *RestartAppResourcesResponse) SetStatusCode(v int32) *RestartAppResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartAppResourcesResponse) SetBody(v *RestartAppResourcesResponseBody) *RestartAppResourcesResponse {
	s.Body = v
	return s
}

func (s *RestartAppResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
