// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPoolResponse
	GetStatusCode() *int32
	SetBody(v *GetPoolResponseBody) *GetPoolResponse
	GetBody() *GetPoolResponseBody
}

type GetPoolResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPoolResponse) GoString() string {
	return s.String()
}

func (s *GetPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPoolResponse) GetBody() *GetPoolResponseBody {
	return s.Body
}

func (s *GetPoolResponse) SetHeaders(v map[string]*string) *GetPoolResponse {
	s.Headers = v
	return s
}

func (s *GetPoolResponse) SetStatusCode(v int32) *GetPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPoolResponse) SetBody(v *GetPoolResponseBody) *GetPoolResponse {
	s.Body = v
	return s
}

func (s *GetPoolResponse) Validate() error {
	return dara.Validate(s)
}
