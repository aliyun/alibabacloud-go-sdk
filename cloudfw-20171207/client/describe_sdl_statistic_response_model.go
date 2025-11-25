// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSdlStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSdlStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSdlStatisticResponseBody) *DescribeSdlStatisticResponse
	GetBody() *DescribeSdlStatisticResponseBody
}

type DescribeSdlStatisticResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSdlStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSdlStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeSdlStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSdlStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSdlStatisticResponse) GetBody() *DescribeSdlStatisticResponseBody {
	return s.Body
}

func (s *DescribeSdlStatisticResponse) SetHeaders(v map[string]*string) *DescribeSdlStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeSdlStatisticResponse) SetStatusCode(v int32) *DescribeSdlStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSdlStatisticResponse) SetBody(v *DescribeSdlStatisticResponseBody) *DescribeSdlStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeSdlStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
