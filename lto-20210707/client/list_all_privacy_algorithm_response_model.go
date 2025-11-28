// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllPrivacyAlgorithmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllPrivacyAlgorithmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllPrivacyAlgorithmResponse
	GetStatusCode() *int32
	SetBody(v *ListAllPrivacyAlgorithmResponseBody) *ListAllPrivacyAlgorithmResponse
	GetBody() *ListAllPrivacyAlgorithmResponseBody
}

type ListAllPrivacyAlgorithmResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllPrivacyAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllPrivacyAlgorithmResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllPrivacyAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *ListAllPrivacyAlgorithmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllPrivacyAlgorithmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllPrivacyAlgorithmResponse) GetBody() *ListAllPrivacyAlgorithmResponseBody {
	return s.Body
}

func (s *ListAllPrivacyAlgorithmResponse) SetHeaders(v map[string]*string) *ListAllPrivacyAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *ListAllPrivacyAlgorithmResponse) SetStatusCode(v int32) *ListAllPrivacyAlgorithmResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllPrivacyAlgorithmResponse) SetBody(v *ListAllPrivacyAlgorithmResponseBody) *ListAllPrivacyAlgorithmResponse {
	s.Body = v
	return s
}

func (s *ListAllPrivacyAlgorithmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
