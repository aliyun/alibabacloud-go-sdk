// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeMiniAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcubeMiniAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcubeMiniAppResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcubeMiniAppResponseBody) *CreateMcubeMiniAppResponse
	GetBody() *CreateMcubeMiniAppResponseBody
}

type CreateMcubeMiniAppResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcubeMiniAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcubeMiniAppResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeMiniAppResponse) GoString() string {
	return s.String()
}

func (s *CreateMcubeMiniAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcubeMiniAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcubeMiniAppResponse) GetBody() *CreateMcubeMiniAppResponseBody {
	return s.Body
}

func (s *CreateMcubeMiniAppResponse) SetHeaders(v map[string]*string) *CreateMcubeMiniAppResponse {
	s.Headers = v
	return s
}

func (s *CreateMcubeMiniAppResponse) SetStatusCode(v int32) *CreateMcubeMiniAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcubeMiniAppResponse) SetBody(v *CreateMcubeMiniAppResponseBody) *CreateMcubeMiniAppResponse {
	s.Body = v
	return s
}

func (s *CreateMcubeMiniAppResponse) Validate() error {
	return dara.Validate(s)
}
