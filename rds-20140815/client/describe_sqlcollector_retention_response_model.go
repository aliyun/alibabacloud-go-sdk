// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLCollectorRetentionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLCollectorRetentionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLCollectorRetentionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLCollectorRetentionResponseBody) *DescribeSQLCollectorRetentionResponse
	GetBody() *DescribeSQLCollectorRetentionResponseBody
}

type DescribeSQLCollectorRetentionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLCollectorRetentionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLCollectorRetentionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLCollectorRetentionResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorRetentionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLCollectorRetentionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLCollectorRetentionResponse) GetBody() *DescribeSQLCollectorRetentionResponseBody {
	return s.Body
}

func (s *DescribeSQLCollectorRetentionResponse) SetHeaders(v map[string]*string) *DescribeSQLCollectorRetentionResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLCollectorRetentionResponse) SetStatusCode(v int32) *DescribeSQLCollectorRetentionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLCollectorRetentionResponse) SetBody(v *DescribeSQLCollectorRetentionResponseBody) *DescribeSQLCollectorRetentionResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLCollectorRetentionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
