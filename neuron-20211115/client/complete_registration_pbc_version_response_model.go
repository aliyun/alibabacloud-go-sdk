// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteRegistrationPbcVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompleteRegistrationPbcVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompleteRegistrationPbcVersionResponse
	GetStatusCode() *int32
	SetBody(v *CatalogCommonResult) *CompleteRegistrationPbcVersionResponse
	GetBody() *CatalogCommonResult
}

type CompleteRegistrationPbcVersionResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CatalogCommonResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompleteRegistrationPbcVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CompleteRegistrationPbcVersionResponse) GoString() string {
	return s.String()
}

func (s *CompleteRegistrationPbcVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompleteRegistrationPbcVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompleteRegistrationPbcVersionResponse) GetBody() *CatalogCommonResult {
	return s.Body
}

func (s *CompleteRegistrationPbcVersionResponse) SetHeaders(v map[string]*string) *CompleteRegistrationPbcVersionResponse {
	s.Headers = v
	return s
}

func (s *CompleteRegistrationPbcVersionResponse) SetStatusCode(v int32) *CompleteRegistrationPbcVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteRegistrationPbcVersionResponse) SetBody(v *CatalogCommonResult) *CompleteRegistrationPbcVersionResponse {
	s.Body = v
	return s
}

func (s *CompleteRegistrationPbcVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
