// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagWanSnatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagWanSnatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagWanSnatResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagWanSnatResponseBody) *DescribeSagWanSnatResponse
	GetBody() *DescribeSagWanSnatResponseBody
}

type DescribeSagWanSnatResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagWanSnatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagWanSnatResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWanSnatResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagWanSnatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagWanSnatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagWanSnatResponse) GetBody() *DescribeSagWanSnatResponseBody {
	return s.Body
}

func (s *DescribeSagWanSnatResponse) SetHeaders(v map[string]*string) *DescribeSagWanSnatResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagWanSnatResponse) SetStatusCode(v int32) *DescribeSagWanSnatResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagWanSnatResponse) SetBody(v *DescribeSagWanSnatResponseBody) *DescribeSagWanSnatResponse {
	s.Body = v
	return s
}

func (s *DescribeSagWanSnatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
