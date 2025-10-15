// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSummaryTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSummaryTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSummaryTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetSummaryTaskResultResponseBody) *GetSummaryTaskResultResponse
	GetBody() *GetSummaryTaskResultResponseBody
}

type GetSummaryTaskResultResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSummaryTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSummaryTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSummaryTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSummaryTaskResultResponse) GetBody() *GetSummaryTaskResultResponseBody {
	return s.Body
}

func (s *GetSummaryTaskResultResponse) SetHeaders(v map[string]*string) *GetSummaryTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetSummaryTaskResultResponse) SetStatusCode(v int32) *GetSummaryTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSummaryTaskResultResponse) SetBody(v *GetSummaryTaskResultResponseBody) *GetSummaryTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetSummaryTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
