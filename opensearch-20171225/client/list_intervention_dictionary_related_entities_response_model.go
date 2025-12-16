// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterventionDictionaryRelatedEntitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInterventionDictionaryRelatedEntitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInterventionDictionaryRelatedEntitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListInterventionDictionaryRelatedEntitiesResponseBody) *ListInterventionDictionaryRelatedEntitiesResponse
	GetBody() *ListInterventionDictionaryRelatedEntitiesResponseBody
}

type ListInterventionDictionaryRelatedEntitiesResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterventionDictionaryRelatedEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterventionDictionaryRelatedEntitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionaryRelatedEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryRelatedEntitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInterventionDictionaryRelatedEntitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInterventionDictionaryRelatedEntitiesResponse) GetBody() *ListInterventionDictionaryRelatedEntitiesResponseBody {
	return s.Body
}

func (s *ListInterventionDictionaryRelatedEntitiesResponse) SetHeaders(v map[string]*string) *ListInterventionDictionaryRelatedEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListInterventionDictionaryRelatedEntitiesResponse) SetStatusCode(v int32) *ListInterventionDictionaryRelatedEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterventionDictionaryRelatedEntitiesResponse) SetBody(v *ListInterventionDictionaryRelatedEntitiesResponseBody) *ListInterventionDictionaryRelatedEntitiesResponse {
	s.Body = v
	return s
}

func (s *ListInterventionDictionaryRelatedEntitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
