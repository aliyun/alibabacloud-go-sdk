// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceListPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceListPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceListPageResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceListPageResponseBody) *GetServiceListPageResponse
	GetBody() *GetServiceListPageResponseBody
}

type GetServiceListPageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceListPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceListPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListPageResponse) GoString() string {
	return s.String()
}

func (s *GetServiceListPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceListPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceListPageResponse) GetBody() *GetServiceListPageResponseBody {
	return s.Body
}

func (s *GetServiceListPageResponse) SetHeaders(v map[string]*string) *GetServiceListPageResponse {
	s.Headers = v
	return s
}

func (s *GetServiceListPageResponse) SetStatusCode(v int32) *GetServiceListPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceListPageResponse) SetBody(v *GetServiceListPageResponseBody) *GetServiceListPageResponse {
	s.Body = v
	return s
}

func (s *GetServiceListPageResponse) Validate() error {
	return dara.Validate(s)
}
