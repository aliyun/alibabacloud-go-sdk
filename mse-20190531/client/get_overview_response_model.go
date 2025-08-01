// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOverviewResponse
	GetStatusCode() *int32
	SetBody(v *GetOverviewResponseBody) *GetOverviewResponse
	GetBody() *GetOverviewResponseBody
}

type GetOverviewResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOverviewResponse) GetBody() *GetOverviewResponseBody {
	return s.Body
}

func (s *GetOverviewResponse) SetHeaders(v map[string]*string) *GetOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetOverviewResponse) SetStatusCode(v int32) *GetOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOverviewResponse) SetBody(v *GetOverviewResponseBody) *GetOverviewResponse {
	s.Body = v
	return s
}

func (s *GetOverviewResponse) Validate() error {
	return dara.Validate(s)
}
