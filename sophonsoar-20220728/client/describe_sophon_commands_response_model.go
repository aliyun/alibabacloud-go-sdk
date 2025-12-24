// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSophonCommandsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSophonCommandsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSophonCommandsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSophonCommandsResponseBody) *DescribeSophonCommandsResponse
	GetBody() *DescribeSophonCommandsResponseBody
}

type DescribeSophonCommandsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSophonCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSophonCommandsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSophonCommandsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSophonCommandsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSophonCommandsResponse) GetBody() *DescribeSophonCommandsResponseBody {
	return s.Body
}

func (s *DescribeSophonCommandsResponse) SetHeaders(v map[string]*string) *DescribeSophonCommandsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSophonCommandsResponse) SetStatusCode(v int32) *DescribeSophonCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSophonCommandsResponse) SetBody(v *DescribeSophonCommandsResponseBody) *DescribeSophonCommandsResponse {
	s.Body = v
	return s
}

func (s *DescribeSophonCommandsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
