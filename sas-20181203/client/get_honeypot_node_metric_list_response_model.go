// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotNodeMetricListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHoneypotNodeMetricListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHoneypotNodeMetricListResponse
	GetStatusCode() *int32
	SetBody(v *GetHoneypotNodeMetricListResponseBody) *GetHoneypotNodeMetricListResponse
	GetBody() *GetHoneypotNodeMetricListResponseBody
}

type GetHoneypotNodeMetricListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHoneypotNodeMetricListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHoneypotNodeMetricListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotNodeMetricListResponse) GoString() string {
	return s.String()
}

func (s *GetHoneypotNodeMetricListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHoneypotNodeMetricListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHoneypotNodeMetricListResponse) GetBody() *GetHoneypotNodeMetricListResponseBody {
	return s.Body
}

func (s *GetHoneypotNodeMetricListResponse) SetHeaders(v map[string]*string) *GetHoneypotNodeMetricListResponse {
	s.Headers = v
	return s
}

func (s *GetHoneypotNodeMetricListResponse) SetStatusCode(v int32) *GetHoneypotNodeMetricListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHoneypotNodeMetricListResponse) SetBody(v *GetHoneypotNodeMetricListResponseBody) *GetHoneypotNodeMetricListResponse {
	s.Body = v
	return s
}

func (s *GetHoneypotNodeMetricListResponse) Validate() error {
	return dara.Validate(s)
}
