// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAifgAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAifgAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAifgAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *StartAifgAnalysisResponseBody) *StartAifgAnalysisResponse
	GetBody() *StartAifgAnalysisResponseBody
}

type StartAifgAnalysisResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAifgAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAifgAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAifgAnalysisResponse) GoString() string {
	return s.String()
}

func (s *StartAifgAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAifgAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAifgAnalysisResponse) GetBody() *StartAifgAnalysisResponseBody {
	return s.Body
}

func (s *StartAifgAnalysisResponse) SetHeaders(v map[string]*string) *StartAifgAnalysisResponse {
	s.Headers = v
	return s
}

func (s *StartAifgAnalysisResponse) SetStatusCode(v int32) *StartAifgAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAifgAnalysisResponse) SetBody(v *StartAifgAnalysisResponseBody) *StartAifgAnalysisResponse {
	s.Body = v
	return s
}

func (s *StartAifgAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
