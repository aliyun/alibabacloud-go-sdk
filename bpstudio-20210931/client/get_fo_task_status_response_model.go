// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFoTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFoTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFoTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetFoTaskStatusResponseBody) *GetFoTaskStatusResponse
	GetBody() *GetFoTaskStatusResponseBody
}

type GetFoTaskStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFoTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFoTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFoTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetFoTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFoTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFoTaskStatusResponse) GetBody() *GetFoTaskStatusResponseBody {
	return s.Body
}

func (s *GetFoTaskStatusResponse) SetHeaders(v map[string]*string) *GetFoTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetFoTaskStatusResponse) SetStatusCode(v int32) *GetFoTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFoTaskStatusResponse) SetBody(v *GetFoTaskStatusResponseBody) *GetFoTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetFoTaskStatusResponse) Validate() error {
	return dara.Validate(s)
}
