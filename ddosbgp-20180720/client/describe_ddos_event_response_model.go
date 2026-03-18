// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDdosEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDdosEventResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDdosEventResponseBody) *DescribeDdosEventResponse
	GetBody() *DescribeDdosEventResponseBody
}

type DescribeDdosEventResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDdosEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDdosEventResponse) GetBody() *DescribeDdosEventResponseBody {
	return s.Body
}

func (s *DescribeDdosEventResponse) SetHeaders(v map[string]*string) *DescribeDdosEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosEventResponse) SetStatusCode(v int32) *DescribeDdosEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosEventResponse) SetBody(v *DescribeDdosEventResponseBody) *DescribeDdosEventResponse {
	s.Body = v
	return s
}

func (s *DescribeDdosEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
