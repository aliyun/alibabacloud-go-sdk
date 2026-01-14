// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpsetsBandwidthLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIpsetsBandwidthLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIpsetsBandwidthLimitResponse
	GetStatusCode() *int32
	SetBody(v *GetIpsetsBandwidthLimitResponseBody) *GetIpsetsBandwidthLimitResponse
	GetBody() *GetIpsetsBandwidthLimitResponseBody
}

type GetIpsetsBandwidthLimitResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIpsetsBandwidthLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIpsetsBandwidthLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIpsetsBandwidthLimitResponse) GoString() string {
	return s.String()
}

func (s *GetIpsetsBandwidthLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIpsetsBandwidthLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIpsetsBandwidthLimitResponse) GetBody() *GetIpsetsBandwidthLimitResponseBody {
	return s.Body
}

func (s *GetIpsetsBandwidthLimitResponse) SetHeaders(v map[string]*string) *GetIpsetsBandwidthLimitResponse {
	s.Headers = v
	return s
}

func (s *GetIpsetsBandwidthLimitResponse) SetStatusCode(v int32) *GetIpsetsBandwidthLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIpsetsBandwidthLimitResponse) SetBody(v *GetIpsetsBandwidthLimitResponseBody) *GetIpsetsBandwidthLimitResponse {
	s.Body = v
	return s
}

func (s *GetIpsetsBandwidthLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
