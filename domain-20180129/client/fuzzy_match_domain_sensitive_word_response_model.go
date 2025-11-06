// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFuzzyMatchDomainSensitiveWordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FuzzyMatchDomainSensitiveWordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FuzzyMatchDomainSensitiveWordResponse
	GetStatusCode() *int32
	SetBody(v *FuzzyMatchDomainSensitiveWordResponseBody) *FuzzyMatchDomainSensitiveWordResponse
	GetBody() *FuzzyMatchDomainSensitiveWordResponseBody
}

type FuzzyMatchDomainSensitiveWordResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FuzzyMatchDomainSensitiveWordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FuzzyMatchDomainSensitiveWordResponse) String() string {
	return dara.Prettify(s)
}

func (s FuzzyMatchDomainSensitiveWordResponse) GoString() string {
	return s.String()
}

func (s *FuzzyMatchDomainSensitiveWordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FuzzyMatchDomainSensitiveWordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FuzzyMatchDomainSensitiveWordResponse) GetBody() *FuzzyMatchDomainSensitiveWordResponseBody {
	return s.Body
}

func (s *FuzzyMatchDomainSensitiveWordResponse) SetHeaders(v map[string]*string) *FuzzyMatchDomainSensitiveWordResponse {
	s.Headers = v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponse) SetStatusCode(v int32) *FuzzyMatchDomainSensitiveWordResponse {
	s.StatusCode = &v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponse) SetBody(v *FuzzyMatchDomainSensitiveWordResponseBody) *FuzzyMatchDomainSensitiveWordResponse {
	s.Body = v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
