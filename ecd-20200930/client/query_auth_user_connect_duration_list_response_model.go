// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuthUserConnectDurationListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAuthUserConnectDurationListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAuthUserConnectDurationListResponse
	GetStatusCode() *int32
	SetBody(v *QueryAuthUserConnectDurationListResponseBody) *QueryAuthUserConnectDurationListResponse
	GetBody() *QueryAuthUserConnectDurationListResponseBody
}

type QueryAuthUserConnectDurationListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAuthUserConnectDurationListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAuthUserConnectDurationListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthUserConnectDurationListResponse) GoString() string {
	return s.String()
}

func (s *QueryAuthUserConnectDurationListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAuthUserConnectDurationListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAuthUserConnectDurationListResponse) GetBody() *QueryAuthUserConnectDurationListResponseBody {
	return s.Body
}

func (s *QueryAuthUserConnectDurationListResponse) SetHeaders(v map[string]*string) *QueryAuthUserConnectDurationListResponse {
	s.Headers = v
	return s
}

func (s *QueryAuthUserConnectDurationListResponse) SetStatusCode(v int32) *QueryAuthUserConnectDurationListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAuthUserConnectDurationListResponse) SetBody(v *QueryAuthUserConnectDurationListResponseBody) *QueryAuthUserConnectDurationListResponse {
	s.Body = v
	return s
}

func (s *QueryAuthUserConnectDurationListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
