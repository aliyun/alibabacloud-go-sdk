// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSensitiveSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSensitiveSwitchResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSensitiveSwitchResponseBody) *DescribeSensitiveSwitchResponse
	GetBody() *DescribeSensitiveSwitchResponseBody
}

type DescribeSensitiveSwitchResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSensitiveSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSensitiveSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveSwitchResponse) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSensitiveSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSensitiveSwitchResponse) GetBody() *DescribeSensitiveSwitchResponseBody {
	return s.Body
}

func (s *DescribeSensitiveSwitchResponse) SetHeaders(v map[string]*string) *DescribeSensitiveSwitchResponse {
	s.Headers = v
	return s
}

func (s *DescribeSensitiveSwitchResponse) SetStatusCode(v int32) *DescribeSensitiveSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSensitiveSwitchResponse) SetBody(v *DescribeSensitiveSwitchResponseBody) *DescribeSensitiveSwitchResponse {
	s.Body = v
	return s
}

func (s *DescribeSensitiveSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
