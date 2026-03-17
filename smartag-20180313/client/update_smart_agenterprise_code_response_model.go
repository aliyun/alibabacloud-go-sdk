// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAGEnterpriseCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmartAGEnterpriseCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmartAGEnterpriseCodeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmartAGEnterpriseCodeResponseBody) *UpdateSmartAGEnterpriseCodeResponse
	GetBody() *UpdateSmartAGEnterpriseCodeResponseBody
}

type UpdateSmartAGEnterpriseCodeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmartAGEnterpriseCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmartAGEnterpriseCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAGEnterpriseCodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartAGEnterpriseCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmartAGEnterpriseCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmartAGEnterpriseCodeResponse) GetBody() *UpdateSmartAGEnterpriseCodeResponseBody {
	return s.Body
}

func (s *UpdateSmartAGEnterpriseCodeResponse) SetHeaders(v map[string]*string) *UpdateSmartAGEnterpriseCodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartAGEnterpriseCodeResponse) SetStatusCode(v int32) *UpdateSmartAGEnterpriseCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmartAGEnterpriseCodeResponse) SetBody(v *UpdateSmartAGEnterpriseCodeResponseBody) *UpdateSmartAGEnterpriseCodeResponse {
	s.Body = v
	return s
}

func (s *UpdateSmartAGEnterpriseCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
