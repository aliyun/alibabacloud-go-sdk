// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHostResponse
	GetStatusCode() *int32
	SetBody(v *GetHostResponseBody) *GetHostResponse
	GetBody() *GetHostResponseBody
}

type GetHostResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHostResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHostResponse) GoString() string {
	return s.String()
}

func (s *GetHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHostResponse) GetBody() *GetHostResponseBody {
	return s.Body
}

func (s *GetHostResponse) SetHeaders(v map[string]*string) *GetHostResponse {
	s.Headers = v
	return s
}

func (s *GetHostResponse) SetStatusCode(v int32) *GetHostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHostResponse) SetBody(v *GetHostResponseBody) *GetHostResponse {
	s.Body = v
	return s
}

func (s *GetHostResponse) Validate() error {
	return dara.Validate(s)
}
