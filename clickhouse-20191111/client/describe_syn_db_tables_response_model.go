// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynDbTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSynDbTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSynDbTablesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSynDbTablesResponseBody) *DescribeSynDbTablesResponse
	GetBody() *DescribeSynDbTablesResponseBody
}

type DescribeSynDbTablesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSynDbTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSynDbTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynDbTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynDbTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSynDbTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSynDbTablesResponse) GetBody() *DescribeSynDbTablesResponseBody {
	return s.Body
}

func (s *DescribeSynDbTablesResponse) SetHeaders(v map[string]*string) *DescribeSynDbTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynDbTablesResponse) SetStatusCode(v int32) *DescribeSynDbTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynDbTablesResponse) SetBody(v *DescribeSynDbTablesResponseBody) *DescribeSynDbTablesResponse {
	s.Body = v
	return s
}

func (s *DescribeSynDbTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
