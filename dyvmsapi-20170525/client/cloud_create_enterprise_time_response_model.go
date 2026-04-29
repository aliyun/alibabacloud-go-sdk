// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateEnterpriseTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudCreateEnterpriseTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudCreateEnterpriseTimeResponse
	GetStatusCode() *int32
	SetBody(v *CloudCreateEnterpriseTimeResponseBody) *CloudCreateEnterpriseTimeResponse
	GetBody() *CloudCreateEnterpriseTimeResponseBody
}

type CloudCreateEnterpriseTimeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudCreateEnterpriseTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudCreateEnterpriseTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateEnterpriseTimeResponse) GoString() string {
	return s.String()
}

func (s *CloudCreateEnterpriseTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudCreateEnterpriseTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudCreateEnterpriseTimeResponse) GetBody() *CloudCreateEnterpriseTimeResponseBody {
	return s.Body
}

func (s *CloudCreateEnterpriseTimeResponse) SetHeaders(v map[string]*string) *CloudCreateEnterpriseTimeResponse {
	s.Headers = v
	return s
}

func (s *CloudCreateEnterpriseTimeResponse) SetStatusCode(v int32) *CloudCreateEnterpriseTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudCreateEnterpriseTimeResponse) SetBody(v *CloudCreateEnterpriseTimeResponseBody) *CloudCreateEnterpriseTimeResponse {
	s.Body = v
	return s
}

func (s *CloudCreateEnterpriseTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
