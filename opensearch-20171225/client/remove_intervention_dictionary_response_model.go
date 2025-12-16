// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInterventionDictionaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveInterventionDictionaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveInterventionDictionaryResponse
	GetStatusCode() *int32
	SetBody(v *RemoveInterventionDictionaryResponseBody) *RemoveInterventionDictionaryResponse
	GetBody() *RemoveInterventionDictionaryResponseBody
}

type RemoveInterventionDictionaryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveInterventionDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveInterventionDictionaryResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveInterventionDictionaryResponse) GoString() string {
	return s.String()
}

func (s *RemoveInterventionDictionaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveInterventionDictionaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveInterventionDictionaryResponse) GetBody() *RemoveInterventionDictionaryResponseBody {
	return s.Body
}

func (s *RemoveInterventionDictionaryResponse) SetHeaders(v map[string]*string) *RemoveInterventionDictionaryResponse {
	s.Headers = v
	return s
}

func (s *RemoveInterventionDictionaryResponse) SetStatusCode(v int32) *RemoveInterventionDictionaryResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveInterventionDictionaryResponse) SetBody(v *RemoveInterventionDictionaryResponseBody) *RemoveInterventionDictionaryResponse {
	s.Body = v
	return s
}

func (s *RemoveInterventionDictionaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
