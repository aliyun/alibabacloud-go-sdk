// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSynchronizationJobStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSynchronizationJobStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSynchronizationJobStatusResponseBody) *DescribeSynchronizationJobStatusResponse
	GetBody() *DescribeSynchronizationJobStatusResponseBody
}

type DescribeSynchronizationJobStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSynchronizationJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSynchronizationJobStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSynchronizationJobStatusResponse) GetBody() *DescribeSynchronizationJobStatusResponseBody {
	return s.Body
}

func (s *DescribeSynchronizationJobStatusResponse) SetHeaders(v map[string]*string) *DescribeSynchronizationJobStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponse) SetStatusCode(v int32) *DescribeSynchronizationJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponse) SetBody(v *DescribeSynchronizationJobStatusResponseBody) *DescribeSynchronizationJobStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
