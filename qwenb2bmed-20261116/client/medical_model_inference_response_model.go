// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalModelInferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MedicalModelInferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MedicalModelInferenceResponse
	GetStatusCode() *int32
	SetId(v string) *MedicalModelInferenceResponse
	GetId() *string
	SetEvent(v string) *MedicalModelInferenceResponse
	GetEvent() *string
	SetBody(v *MedicalModelInferenceResponseBody) *MedicalModelInferenceResponse
	GetBody() *MedicalModelInferenceResponseBody
}

type MedicalModelInferenceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *MedicalModelInferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MedicalModelInferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceResponse) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MedicalModelInferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MedicalModelInferenceResponse) GetId() *string {
	return s.Id
}

func (s *MedicalModelInferenceResponse) GetEvent() *string {
	return s.Event
}

func (s *MedicalModelInferenceResponse) GetBody() *MedicalModelInferenceResponseBody {
	return s.Body
}

func (s *MedicalModelInferenceResponse) SetHeaders(v map[string]*string) *MedicalModelInferenceResponse {
	s.Headers = v
	return s
}

func (s *MedicalModelInferenceResponse) SetStatusCode(v int32) *MedicalModelInferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *MedicalModelInferenceResponse) SetId(v string) *MedicalModelInferenceResponse {
	s.Id = &v
	return s
}

func (s *MedicalModelInferenceResponse) SetEvent(v string) *MedicalModelInferenceResponse {
	s.Event = &v
	return s
}

func (s *MedicalModelInferenceResponse) SetBody(v *MedicalModelInferenceResponseBody) *MedicalModelInferenceResponse {
	s.Body = v
	return s
}

func (s *MedicalModelInferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
