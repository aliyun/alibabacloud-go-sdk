// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDetailCdrIbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkDetailCdrIbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkDetailCdrIbResponse
	GetStatusCode() *int32
	SetBody(v *ClinkDetailCdrIbResponseBody) *ClinkDetailCdrIbResponse
	GetBody() *ClinkDetailCdrIbResponseBody
}

type ClinkDetailCdrIbResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkDetailCdrIbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkDetailCdrIbResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkDetailCdrIbResponse) GoString() string {
	return s.String()
}

func (s *ClinkDetailCdrIbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkDetailCdrIbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkDetailCdrIbResponse) GetBody() *ClinkDetailCdrIbResponseBody {
	return s.Body
}

func (s *ClinkDetailCdrIbResponse) SetHeaders(v map[string]*string) *ClinkDetailCdrIbResponse {
	s.Headers = v
	return s
}

func (s *ClinkDetailCdrIbResponse) SetStatusCode(v int32) *ClinkDetailCdrIbResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkDetailCdrIbResponse) SetBody(v *ClinkDetailCdrIbResponseBody) *ClinkDetailCdrIbResponse {
	s.Body = v
	return s
}

func (s *ClinkDetailCdrIbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
