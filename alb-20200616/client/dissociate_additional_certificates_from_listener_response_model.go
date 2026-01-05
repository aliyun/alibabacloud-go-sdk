// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateAdditionalCertificatesFromListenerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateAdditionalCertificatesFromListenerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateAdditionalCertificatesFromListenerResponse
	GetStatusCode() *int32
	SetBody(v *DissociateAdditionalCertificatesFromListenerResponseBody) *DissociateAdditionalCertificatesFromListenerResponse
	GetBody() *DissociateAdditionalCertificatesFromListenerResponseBody
}

type DissociateAdditionalCertificatesFromListenerResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateAdditionalCertificatesFromListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateAdditionalCertificatesFromListenerResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerResponse) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) GetBody() *DissociateAdditionalCertificatesFromListenerResponseBody {
	return s.Body
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) SetHeaders(v map[string]*string) *DissociateAdditionalCertificatesFromListenerResponse {
	s.Headers = v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) SetStatusCode(v int32) *DissociateAdditionalCertificatesFromListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) SetBody(v *DissociateAdditionalCertificatesFromListenerResponseBody) *DissociateAdditionalCertificatesFromListenerResponse {
	s.Body = v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
