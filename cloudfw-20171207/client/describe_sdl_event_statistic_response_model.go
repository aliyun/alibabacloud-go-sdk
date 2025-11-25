// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlEventStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSdlEventStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSdlEventStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSdlEventStatisticResponseBody) *DescribeSdlEventStatisticResponse
	GetBody() *DescribeSdlEventStatisticResponseBody
}

type DescribeSdlEventStatisticResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSdlEventStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSdlEventStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSdlEventStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSdlEventStatisticResponse) GetBody() *DescribeSdlEventStatisticResponseBody {
	return s.Body
}

func (s *DescribeSdlEventStatisticResponse) SetHeaders(v map[string]*string) *DescribeSdlEventStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeSdlEventStatisticResponse) SetStatusCode(v int32) *DescribeSdlEventStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSdlEventStatisticResponse) SetBody(v *DescribeSdlEventStatisticResponseBody) *DescribeSdlEventStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeSdlEventStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
