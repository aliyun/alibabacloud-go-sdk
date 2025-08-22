// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppResponse
	GetStatusCode() *int32
	SetBody(v *GetAppResponseBody) *GetAppResponse
	GetBody() *GetAppResponseBody
}

type GetAppResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppResponse) GoString() string {
	return s.String()
}

func (s *GetAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppResponse) GetBody() *GetAppResponseBody {
	return s.Body
}

func (s *GetAppResponse) SetHeaders(v map[string]*string) *GetAppResponse {
	s.Headers = v
	return s
}

func (s *GetAppResponse) SetStatusCode(v int32) *GetAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppResponse) SetBody(v *GetAppResponseBody) *GetAppResponse {
	s.Body = v
	return s
}

func (s *GetAppResponse) Validate() error {
	return dara.Validate(s)
}
