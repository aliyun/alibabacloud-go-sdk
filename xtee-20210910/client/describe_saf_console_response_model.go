// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSafConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSafConsoleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSafConsoleResponseBody) *DescribeSafConsoleResponse
	GetBody() *DescribeSafConsoleResponseBody
}

type DescribeSafConsoleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSafConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSafConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafConsoleResponse) GoString() string {
	return s.String()
}

func (s *DescribeSafConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSafConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSafConsoleResponse) GetBody() *DescribeSafConsoleResponseBody {
	return s.Body
}

func (s *DescribeSafConsoleResponse) SetHeaders(v map[string]*string) *DescribeSafConsoleResponse {
	s.Headers = v
	return s
}

func (s *DescribeSafConsoleResponse) SetStatusCode(v int32) *DescribeSafConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSafConsoleResponse) SetBody(v *DescribeSafConsoleResponseBody) *DescribeSafConsoleResponse {
	s.Body = v
	return s
}

func (s *DescribeSafConsoleResponse) Validate() error {
	return dara.Validate(s)
}
