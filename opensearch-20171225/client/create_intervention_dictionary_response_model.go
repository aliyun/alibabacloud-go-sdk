// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInterventionDictionaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInterventionDictionaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInterventionDictionaryResponse
	GetStatusCode() *int32
	SetBody(v *CreateInterventionDictionaryResponseBody) *CreateInterventionDictionaryResponse
	GetBody() *CreateInterventionDictionaryResponseBody
}

type CreateInterventionDictionaryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInterventionDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInterventionDictionaryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInterventionDictionaryResponse) GoString() string {
	return s.String()
}

func (s *CreateInterventionDictionaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInterventionDictionaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInterventionDictionaryResponse) GetBody() *CreateInterventionDictionaryResponseBody {
	return s.Body
}

func (s *CreateInterventionDictionaryResponse) SetHeaders(v map[string]*string) *CreateInterventionDictionaryResponse {
	s.Headers = v
	return s
}

func (s *CreateInterventionDictionaryResponse) SetStatusCode(v int32) *CreateInterventionDictionaryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInterventionDictionaryResponse) SetBody(v *CreateInterventionDictionaryResponseBody) *CreateInterventionDictionaryResponse {
	s.Body = v
	return s
}

func (s *CreateInterventionDictionaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
