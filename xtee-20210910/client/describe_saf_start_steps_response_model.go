// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafStartStepsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSafStartStepsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSafStartStepsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSafStartStepsResponseBody) *DescribeSafStartStepsResponse
	GetBody() *DescribeSafStartStepsResponseBody
}

type DescribeSafStartStepsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSafStartStepsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSafStartStepsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafStartStepsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSafStartStepsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSafStartStepsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSafStartStepsResponse) GetBody() *DescribeSafStartStepsResponseBody {
	return s.Body
}

func (s *DescribeSafStartStepsResponse) SetHeaders(v map[string]*string) *DescribeSafStartStepsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSafStartStepsResponse) SetStatusCode(v int32) *DescribeSafStartStepsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSafStartStepsResponse) SetBody(v *DescribeSafStartStepsResponseBody) *DescribeSafStartStepsResponse {
	s.Body = v
	return s
}

func (s *DescribeSafStartStepsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
