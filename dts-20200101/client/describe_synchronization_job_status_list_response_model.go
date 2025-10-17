// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobStatusListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSynchronizationJobStatusListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSynchronizationJobStatusListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSynchronizationJobStatusListResponseBody) *DescribeSynchronizationJobStatusListResponse
	GetBody() *DescribeSynchronizationJobStatusListResponseBody
}

type DescribeSynchronizationJobStatusListResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSynchronizationJobStatusListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSynchronizationJobStatusListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSynchronizationJobStatusListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSynchronizationJobStatusListResponse) GetBody() *DescribeSynchronizationJobStatusListResponseBody {
	return s.Body
}

func (s *DescribeSynchronizationJobStatusListResponse) SetHeaders(v map[string]*string) *DescribeSynchronizationJobStatusListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponse) SetStatusCode(v int32) *DescribeSynchronizationJobStatusListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponse) SetBody(v *DescribeSynchronizationJobStatusListResponseBody) *DescribeSynchronizationJobStatusListResponse {
	s.Body = v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
