// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListEnterpriseTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudListEnterpriseTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudListEnterpriseTimeResponse
	GetStatusCode() *int32
	SetBody(v *CloudListEnterpriseTimeResponseBody) *CloudListEnterpriseTimeResponse
	GetBody() *CloudListEnterpriseTimeResponseBody
}

type CloudListEnterpriseTimeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudListEnterpriseTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudListEnterpriseTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudListEnterpriseTimeResponse) GoString() string {
	return s.String()
}

func (s *CloudListEnterpriseTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudListEnterpriseTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudListEnterpriseTimeResponse) GetBody() *CloudListEnterpriseTimeResponseBody {
	return s.Body
}

func (s *CloudListEnterpriseTimeResponse) SetHeaders(v map[string]*string) *CloudListEnterpriseTimeResponse {
	s.Headers = v
	return s
}

func (s *CloudListEnterpriseTimeResponse) SetStatusCode(v int32) *CloudListEnterpriseTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudListEnterpriseTimeResponse) SetBody(v *CloudListEnterpriseTimeResponseBody) *CloudListEnterpriseTimeResponse {
	s.Body = v
	return s
}

func (s *CloudListEnterpriseTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
