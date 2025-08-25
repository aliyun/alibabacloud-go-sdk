// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeMiniAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcubeMiniAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcubeMiniAppResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcubeMiniAppResponseBody) *DeleteMcubeMiniAppResponse
	GetBody() *DeleteMcubeMiniAppResponseBody
}

type DeleteMcubeMiniAppResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcubeMiniAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcubeMiniAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeMiniAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcubeMiniAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcubeMiniAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcubeMiniAppResponse) GetBody() *DeleteMcubeMiniAppResponseBody {
	return s.Body
}

func (s *DeleteMcubeMiniAppResponse) SetHeaders(v map[string]*string) *DeleteMcubeMiniAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcubeMiniAppResponse) SetStatusCode(v int32) *DeleteMcubeMiniAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcubeMiniAppResponse) SetBody(v *DeleteMcubeMiniAppResponseBody) *DeleteMcubeMiniAppResponse {
	s.Body = v
	return s
}

func (s *DeleteMcubeMiniAppResponse) Validate() error {
	return dara.Validate(s)
}
