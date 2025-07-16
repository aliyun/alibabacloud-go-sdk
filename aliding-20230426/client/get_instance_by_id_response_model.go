// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceByIdResponseBody) *GetInstanceByIdResponse
	GetBody() *GetInstanceByIdResponseBody
}

type GetInstanceByIdResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceByIdResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceByIdResponse) GetBody() *GetInstanceByIdResponseBody {
	return s.Body
}

func (s *GetInstanceByIdResponse) SetHeaders(v map[string]*string) *GetInstanceByIdResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceByIdResponse) SetStatusCode(v int32) *GetInstanceByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceByIdResponse) SetBody(v *GetInstanceByIdResponseBody) *GetInstanceByIdResponse {
	s.Body = v
	return s
}

func (s *GetInstanceByIdResponse) Validate() error {
	return dara.Validate(s)
}
