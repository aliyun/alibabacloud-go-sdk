// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHibernateDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HibernateDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HibernateDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *HibernateDesktopsResponseBody) *HibernateDesktopsResponse
	GetBody() *HibernateDesktopsResponseBody
}

type HibernateDesktopsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HibernateDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HibernateDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s HibernateDesktopsResponse) GoString() string {
	return s.String()
}

func (s *HibernateDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HibernateDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HibernateDesktopsResponse) GetBody() *HibernateDesktopsResponseBody {
	return s.Body
}

func (s *HibernateDesktopsResponse) SetHeaders(v map[string]*string) *HibernateDesktopsResponse {
	s.Headers = v
	return s
}

func (s *HibernateDesktopsResponse) SetStatusCode(v int32) *HibernateDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *HibernateDesktopsResponse) SetBody(v *HibernateDesktopsResponseBody) *HibernateDesktopsResponse {
	s.Body = v
	return s
}

func (s *HibernateDesktopsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
