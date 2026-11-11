// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHibernateClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HibernateClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HibernateClusterResponse
	GetStatusCode() *int32
	SetBody(v *HibernateClusterResponseBody) *HibernateClusterResponse
	GetBody() *HibernateClusterResponseBody
}

type HibernateClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HibernateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HibernateClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s HibernateClusterResponse) GoString() string {
	return s.String()
}

func (s *HibernateClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HibernateClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HibernateClusterResponse) GetBody() *HibernateClusterResponseBody {
	return s.Body
}

func (s *HibernateClusterResponse) SetHeaders(v map[string]*string) *HibernateClusterResponse {
	s.Headers = v
	return s
}

func (s *HibernateClusterResponse) SetStatusCode(v int32) *HibernateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *HibernateClusterResponse) SetBody(v *HibernateClusterResponseBody) *HibernateClusterResponse {
	s.Body = v
	return s
}

func (s *HibernateClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
