// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagManagementPortResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagManagementPortResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagManagementPortResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagManagementPortResponseBody) *DescribeSagManagementPortResponse
	GetBody() *DescribeSagManagementPortResponseBody
}

type DescribeSagManagementPortResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagManagementPortResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagManagementPortResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagManagementPortResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagManagementPortResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagManagementPortResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagManagementPortResponse) GetBody() *DescribeSagManagementPortResponseBody {
	return s.Body
}

func (s *DescribeSagManagementPortResponse) SetHeaders(v map[string]*string) *DescribeSagManagementPortResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagManagementPortResponse) SetStatusCode(v int32) *DescribeSagManagementPortResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagManagementPortResponse) SetBody(v *DescribeSagManagementPortResponseBody) *DescribeSagManagementPortResponse {
	s.Body = v
	return s
}

func (s *DescribeSagManagementPortResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
