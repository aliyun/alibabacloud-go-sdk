// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIndexResponse
	GetStatusCode() *int32
	SetBody(v *GetIndexResponseBody) *GetIndexResponse
	GetBody() *GetIndexResponseBody
}

type GetIndexResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIndexResponse) GoString() string {
	return s.String()
}

func (s *GetIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIndexResponse) GetBody() *GetIndexResponseBody {
	return s.Body
}

func (s *GetIndexResponse) SetHeaders(v map[string]*string) *GetIndexResponse {
	s.Headers = v
	return s
}

func (s *GetIndexResponse) SetStatusCode(v int32) *GetIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexResponse) SetBody(v *GetIndexResponseBody) *GetIndexResponse {
	s.Body = v
	return s
}

func (s *GetIndexResponse) Validate() error {
	return dara.Validate(s)
}
