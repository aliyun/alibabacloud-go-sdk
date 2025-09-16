// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeployGraphResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeployGraphResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeployGraphResponse
	GetStatusCode() *int32
	SetBody(v *GetDeployGraphResponseBody) *GetDeployGraphResponse
	GetBody() *GetDeployGraphResponseBody
}

type GetDeployGraphResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeployGraphResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeployGraphResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeployGraphResponse) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeployGraphResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeployGraphResponse) GetBody() *GetDeployGraphResponseBody {
	return s.Body
}

func (s *GetDeployGraphResponse) SetHeaders(v map[string]*string) *GetDeployGraphResponse {
	s.Headers = v
	return s
}

func (s *GetDeployGraphResponse) SetStatusCode(v int32) *GetDeployGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeployGraphResponse) SetBody(v *GetDeployGraphResponseBody) *GetDeployGraphResponse {
	s.Body = v
	return s
}

func (s *GetDeployGraphResponse) Validate() error {
	return dara.Validate(s)
}
