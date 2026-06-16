// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertRecordAnalysisResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlertRecordAnalysisResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlertRecordAnalysisResultResponse
	GetStatusCode() *int32
	SetBody(v *GetAlertRecordAnalysisResultResponseBody) *GetAlertRecordAnalysisResultResponse
	GetBody() *GetAlertRecordAnalysisResultResponseBody
}

type GetAlertRecordAnalysisResultResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlertRecordAnalysisResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlertRecordAnalysisResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRecordAnalysisResultResponse) GoString() string {
	return s.String()
}

func (s *GetAlertRecordAnalysisResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlertRecordAnalysisResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlertRecordAnalysisResultResponse) GetBody() *GetAlertRecordAnalysisResultResponseBody {
	return s.Body
}

func (s *GetAlertRecordAnalysisResultResponse) SetHeaders(v map[string]*string) *GetAlertRecordAnalysisResultResponse {
	s.Headers = v
	return s
}

func (s *GetAlertRecordAnalysisResultResponse) SetStatusCode(v int32) *GetAlertRecordAnalysisResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlertRecordAnalysisResultResponse) SetBody(v *GetAlertRecordAnalysisResultResponseBody) *GetAlertRecordAnalysisResultResponse {
	s.Body = v
	return s
}

func (s *GetAlertRecordAnalysisResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
