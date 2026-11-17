// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEduHomeworkCorrectionResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEduHomeworkCorrectionResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEduHomeworkCorrectionResultResponse
	GetStatusCode() *int32
	SetBody(v *GetEduHomeworkCorrectionResultResponseBody) *GetEduHomeworkCorrectionResultResponse
	GetBody() *GetEduHomeworkCorrectionResultResponseBody
}

type GetEduHomeworkCorrectionResultResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEduHomeworkCorrectionResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEduHomeworkCorrectionResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEduHomeworkCorrectionResultResponse) GoString() string {
	return s.String()
}

func (s *GetEduHomeworkCorrectionResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEduHomeworkCorrectionResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEduHomeworkCorrectionResultResponse) GetBody() *GetEduHomeworkCorrectionResultResponseBody {
	return s.Body
}

func (s *GetEduHomeworkCorrectionResultResponse) SetHeaders(v map[string]*string) *GetEduHomeworkCorrectionResultResponse {
	s.Headers = v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponse) SetStatusCode(v int32) *GetEduHomeworkCorrectionResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponse) SetBody(v *GetEduHomeworkCorrectionResultResponseBody) *GetEduHomeworkCorrectionResultResponse {
	s.Body = v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
