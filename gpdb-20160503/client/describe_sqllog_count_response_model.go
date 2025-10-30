// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLLogCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLLogCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLLogCountResponseBody) *DescribeSQLLogCountResponse
	GetBody() *DescribeSQLLogCountResponseBody
}

type DescribeSQLLogCountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLLogCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLLogCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLLogCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLLogCountResponse) GetBody() *DescribeSQLLogCountResponseBody {
	return s.Body
}

func (s *DescribeSQLLogCountResponse) SetHeaders(v map[string]*string) *DescribeSQLLogCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogCountResponse) SetStatusCode(v int32) *DescribeSQLLogCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogCountResponse) SetBody(v *DescribeSQLLogCountResponseBody) *DescribeSQLLogCountResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLLogCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
