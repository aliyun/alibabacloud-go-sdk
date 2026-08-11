// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiAppWarningByPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAiAppWarningByPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAiAppWarningByPageResponse
	GetStatusCode() *int32
	SetBody(v *ListAiAppWarningByPageResponseBody) *ListAiAppWarningByPageResponse
	GetBody() *ListAiAppWarningByPageResponseBody
}

type ListAiAppWarningByPageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAiAppWarningByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAiAppWarningByPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppWarningByPageResponse) GoString() string {
	return s.String()
}

func (s *ListAiAppWarningByPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAiAppWarningByPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAiAppWarningByPageResponse) GetBody() *ListAiAppWarningByPageResponseBody {
	return s.Body
}

func (s *ListAiAppWarningByPageResponse) SetHeaders(v map[string]*string) *ListAiAppWarningByPageResponse {
	s.Headers = v
	return s
}

func (s *ListAiAppWarningByPageResponse) SetStatusCode(v int32) *ListAiAppWarningByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAiAppWarningByPageResponse) SetBody(v *ListAiAppWarningByPageResponseBody) *ListAiAppWarningByPageResponse {
	s.Body = v
	return s
}

func (s *ListAiAppWarningByPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
