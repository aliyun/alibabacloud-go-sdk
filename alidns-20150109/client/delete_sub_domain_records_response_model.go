// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubDomainRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSubDomainRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSubDomainRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSubDomainRecordsResponseBody) *DeleteSubDomainRecordsResponse
	GetBody() *DeleteSubDomainRecordsResponseBody
}

type DeleteSubDomainRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSubDomainRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSubDomainRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubDomainRecordsResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubDomainRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSubDomainRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSubDomainRecordsResponse) GetBody() *DeleteSubDomainRecordsResponseBody {
	return s.Body
}

func (s *DeleteSubDomainRecordsResponse) SetHeaders(v map[string]*string) *DeleteSubDomainRecordsResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubDomainRecordsResponse) SetStatusCode(v int32) *DeleteSubDomainRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubDomainRecordsResponse) SetBody(v *DeleteSubDomainRecordsResponseBody) *DeleteSubDomainRecordsResponse {
	s.Body = v
	return s
}

func (s *DeleteSubDomainRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
