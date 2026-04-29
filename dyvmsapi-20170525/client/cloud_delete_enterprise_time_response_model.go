// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteEnterpriseTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudDeleteEnterpriseTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudDeleteEnterpriseTimeResponse
	GetStatusCode() *int32
	SetBody(v *CloudDeleteEnterpriseTimeResponseBody) *CloudDeleteEnterpriseTimeResponse
	GetBody() *CloudDeleteEnterpriseTimeResponseBody
}

type CloudDeleteEnterpriseTimeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudDeleteEnterpriseTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudDeleteEnterpriseTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteEnterpriseTimeResponse) GoString() string {
	return s.String()
}

func (s *CloudDeleteEnterpriseTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudDeleteEnterpriseTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudDeleteEnterpriseTimeResponse) GetBody() *CloudDeleteEnterpriseTimeResponseBody {
	return s.Body
}

func (s *CloudDeleteEnterpriseTimeResponse) SetHeaders(v map[string]*string) *CloudDeleteEnterpriseTimeResponse {
	s.Headers = v
	return s
}

func (s *CloudDeleteEnterpriseTimeResponse) SetStatusCode(v int32) *CloudDeleteEnterpriseTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudDeleteEnterpriseTimeResponse) SetBody(v *CloudDeleteEnterpriseTimeResponseBody) *CloudDeleteEnterpriseTimeResponse {
	s.Body = v
	return s
}

func (s *CloudDeleteEnterpriseTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
