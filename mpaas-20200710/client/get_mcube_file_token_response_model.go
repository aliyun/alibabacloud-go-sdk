// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeFileTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcubeFileTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcubeFileTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetMcubeFileTokenResponseBody) *GetMcubeFileTokenResponse
	GetBody() *GetMcubeFileTokenResponseBody
}

type GetMcubeFileTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcubeFileTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcubeFileTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeFileTokenResponse) GoString() string {
	return s.String()
}

func (s *GetMcubeFileTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcubeFileTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcubeFileTokenResponse) GetBody() *GetMcubeFileTokenResponseBody {
	return s.Body
}

func (s *GetMcubeFileTokenResponse) SetHeaders(v map[string]*string) *GetMcubeFileTokenResponse {
	s.Headers = v
	return s
}

func (s *GetMcubeFileTokenResponse) SetStatusCode(v int32) *GetMcubeFileTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcubeFileTokenResponse) SetBody(v *GetMcubeFileTokenResponseBody) *GetMcubeFileTokenResponse {
	s.Body = v
	return s
}

func (s *GetMcubeFileTokenResponse) Validate() error {
	return dara.Validate(s)
}
