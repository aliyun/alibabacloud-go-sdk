// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessageDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsMessageDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsMessageDetailResponse
	GetStatusCode() *int32
	SetBody(v *OnsMessageDetailResponseBody) *OnsMessageDetailResponse
	GetBody() *OnsMessageDetailResponseBody
}

type OnsMessageDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsMessageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsMessageDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageDetailResponse) GoString() string {
	return s.String()
}

func (s *OnsMessageDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsMessageDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsMessageDetailResponse) GetBody() *OnsMessageDetailResponseBody {
	return s.Body
}

func (s *OnsMessageDetailResponse) SetHeaders(v map[string]*string) *OnsMessageDetailResponse {
	s.Headers = v
	return s
}

func (s *OnsMessageDetailResponse) SetStatusCode(v int32) *OnsMessageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsMessageDetailResponse) SetBody(v *OnsMessageDetailResponseBody) *OnsMessageDetailResponse {
	s.Body = v
	return s
}

func (s *OnsMessageDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
