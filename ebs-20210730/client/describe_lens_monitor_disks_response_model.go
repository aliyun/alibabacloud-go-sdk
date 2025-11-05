// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLensMonitorDisksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLensMonitorDisksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLensMonitorDisksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLensMonitorDisksResponseBody) *DescribeLensMonitorDisksResponse
	GetBody() *DescribeLensMonitorDisksResponseBody
}

type DescribeLensMonitorDisksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLensMonitorDisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLensMonitorDisksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLensMonitorDisksResponse) GoString() string {
	return s.String()
}

func (s *DescribeLensMonitorDisksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLensMonitorDisksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLensMonitorDisksResponse) GetBody() *DescribeLensMonitorDisksResponseBody {
	return s.Body
}

func (s *DescribeLensMonitorDisksResponse) SetHeaders(v map[string]*string) *DescribeLensMonitorDisksResponse {
	s.Headers = v
	return s
}

func (s *DescribeLensMonitorDisksResponse) SetStatusCode(v int32) *DescribeLensMonitorDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLensMonitorDisksResponse) SetBody(v *DescribeLensMonitorDisksResponseBody) *DescribeLensMonitorDisksResponse {
	s.Body = v
	return s
}

func (s *DescribeLensMonitorDisksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
