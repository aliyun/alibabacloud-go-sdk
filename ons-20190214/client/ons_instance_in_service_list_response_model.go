// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceInServiceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsInstanceInServiceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsInstanceInServiceListResponse
	GetStatusCode() *int32
	SetBody(v *OnsInstanceInServiceListResponseBody) *OnsInstanceInServiceListResponse
	GetBody() *OnsInstanceInServiceListResponseBody
}

type OnsInstanceInServiceListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsInstanceInServiceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsInstanceInServiceListResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceInServiceListResponse) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsInstanceInServiceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsInstanceInServiceListResponse) GetBody() *OnsInstanceInServiceListResponseBody {
	return s.Body
}

func (s *OnsInstanceInServiceListResponse) SetHeaders(v map[string]*string) *OnsInstanceInServiceListResponse {
	s.Headers = v
	return s
}

func (s *OnsInstanceInServiceListResponse) SetStatusCode(v int32) *OnsInstanceInServiceListResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsInstanceInServiceListResponse) SetBody(v *OnsInstanceInServiceListResponseBody) *OnsInstanceInServiceListResponse {
	s.Body = v
	return s
}

func (s *OnsInstanceInServiceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
