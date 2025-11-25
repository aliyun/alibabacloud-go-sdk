// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStsGrantStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStsGrantStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStsGrantStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStsGrantStatusResponseBody) *DescribeStsGrantStatusResponse
	GetBody() *DescribeStsGrantStatusResponseBody
}

type DescribeStsGrantStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStsGrantStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStsGrantStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStsGrantStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeStsGrantStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStsGrantStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStsGrantStatusResponse) GetBody() *DescribeStsGrantStatusResponseBody {
	return s.Body
}

func (s *DescribeStsGrantStatusResponse) SetHeaders(v map[string]*string) *DescribeStsGrantStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeStsGrantStatusResponse) SetStatusCode(v int32) *DescribeStsGrantStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStsGrantStatusResponse) SetBody(v *DescribeStsGrantStatusResponseBody) *DescribeStsGrantStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeStsGrantStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
