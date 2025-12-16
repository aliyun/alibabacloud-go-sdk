// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterventionDictionariesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInterventionDictionariesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInterventionDictionariesResponse
	GetStatusCode() *int32
	SetBody(v *ListInterventionDictionariesResponseBody) *ListInterventionDictionariesResponse
	GetBody() *ListInterventionDictionariesResponseBody
}

type ListInterventionDictionariesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterventionDictionariesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterventionDictionariesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionariesResponse) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionariesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInterventionDictionariesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInterventionDictionariesResponse) GetBody() *ListInterventionDictionariesResponseBody {
	return s.Body
}

func (s *ListInterventionDictionariesResponse) SetHeaders(v map[string]*string) *ListInterventionDictionariesResponse {
	s.Headers = v
	return s
}

func (s *ListInterventionDictionariesResponse) SetStatusCode(v int32) *ListInterventionDictionariesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterventionDictionariesResponse) SetBody(v *ListInterventionDictionariesResponseBody) *ListInterventionDictionariesResponse {
	s.Body = v
	return s
}

func (s *ListInterventionDictionariesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
