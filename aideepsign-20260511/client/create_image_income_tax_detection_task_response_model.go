// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageIncomeTaxDetectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageIncomeTaxDetectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageIncomeTaxDetectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageIncomeTaxDetectionTaskResponseBody) *CreateImageIncomeTaxDetectionTaskResponse
	GetBody() *CreateImageIncomeTaxDetectionTaskResponseBody
}

type CreateImageIncomeTaxDetectionTaskResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageIncomeTaxDetectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageIncomeTaxDetectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageIncomeTaxDetectionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageIncomeTaxDetectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageIncomeTaxDetectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageIncomeTaxDetectionTaskResponse) GetBody() *CreateImageIncomeTaxDetectionTaskResponseBody {
	return s.Body
}

func (s *CreateImageIncomeTaxDetectionTaskResponse) SetHeaders(v map[string]*string) *CreateImageIncomeTaxDetectionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageIncomeTaxDetectionTaskResponse) SetStatusCode(v int32) *CreateImageIncomeTaxDetectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageIncomeTaxDetectionTaskResponse) SetBody(v *CreateImageIncomeTaxDetectionTaskResponseBody) *CreateImageIncomeTaxDetectionTaskResponse {
	s.Body = v
	return s
}

func (s *CreateImageIncomeTaxDetectionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
