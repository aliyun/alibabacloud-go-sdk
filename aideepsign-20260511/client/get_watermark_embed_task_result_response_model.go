// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarkEmbedTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWatermarkEmbedTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWatermarkEmbedTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetWatermarkEmbedTaskResultResponseBody) *GetWatermarkEmbedTaskResultResponse
	GetBody() *GetWatermarkEmbedTaskResultResponseBody
}

type GetWatermarkEmbedTaskResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWatermarkEmbedTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWatermarkEmbedTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkEmbedTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetWatermarkEmbedTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWatermarkEmbedTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWatermarkEmbedTaskResultResponse) GetBody() *GetWatermarkEmbedTaskResultResponseBody {
	return s.Body
}

func (s *GetWatermarkEmbedTaskResultResponse) SetHeaders(v map[string]*string) *GetWatermarkEmbedTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetWatermarkEmbedTaskResultResponse) SetStatusCode(v int32) *GetWatermarkEmbedTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWatermarkEmbedTaskResultResponse) SetBody(v *GetWatermarkEmbedTaskResultResponseBody) *GetWatermarkEmbedTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetWatermarkEmbedTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
