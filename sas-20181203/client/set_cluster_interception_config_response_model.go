// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetClusterInterceptionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetClusterInterceptionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetClusterInterceptionConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetClusterInterceptionConfigResponseBody) *SetClusterInterceptionConfigResponse
	GetBody() *SetClusterInterceptionConfigResponseBody
}

type SetClusterInterceptionConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetClusterInterceptionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetClusterInterceptionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetClusterInterceptionConfigResponse) GoString() string {
	return s.String()
}

func (s *SetClusterInterceptionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetClusterInterceptionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetClusterInterceptionConfigResponse) GetBody() *SetClusterInterceptionConfigResponseBody {
	return s.Body
}

func (s *SetClusterInterceptionConfigResponse) SetHeaders(v map[string]*string) *SetClusterInterceptionConfigResponse {
	s.Headers = v
	return s
}

func (s *SetClusterInterceptionConfigResponse) SetStatusCode(v int32) *SetClusterInterceptionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetClusterInterceptionConfigResponse) SetBody(v *SetClusterInterceptionConfigResponseBody) *SetClusterInterceptionConfigResponse {
	s.Body = v
	return s
}

func (s *SetClusterInterceptionConfigResponse) Validate() error {
	return dara.Validate(s)
}
