// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageAnalysisResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStorageAnalysisResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStorageAnalysisResultResponse
	GetStatusCode() *int32
	SetBody(v *GetStorageAnalysisResultResponseBody) *GetStorageAnalysisResultResponse
	GetBody() *GetStorageAnalysisResultResponseBody
}

type GetStorageAnalysisResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStorageAnalysisResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStorageAnalysisResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStorageAnalysisResultResponse) GoString() string {
	return s.String()
}

func (s *GetStorageAnalysisResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStorageAnalysisResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStorageAnalysisResultResponse) GetBody() *GetStorageAnalysisResultResponseBody {
	return s.Body
}

func (s *GetStorageAnalysisResultResponse) SetHeaders(v map[string]*string) *GetStorageAnalysisResultResponse {
	s.Headers = v
	return s
}

func (s *GetStorageAnalysisResultResponse) SetStatusCode(v int32) *GetStorageAnalysisResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStorageAnalysisResultResponse) SetBody(v *GetStorageAnalysisResultResponseBody) *GetStorageAnalysisResultResponse {
	s.Body = v
	return s
}

func (s *GetStorageAnalysisResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
