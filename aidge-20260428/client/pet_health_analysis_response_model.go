// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPetHealthAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PetHealthAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PetHealthAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *PetHealthAnalysisResponseBody) *PetHealthAnalysisResponse
	GetBody() *PetHealthAnalysisResponseBody
}

type PetHealthAnalysisResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PetHealthAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PetHealthAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s PetHealthAnalysisResponse) GoString() string {
	return s.String()
}

func (s *PetHealthAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PetHealthAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PetHealthAnalysisResponse) GetBody() *PetHealthAnalysisResponseBody {
	return s.Body
}

func (s *PetHealthAnalysisResponse) SetHeaders(v map[string]*string) *PetHealthAnalysisResponse {
	s.Headers = v
	return s
}

func (s *PetHealthAnalysisResponse) SetStatusCode(v int32) *PetHealthAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *PetHealthAnalysisResponse) SetBody(v *PetHealthAnalysisResponseBody) *PetHealthAnalysisResponse {
	s.Body = v
	return s
}

func (s *PetHealthAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
