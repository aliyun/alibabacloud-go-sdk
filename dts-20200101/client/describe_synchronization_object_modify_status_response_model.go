// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationObjectModifyStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSynchronizationObjectModifyStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSynchronizationObjectModifyStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSynchronizationObjectModifyStatusResponseBody) *DescribeSynchronizationObjectModifyStatusResponse
	GetBody() *DescribeSynchronizationObjectModifyStatusResponseBody
}

type DescribeSynchronizationObjectModifyStatusResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSynchronizationObjectModifyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSynchronizationObjectModifyStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSynchronizationObjectModifyStatusResponse) GetBody() *DescribeSynchronizationObjectModifyStatusResponseBody {
	return s.Body
}

func (s *DescribeSynchronizationObjectModifyStatusResponse) SetHeaders(v map[string]*string) *DescribeSynchronizationObjectModifyStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponse) SetStatusCode(v int32) *DescribeSynchronizationObjectModifyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponse) SetBody(v *DescribeSynchronizationObjectModifyStatusResponseBody) *DescribeSynchronizationObjectModifyStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
