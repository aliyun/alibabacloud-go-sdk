// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeConfigResponseBody) *GetNodeConfigResponse
	GetBody() *GetNodeConfigResponseBody
}

type GetNodeConfigResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeConfigResponse) GoString() string {
	return s.String()
}

func (s *GetNodeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeConfigResponse) GetBody() *GetNodeConfigResponseBody {
	return s.Body
}

func (s *GetNodeConfigResponse) SetHeaders(v map[string]*string) *GetNodeConfigResponse {
	s.Headers = v
	return s
}

func (s *GetNodeConfigResponse) SetStatusCode(v int32) *GetNodeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeConfigResponse) SetBody(v *GetNodeConfigResponseBody) *GetNodeConfigResponse {
	s.Body = v
	return s
}

func (s *GetNodeConfigResponse) Validate() error {
	return dara.Validate(s)
}
