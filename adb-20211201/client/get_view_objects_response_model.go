// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetViewObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetViewObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetViewObjectsResponse
	GetStatusCode() *int32
	SetBody(v *GetViewObjectsResponseBody) *GetViewObjectsResponse
	GetBody() *GetViewObjectsResponseBody
}

type GetViewObjectsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetViewObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetViewObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetViewObjectsResponse) GoString() string {
	return s.String()
}

func (s *GetViewObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetViewObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetViewObjectsResponse) GetBody() *GetViewObjectsResponseBody {
	return s.Body
}

func (s *GetViewObjectsResponse) SetHeaders(v map[string]*string) *GetViewObjectsResponse {
	s.Headers = v
	return s
}

func (s *GetViewObjectsResponse) SetStatusCode(v int32) *GetViewObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetViewObjectsResponse) SetBody(v *GetViewObjectsResponseBody) *GetViewObjectsResponse {
	s.Body = v
	return s
}

func (s *GetViewObjectsResponse) Validate() error {
	return dara.Validate(s)
}
