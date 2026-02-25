// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaActiveUserUsagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQuotaActiveUserUsagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQuotaActiveUserUsagesResponse
	GetStatusCode() *int32
	SetBody(v *ListQuotaActiveUserUsagesResponseBody) *ListQuotaActiveUserUsagesResponse
	GetBody() *ListQuotaActiveUserUsagesResponseBody
}

type ListQuotaActiveUserUsagesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaActiveUserUsagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaActiveUserUsagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaActiveUserUsagesResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaActiveUserUsagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQuotaActiveUserUsagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQuotaActiveUserUsagesResponse) GetBody() *ListQuotaActiveUserUsagesResponseBody {
	return s.Body
}

func (s *ListQuotaActiveUserUsagesResponse) SetHeaders(v map[string]*string) *ListQuotaActiveUserUsagesResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaActiveUserUsagesResponse) SetStatusCode(v int32) *ListQuotaActiveUserUsagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaActiveUserUsagesResponse) SetBody(v *ListQuotaActiveUserUsagesResponseBody) *ListQuotaActiveUserUsagesResponse {
	s.Body = v
	return s
}

func (s *ListQuotaActiveUserUsagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
