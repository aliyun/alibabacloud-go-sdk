// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterventionDictionaryNerResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInterventionDictionaryNerResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInterventionDictionaryNerResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListInterventionDictionaryNerResultsResponseBody) *ListInterventionDictionaryNerResultsResponse
	GetBody() *ListInterventionDictionaryNerResultsResponseBody
}

type ListInterventionDictionaryNerResultsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterventionDictionaryNerResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterventionDictionaryNerResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionaryNerResultsResponse) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryNerResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInterventionDictionaryNerResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInterventionDictionaryNerResultsResponse) GetBody() *ListInterventionDictionaryNerResultsResponseBody {
	return s.Body
}

func (s *ListInterventionDictionaryNerResultsResponse) SetHeaders(v map[string]*string) *ListInterventionDictionaryNerResultsResponse {
	s.Headers = v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponse) SetStatusCode(v int32) *ListInterventionDictionaryNerResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponse) SetBody(v *ListInterventionDictionaryNerResultsResponseBody) *ListInterventionDictionaryNerResultsResponse {
	s.Body = v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
