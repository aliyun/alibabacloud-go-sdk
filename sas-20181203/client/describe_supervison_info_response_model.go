// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupervisonInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSupervisonInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSupervisonInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSupervisonInfoResponseBody) *DescribeSupervisonInfoResponse
	GetBody() *DescribeSupervisonInfoResponseBody
}

type DescribeSupervisonInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupervisonInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupervisonInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupervisonInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupervisonInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSupervisonInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSupervisonInfoResponse) GetBody() *DescribeSupervisonInfoResponseBody {
	return s.Body
}

func (s *DescribeSupervisonInfoResponse) SetHeaders(v map[string]*string) *DescribeSupervisonInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupervisonInfoResponse) SetStatusCode(v int32) *DescribeSupervisonInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupervisonInfoResponse) SetBody(v *DescribeSupervisonInfoResponseBody) *DescribeSupervisonInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeSupervisonInfoResponse) Validate() error {
	return dara.Validate(s)
}
