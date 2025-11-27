// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlotsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlotsResponseBody) *DescribeSlotsResponse
	GetBody() *DescribeSlotsResponseBody
}

type DescribeSlotsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlotsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlotsResponse) GetBody() *DescribeSlotsResponseBody {
	return s.Body
}

func (s *DescribeSlotsResponse) SetHeaders(v map[string]*string) *DescribeSlotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlotsResponse) SetStatusCode(v int32) *DescribeSlotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlotsResponse) SetBody(v *DescribeSlotsResponseBody) *DescribeSlotsResponse {
	s.Body = v
	return s
}

func (s *DescribeSlotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
