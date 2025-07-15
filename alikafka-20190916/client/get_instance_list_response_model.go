// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceListResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceListResponseBody) *GetInstanceListResponse
	GetBody() *GetInstanceListResponseBody
}

type GetInstanceListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceListResponse) GetBody() *GetInstanceListResponseBody {
	return s.Body
}

func (s *GetInstanceListResponse) SetHeaders(v map[string]*string) *GetInstanceListResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceListResponse) SetStatusCode(v int32) *GetInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceListResponse) SetBody(v *GetInstanceListResponseBody) *GetInstanceListResponse {
	s.Body = v
	return s
}

func (s *GetInstanceListResponse) Validate() error {
	return dara.Validate(s)
}
