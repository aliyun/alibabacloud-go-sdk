// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSynchronizationJobAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSynchronizationJobAlertResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSynchronizationJobAlertResponseBody) *DescribeSynchronizationJobAlertResponse
	GetBody() *DescribeSynchronizationJobAlertResponseBody
}

type DescribeSynchronizationJobAlertResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSynchronizationJobAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSynchronizationJobAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobAlertResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSynchronizationJobAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSynchronizationJobAlertResponse) GetBody() *DescribeSynchronizationJobAlertResponseBody {
	return s.Body
}

func (s *DescribeSynchronizationJobAlertResponse) SetHeaders(v map[string]*string) *DescribeSynchronizationJobAlertResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronizationJobAlertResponse) SetStatusCode(v int32) *DescribeSynchronizationJobAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponse) SetBody(v *DescribeSynchronizationJobAlertResponseBody) *DescribeSynchronizationJobAlertResponse {
	s.Body = v
	return s
}

func (s *DescribeSynchronizationJobAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
