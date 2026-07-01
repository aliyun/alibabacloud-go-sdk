// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeToRCSSignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeToRCSSignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeToRCSSignatureResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeToRCSSignatureResponseBody) *UpgradeToRCSSignatureResponse
	GetBody() *UpgradeToRCSSignatureResponseBody
}

type UpgradeToRCSSignatureResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeToRCSSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeToRCSSignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeToRCSSignatureResponse) GoString() string {
	return s.String()
}

func (s *UpgradeToRCSSignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeToRCSSignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeToRCSSignatureResponse) GetBody() *UpgradeToRCSSignatureResponseBody {
	return s.Body
}

func (s *UpgradeToRCSSignatureResponse) SetHeaders(v map[string]*string) *UpgradeToRCSSignatureResponse {
	s.Headers = v
	return s
}

func (s *UpgradeToRCSSignatureResponse) SetStatusCode(v int32) *UpgradeToRCSSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeToRCSSignatureResponse) SetBody(v *UpgradeToRCSSignatureResponseBody) *UpgradeToRCSSignatureResponse {
	s.Body = v
	return s
}

func (s *UpgradeToRCSSignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
