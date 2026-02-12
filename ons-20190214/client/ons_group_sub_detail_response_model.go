// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupSubDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsGroupSubDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsGroupSubDetailResponse
	GetStatusCode() *int32
	SetBody(v *OnsGroupSubDetailResponseBody) *OnsGroupSubDetailResponse
	GetBody() *OnsGroupSubDetailResponseBody
}

type OnsGroupSubDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsGroupSubDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsGroupSubDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupSubDetailResponse) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsGroupSubDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsGroupSubDetailResponse) GetBody() *OnsGroupSubDetailResponseBody {
	return s.Body
}

func (s *OnsGroupSubDetailResponse) SetHeaders(v map[string]*string) *OnsGroupSubDetailResponse {
	s.Headers = v
	return s
}

func (s *OnsGroupSubDetailResponse) SetStatusCode(v int32) *OnsGroupSubDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsGroupSubDetailResponse) SetBody(v *OnsGroupSubDetailResponseBody) *OnsGroupSubDetailResponse {
	s.Body = v
	return s
}

func (s *OnsGroupSubDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
