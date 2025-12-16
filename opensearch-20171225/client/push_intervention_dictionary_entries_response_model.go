// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushInterventionDictionaryEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushInterventionDictionaryEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushInterventionDictionaryEntriesResponse
	GetStatusCode() *int32
	SetBody(v *PushInterventionDictionaryEntriesResponseBody) *PushInterventionDictionaryEntriesResponse
	GetBody() *PushInterventionDictionaryEntriesResponseBody
}

type PushInterventionDictionaryEntriesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushInterventionDictionaryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushInterventionDictionaryEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s PushInterventionDictionaryEntriesResponse) GoString() string {
	return s.String()
}

func (s *PushInterventionDictionaryEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushInterventionDictionaryEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushInterventionDictionaryEntriesResponse) GetBody() *PushInterventionDictionaryEntriesResponseBody {
	return s.Body
}

func (s *PushInterventionDictionaryEntriesResponse) SetHeaders(v map[string]*string) *PushInterventionDictionaryEntriesResponse {
	s.Headers = v
	return s
}

func (s *PushInterventionDictionaryEntriesResponse) SetStatusCode(v int32) *PushInterventionDictionaryEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *PushInterventionDictionaryEntriesResponse) SetBody(v *PushInterventionDictionaryEntriesResponseBody) *PushInterventionDictionaryEntriesResponse {
	s.Body = v
	return s
}

func (s *PushInterventionDictionaryEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
