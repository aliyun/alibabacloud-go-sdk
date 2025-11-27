// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLLogFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLLogFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLLogFilesResponseBody) *DescribeSQLLogFilesResponse
	GetBody() *DescribeSQLLogFilesResponseBody
}

type DescribeSQLLogFilesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLLogFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLLogFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLLogFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLLogFilesResponse) GetBody() *DescribeSQLLogFilesResponseBody {
	return s.Body
}

func (s *DescribeSQLLogFilesResponse) SetHeaders(v map[string]*string) *DescribeSQLLogFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLLogFilesResponse) SetStatusCode(v int32) *DescribeSQLLogFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLLogFilesResponse) SetBody(v *DescribeSQLLogFilesResponseBody) *DescribeSQLLogFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLLogFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
