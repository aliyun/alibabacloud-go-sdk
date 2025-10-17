// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverDBClusterZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FailoverDBClusterZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FailoverDBClusterZonalResponse
	GetStatusCode() *int32
	SetBody(v *FailoverDBClusterZonalResponseBody) *FailoverDBClusterZonalResponse
	GetBody() *FailoverDBClusterZonalResponseBody
}

type FailoverDBClusterZonalResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FailoverDBClusterZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FailoverDBClusterZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s FailoverDBClusterZonalResponse) GoString() string {
	return s.String()
}

func (s *FailoverDBClusterZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FailoverDBClusterZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FailoverDBClusterZonalResponse) GetBody() *FailoverDBClusterZonalResponseBody {
	return s.Body
}

func (s *FailoverDBClusterZonalResponse) SetHeaders(v map[string]*string) *FailoverDBClusterZonalResponse {
	s.Headers = v
	return s
}

func (s *FailoverDBClusterZonalResponse) SetStatusCode(v int32) *FailoverDBClusterZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *FailoverDBClusterZonalResponse) SetBody(v *FailoverDBClusterZonalResponseBody) *FailoverDBClusterZonalResponse {
	s.Body = v
	return s
}

func (s *FailoverDBClusterZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
