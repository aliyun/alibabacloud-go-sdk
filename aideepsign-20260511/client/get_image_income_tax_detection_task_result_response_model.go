// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageIncomeTaxDetectionTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageIncomeTaxDetectionTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageIncomeTaxDetectionTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetImageIncomeTaxDetectionTaskResultResponseBody) *GetImageIncomeTaxDetectionTaskResultResponse
	GetBody() *GetImageIncomeTaxDetectionTaskResultResponseBody
}

type GetImageIncomeTaxDetectionTaskResultResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageIncomeTaxDetectionTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageIncomeTaxDetectionTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageIncomeTaxDetectionTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetImageIncomeTaxDetectionTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageIncomeTaxDetectionTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageIncomeTaxDetectionTaskResultResponse) GetBody() *GetImageIncomeTaxDetectionTaskResultResponseBody {
	return s.Body
}

func (s *GetImageIncomeTaxDetectionTaskResultResponse) SetHeaders(v map[string]*string) *GetImageIncomeTaxDetectionTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponse) SetStatusCode(v int32) *GetImageIncomeTaxDetectionTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponse) SetBody(v *GetImageIncomeTaxDetectionTaskResultResponseBody) *GetImageIncomeTaxDetectionTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
