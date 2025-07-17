// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemoryNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemoryNodeResponse
	GetStatusCode() *int32
	SetBody(v *GetMemoryNodeResponseBody) *GetMemoryNodeResponse
	GetBody() *GetMemoryNodeResponseBody
}

type GetMemoryNodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemoryNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoryNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryNodeResponse) GoString() string {
	return s.String()
}

func (s *GetMemoryNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemoryNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemoryNodeResponse) GetBody() *GetMemoryNodeResponseBody {
	return s.Body
}

func (s *GetMemoryNodeResponse) SetHeaders(v map[string]*string) *GetMemoryNodeResponse {
	s.Headers = v
	return s
}

func (s *GetMemoryNodeResponse) SetStatusCode(v int32) *GetMemoryNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoryNodeResponse) SetBody(v *GetMemoryNodeResponseBody) *GetMemoryNodeResponse {
	s.Body = v
	return s
}

func (s *GetMemoryNodeResponse) Validate() error {
	return dara.Validate(s)
}
