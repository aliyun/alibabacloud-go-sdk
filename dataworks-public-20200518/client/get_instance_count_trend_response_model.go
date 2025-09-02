// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceCountTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceCountTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceCountTrendResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceCountTrendResponseBody) *GetInstanceCountTrendResponse
	GetBody() *GetInstanceCountTrendResponseBody
}

type GetInstanceCountTrendResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceCountTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceCountTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceCountTrendResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceCountTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceCountTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceCountTrendResponse) GetBody() *GetInstanceCountTrendResponseBody {
	return s.Body
}

func (s *GetInstanceCountTrendResponse) SetHeaders(v map[string]*string) *GetInstanceCountTrendResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceCountTrendResponse) SetStatusCode(v int32) *GetInstanceCountTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceCountTrendResponse) SetBody(v *GetInstanceCountTrendResponseBody) *GetInstanceCountTrendResponse {
	s.Body = v
	return s
}

func (s *GetInstanceCountTrendResponse) Validate() error {
	return dara.Validate(s)
}
