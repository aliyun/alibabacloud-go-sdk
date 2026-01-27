// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportTaskValidationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImportTaskValidationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImportTaskValidationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImportTaskValidationResponseBody) *DescribeImportTaskValidationResponse
	GetBody() *DescribeImportTaskValidationResponseBody
}

type DescribeImportTaskValidationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImportTaskValidationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImportTaskValidationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportTaskValidationResponse) GoString() string {
	return s.String()
}

func (s *DescribeImportTaskValidationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImportTaskValidationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImportTaskValidationResponse) GetBody() *DescribeImportTaskValidationResponseBody {
	return s.Body
}

func (s *DescribeImportTaskValidationResponse) SetHeaders(v map[string]*string) *DescribeImportTaskValidationResponse {
	s.Headers = v
	return s
}

func (s *DescribeImportTaskValidationResponse) SetStatusCode(v int32) *DescribeImportTaskValidationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImportTaskValidationResponse) SetBody(v *DescribeImportTaskValidationResponseBody) *DescribeImportTaskValidationResponse {
	s.Body = v
	return s
}

func (s *DescribeImportTaskValidationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
