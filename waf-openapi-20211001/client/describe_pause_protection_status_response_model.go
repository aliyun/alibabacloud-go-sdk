// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePauseProtectionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePauseProtectionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePauseProtectionStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribePauseProtectionStatusResponseBody) *DescribePauseProtectionStatusResponse
	GetBody() *DescribePauseProtectionStatusResponseBody
}

type DescribePauseProtectionStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePauseProtectionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePauseProtectionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePauseProtectionStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePauseProtectionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePauseProtectionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePauseProtectionStatusResponse) GetBody() *DescribePauseProtectionStatusResponseBody {
	return s.Body
}

func (s *DescribePauseProtectionStatusResponse) SetHeaders(v map[string]*string) *DescribePauseProtectionStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePauseProtectionStatusResponse) SetStatusCode(v int32) *DescribePauseProtectionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePauseProtectionStatusResponse) SetBody(v *DescribePauseProtectionStatusResponseBody) *DescribePauseProtectionStatusResponse {
	s.Body = v
	return s
}

func (s *DescribePauseProtectionStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
