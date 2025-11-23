// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStructSyncJobAnalyzeResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStructSyncJobAnalyzeResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStructSyncJobAnalyzeResultResponse
	GetStatusCode() *int32
	SetBody(v *GetStructSyncJobAnalyzeResultResponseBody) *GetStructSyncJobAnalyzeResultResponse
	GetBody() *GetStructSyncJobAnalyzeResultResponseBody
}

type GetStructSyncJobAnalyzeResultResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStructSyncJobAnalyzeResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStructSyncJobAnalyzeResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncJobAnalyzeResultResponse) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobAnalyzeResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStructSyncJobAnalyzeResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStructSyncJobAnalyzeResultResponse) GetBody() *GetStructSyncJobAnalyzeResultResponseBody {
	return s.Body
}

func (s *GetStructSyncJobAnalyzeResultResponse) SetHeaders(v map[string]*string) *GetStructSyncJobAnalyzeResultResponse {
	s.Headers = v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponse) SetStatusCode(v int32) *GetStructSyncJobAnalyzeResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponse) SetBody(v *GetStructSyncJobAnalyzeResultResponseBody) *GetStructSyncJobAnalyzeResultResponse {
	s.Body = v
	return s
}

func (s *GetStructSyncJobAnalyzeResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
