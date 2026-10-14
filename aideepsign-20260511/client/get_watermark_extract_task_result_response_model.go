// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarkExtractTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWatermarkExtractTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWatermarkExtractTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetWatermarkExtractTaskResultResponseBody) *GetWatermarkExtractTaskResultResponse
	GetBody() *GetWatermarkExtractTaskResultResponseBody
}

type GetWatermarkExtractTaskResultResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWatermarkExtractTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWatermarkExtractTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkExtractTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetWatermarkExtractTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWatermarkExtractTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWatermarkExtractTaskResultResponse) GetBody() *GetWatermarkExtractTaskResultResponseBody {
	return s.Body
}

func (s *GetWatermarkExtractTaskResultResponse) SetHeaders(v map[string]*string) *GetWatermarkExtractTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetWatermarkExtractTaskResultResponse) SetStatusCode(v int32) *GetWatermarkExtractTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWatermarkExtractTaskResultResponse) SetBody(v *GetWatermarkExtractTaskResultResponseBody) *GetWatermarkExtractTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetWatermarkExtractTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
