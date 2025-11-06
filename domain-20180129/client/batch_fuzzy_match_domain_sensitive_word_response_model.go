// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchFuzzyMatchDomainSensitiveWordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchFuzzyMatchDomainSensitiveWordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchFuzzyMatchDomainSensitiveWordResponse
	GetStatusCode() *int32
	SetBody(v *BatchFuzzyMatchDomainSensitiveWordResponseBody) *BatchFuzzyMatchDomainSensitiveWordResponse
	GetBody() *BatchFuzzyMatchDomainSensitiveWordResponseBody
}

type BatchFuzzyMatchDomainSensitiveWordResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchFuzzyMatchDomainSensitiveWordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchFuzzyMatchDomainSensitiveWordResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordResponse) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponse) GetBody() *BatchFuzzyMatchDomainSensitiveWordResponseBody {
	return s.Body
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponse) SetHeaders(v map[string]*string) *BatchFuzzyMatchDomainSensitiveWordResponse {
	s.Headers = v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponse) SetStatusCode(v int32) *BatchFuzzyMatchDomainSensitiveWordResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponse) SetBody(v *BatchFuzzyMatchDomainSensitiveWordResponseBody) *BatchFuzzyMatchDomainSensitiveWordResponse {
	s.Body = v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
