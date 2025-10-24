// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveRequestLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSensitiveRequestLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSensitiveRequestLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSensitiveRequestLogResponseBody) *DescribeSensitiveRequestLogResponse
	GetBody() *DescribeSensitiveRequestLogResponseBody
}

type DescribeSensitiveRequestLogResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSensitiveRequestLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSensitiveRequestLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveRequestLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveRequestLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSensitiveRequestLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSensitiveRequestLogResponse) GetBody() *DescribeSensitiveRequestLogResponseBody {
	return s.Body
}

func (s *DescribeSensitiveRequestLogResponse) SetHeaders(v map[string]*string) *DescribeSensitiveRequestLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeSensitiveRequestLogResponse) SetStatusCode(v int32) *DescribeSensitiveRequestLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSensitiveRequestLogResponse) SetBody(v *DescribeSensitiveRequestLogResponseBody) *DescribeSensitiveRequestLogResponse {
	s.Body = v
	return s
}

func (s *DescribeSensitiveRequestLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
