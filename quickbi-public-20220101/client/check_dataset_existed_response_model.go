// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDatasetExistedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDatasetExistedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDatasetExistedResponse
	GetStatusCode() *int32
	SetBody(v *CheckDatasetExistedResponseBody) *CheckDatasetExistedResponse
	GetBody() *CheckDatasetExistedResponseBody
}

type CheckDatasetExistedResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDatasetExistedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDatasetExistedResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDatasetExistedResponse) GoString() string {
	return s.String()
}

func (s *CheckDatasetExistedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDatasetExistedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDatasetExistedResponse) GetBody() *CheckDatasetExistedResponseBody {
	return s.Body
}

func (s *CheckDatasetExistedResponse) SetHeaders(v map[string]*string) *CheckDatasetExistedResponse {
	s.Headers = v
	return s
}

func (s *CheckDatasetExistedResponse) SetStatusCode(v int32) *CheckDatasetExistedResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDatasetExistedResponse) SetBody(v *CheckDatasetExistedResponseBody) *CheckDatasetExistedResponse {
	s.Body = v
	return s
}

func (s *CheckDatasetExistedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
