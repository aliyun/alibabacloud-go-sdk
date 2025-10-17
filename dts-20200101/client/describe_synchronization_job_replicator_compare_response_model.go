// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobReplicatorCompareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSynchronizationJobReplicatorCompareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSynchronizationJobReplicatorCompareResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSynchronizationJobReplicatorCompareResponseBody) *DescribeSynchronizationJobReplicatorCompareResponse
	GetBody() *DescribeSynchronizationJobReplicatorCompareResponseBody
}

type DescribeSynchronizationJobReplicatorCompareResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSynchronizationJobReplicatorCompareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSynchronizationJobReplicatorCompareResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobReplicatorCompareResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobReplicatorCompareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSynchronizationJobReplicatorCompareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSynchronizationJobReplicatorCompareResponse) GetBody() *DescribeSynchronizationJobReplicatorCompareResponseBody {
	return s.Body
}

func (s *DescribeSynchronizationJobReplicatorCompareResponse) SetHeaders(v map[string]*string) *DescribeSynchronizationJobReplicatorCompareResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponse) SetStatusCode(v int32) *DescribeSynchronizationJobReplicatorCompareResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponse) SetBody(v *DescribeSynchronizationJobReplicatorCompareResponseBody) *DescribeSynchronizationJobReplicatorCompareResponse {
	s.Body = v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
