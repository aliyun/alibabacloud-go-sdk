// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPostpaidSitePlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPostpaidSitePlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPostpaidSitePlansResponse
	GetStatusCode() *int32
	SetBody(v *ListPostpaidSitePlansResponseBody) *ListPostpaidSitePlansResponse
	GetBody() *ListPostpaidSitePlansResponseBody
}

type ListPostpaidSitePlansResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPostpaidSitePlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPostpaidSitePlansResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPostpaidSitePlansResponse) GoString() string {
	return s.String()
}

func (s *ListPostpaidSitePlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPostpaidSitePlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPostpaidSitePlansResponse) GetBody() *ListPostpaidSitePlansResponseBody {
	return s.Body
}

func (s *ListPostpaidSitePlansResponse) SetHeaders(v map[string]*string) *ListPostpaidSitePlansResponse {
	s.Headers = v
	return s
}

func (s *ListPostpaidSitePlansResponse) SetStatusCode(v int32) *ListPostpaidSitePlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPostpaidSitePlansResponse) SetBody(v *ListPostpaidSitePlansResponseBody) *ListPostpaidSitePlansResponse {
	s.Body = v
	return s
}

func (s *ListPostpaidSitePlansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
