// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToPdfJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitConvertImageToPdfJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitConvertImageToPdfJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitConvertImageToPdfJobResponseBody) *SubmitConvertImageToPdfJobResponse
	GetBody() *SubmitConvertImageToPdfJobResponseBody
}

type SubmitConvertImageToPdfJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertImageToPdfJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitConvertImageToPdfJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToPdfJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToPdfJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitConvertImageToPdfJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitConvertImageToPdfJobResponse) GetBody() *SubmitConvertImageToPdfJobResponseBody {
	return s.Body
}

func (s *SubmitConvertImageToPdfJobResponse) SetHeaders(v map[string]*string) *SubmitConvertImageToPdfJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertImageToPdfJobResponse) SetStatusCode(v int32) *SubmitConvertImageToPdfJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertImageToPdfJobResponse) SetBody(v *SubmitConvertImageToPdfJobResponseBody) *SubmitConvertImageToPdfJobResponse {
	s.Body = v
	return s
}

func (s *SubmitConvertImageToPdfJobResponse) Validate() error {
	return dara.Validate(s)
}
