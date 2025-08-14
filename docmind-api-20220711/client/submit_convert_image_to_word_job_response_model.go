// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToWordJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitConvertImageToWordJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitConvertImageToWordJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitConvertImageToWordJobResponseBody) *SubmitConvertImageToWordJobResponse
	GetBody() *SubmitConvertImageToWordJobResponseBody
}

type SubmitConvertImageToWordJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertImageToWordJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitConvertImageToWordJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToWordJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToWordJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitConvertImageToWordJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitConvertImageToWordJobResponse) GetBody() *SubmitConvertImageToWordJobResponseBody {
	return s.Body
}

func (s *SubmitConvertImageToWordJobResponse) SetHeaders(v map[string]*string) *SubmitConvertImageToWordJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertImageToWordJobResponse) SetStatusCode(v int32) *SubmitConvertImageToWordJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertImageToWordJobResponse) SetBody(v *SubmitConvertImageToWordJobResponseBody) *SubmitConvertImageToWordJobResponse {
	s.Body = v
	return s
}

func (s *SubmitConvertImageToWordJobResponse) Validate() error {
	return dara.Validate(s)
}
