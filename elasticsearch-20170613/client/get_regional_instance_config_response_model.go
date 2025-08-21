// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegionalInstanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRegionalInstanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRegionalInstanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetRegionalInstanceConfigResponseBody) *GetRegionalInstanceConfigResponse
	GetBody() *GetRegionalInstanceConfigResponseBody
}

type GetRegionalInstanceConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegionalInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegionalInstanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRegionalInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *GetRegionalInstanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRegionalInstanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRegionalInstanceConfigResponse) GetBody() *GetRegionalInstanceConfigResponseBody {
	return s.Body
}

func (s *GetRegionalInstanceConfigResponse) SetHeaders(v map[string]*string) *GetRegionalInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *GetRegionalInstanceConfigResponse) SetStatusCode(v int32) *GetRegionalInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegionalInstanceConfigResponse) SetBody(v *GetRegionalInstanceConfigResponseBody) *GetRegionalInstanceConfigResponse {
	s.Body = v
	return s
}

func (s *GetRegionalInstanceConfigResponse) Validate() error {
	return dara.Validate(s)
}
