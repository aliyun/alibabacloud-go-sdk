// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWebContainerConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWebContainerConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWebContainerConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetWebContainerConfigResponseBody) *GetWebContainerConfigResponse
	GetBody() *GetWebContainerConfigResponseBody
}

type GetWebContainerConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWebContainerConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWebContainerConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWebContainerConfigResponse) GoString() string {
	return s.String()
}

func (s *GetWebContainerConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWebContainerConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWebContainerConfigResponse) GetBody() *GetWebContainerConfigResponseBody {
	return s.Body
}

func (s *GetWebContainerConfigResponse) SetHeaders(v map[string]*string) *GetWebContainerConfigResponse {
	s.Headers = v
	return s
}

func (s *GetWebContainerConfigResponse) SetStatusCode(v int32) *GetWebContainerConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWebContainerConfigResponse) SetBody(v *GetWebContainerConfigResponseBody) *GetWebContainerConfigResponse {
	s.Body = v
	return s
}

func (s *GetWebContainerConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
