// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsGroupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsGroupListResponse
	GetStatusCode() *int32
	SetBody(v *OnsGroupListResponseBody) *OnsGroupListResponse
	GetBody() *OnsGroupListResponseBody
}

type OnsGroupListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsGroupListResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupListResponse) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsGroupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsGroupListResponse) GetBody() *OnsGroupListResponseBody {
	return s.Body
}

func (s *OnsGroupListResponse) SetHeaders(v map[string]*string) *OnsGroupListResponse {
	s.Headers = v
	return s
}

func (s *OnsGroupListResponse) SetStatusCode(v int32) *OnsGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsGroupListResponse) SetBody(v *OnsGroupListResponseBody) *OnsGroupListResponse {
	s.Body = v
	return s
}

func (s *OnsGroupListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
