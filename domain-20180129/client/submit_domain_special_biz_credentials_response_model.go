// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDomainSpecialBizCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDomainSpecialBizCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDomainSpecialBizCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDomainSpecialBizCredentialsResponseBody) *SubmitDomainSpecialBizCredentialsResponse
	GetBody() *SubmitDomainSpecialBizCredentialsResponseBody
}

type SubmitDomainSpecialBizCredentialsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDomainSpecialBizCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDomainSpecialBizCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDomainSpecialBizCredentialsResponse) GoString() string {
	return s.String()
}

func (s *SubmitDomainSpecialBizCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDomainSpecialBizCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDomainSpecialBizCredentialsResponse) GetBody() *SubmitDomainSpecialBizCredentialsResponseBody {
	return s.Body
}

func (s *SubmitDomainSpecialBizCredentialsResponse) SetHeaders(v map[string]*string) *SubmitDomainSpecialBizCredentialsResponse {
	s.Headers = v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponse) SetStatusCode(v int32) *SubmitDomainSpecialBizCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponse) SetBody(v *SubmitDomainSpecialBizCredentialsResponseBody) *SubmitDomainSpecialBizCredentialsResponse {
	s.Body = v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
