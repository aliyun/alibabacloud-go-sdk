// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmarttagJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSmarttagJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSmarttagJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSmarttagJobResponseBody) *SubmitSmarttagJobResponse
	GetBody() *SubmitSmarttagJobResponseBody
}

type SubmitSmarttagJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSmarttagJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSmarttagJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmarttagJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSmarttagJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSmarttagJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSmarttagJobResponse) GetBody() *SubmitSmarttagJobResponseBody {
	return s.Body
}

func (s *SubmitSmarttagJobResponse) SetHeaders(v map[string]*string) *SubmitSmarttagJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSmarttagJobResponse) SetStatusCode(v int32) *SubmitSmarttagJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSmarttagJobResponse) SetBody(v *SubmitSmarttagJobResponseBody) *SubmitSmarttagJobResponse {
	s.Body = v
	return s
}

func (s *SubmitSmarttagJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
