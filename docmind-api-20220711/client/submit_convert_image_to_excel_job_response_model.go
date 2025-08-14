// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToExcelJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitConvertImageToExcelJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitConvertImageToExcelJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitConvertImageToExcelJobResponseBody) *SubmitConvertImageToExcelJobResponse
	GetBody() *SubmitConvertImageToExcelJobResponseBody
}

type SubmitConvertImageToExcelJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitConvertImageToExcelJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitConvertImageToExcelJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToExcelJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToExcelJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitConvertImageToExcelJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitConvertImageToExcelJobResponse) GetBody() *SubmitConvertImageToExcelJobResponseBody {
	return s.Body
}

func (s *SubmitConvertImageToExcelJobResponse) SetHeaders(v map[string]*string) *SubmitConvertImageToExcelJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitConvertImageToExcelJobResponse) SetStatusCode(v int32) *SubmitConvertImageToExcelJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitConvertImageToExcelJobResponse) SetBody(v *SubmitConvertImageToExcelJobResponseBody) *SubmitConvertImageToExcelJobResponse {
	s.Body = v
	return s
}

func (s *SubmitConvertImageToExcelJobResponse) Validate() error {
	return dara.Validate(s)
}
