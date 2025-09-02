// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgWhiteListDeleteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgWhiteListDeleteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgWhiteListDeleteListResponse
	GetStatusCode() *int32
	SetBody(v *DsgWhiteListDeleteListResponseBody) *DsgWhiteListDeleteListResponse
	GetBody() *DsgWhiteListDeleteListResponseBody
}

type DsgWhiteListDeleteListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgWhiteListDeleteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgWhiteListDeleteListResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListDeleteListResponse) GoString() string {
	return s.String()
}

func (s *DsgWhiteListDeleteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgWhiteListDeleteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgWhiteListDeleteListResponse) GetBody() *DsgWhiteListDeleteListResponseBody {
	return s.Body
}

func (s *DsgWhiteListDeleteListResponse) SetHeaders(v map[string]*string) *DsgWhiteListDeleteListResponse {
	s.Headers = v
	return s
}

func (s *DsgWhiteListDeleteListResponse) SetStatusCode(v int32) *DsgWhiteListDeleteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgWhiteListDeleteListResponse) SetBody(v *DsgWhiteListDeleteListResponseBody) *DsgWhiteListDeleteListResponse {
	s.Body = v
	return s
}

func (s *DsgWhiteListDeleteListResponse) Validate() error {
	return dara.Validate(s)
}
