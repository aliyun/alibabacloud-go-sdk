// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterventionDictionaryEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInterventionDictionaryEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInterventionDictionaryEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ListInterventionDictionaryEntriesResponseBody) *ListInterventionDictionaryEntriesResponse
	GetBody() *ListInterventionDictionaryEntriesResponseBody
}

type ListInterventionDictionaryEntriesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterventionDictionaryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterventionDictionaryEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionaryEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInterventionDictionaryEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInterventionDictionaryEntriesResponse) GetBody() *ListInterventionDictionaryEntriesResponseBody {
	return s.Body
}

func (s *ListInterventionDictionaryEntriesResponse) SetHeaders(v map[string]*string) *ListInterventionDictionaryEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListInterventionDictionaryEntriesResponse) SetStatusCode(v int32) *ListInterventionDictionaryEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponse) SetBody(v *ListInterventionDictionaryEntriesResponseBody) *ListInterventionDictionaryEntriesResponse {
	s.Body = v
	return s
}

func (s *ListInterventionDictionaryEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
