// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmartAccessGatewaySoftwareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSmartAccessGatewaySoftwareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSmartAccessGatewaySoftwareResponse
	GetStatusCode() *int32
	SetBody(v *CreateSmartAccessGatewaySoftwareResponseBody) *CreateSmartAccessGatewaySoftwareResponse
	GetBody() *CreateSmartAccessGatewaySoftwareResponseBody
}

type CreateSmartAccessGatewaySoftwareResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmartAccessGatewaySoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmartAccessGatewaySoftwareResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSmartAccessGatewaySoftwareResponse) GoString() string {
	return s.String()
}

func (s *CreateSmartAccessGatewaySoftwareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSmartAccessGatewaySoftwareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSmartAccessGatewaySoftwareResponse) GetBody() *CreateSmartAccessGatewaySoftwareResponseBody {
	return s.Body
}

func (s *CreateSmartAccessGatewaySoftwareResponse) SetHeaders(v map[string]*string) *CreateSmartAccessGatewaySoftwareResponse {
	s.Headers = v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareResponse) SetStatusCode(v int32) *CreateSmartAccessGatewaySoftwareResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareResponse) SetBody(v *CreateSmartAccessGatewaySoftwareResponseBody) *CreateSmartAccessGatewaySoftwareResponse {
	s.Body = v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
