// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWatermarkEmbedTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWatermarkEmbedTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWatermarkEmbedTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateWatermarkEmbedTaskResponseBody) *CreateWatermarkEmbedTaskResponse
	GetBody() *CreateWatermarkEmbedTaskResponseBody
}

type CreateWatermarkEmbedTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWatermarkEmbedTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWatermarkEmbedTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWatermarkEmbedTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateWatermarkEmbedTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWatermarkEmbedTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWatermarkEmbedTaskResponse) GetBody() *CreateWatermarkEmbedTaskResponseBody {
	return s.Body
}

func (s *CreateWatermarkEmbedTaskResponse) SetHeaders(v map[string]*string) *CreateWatermarkEmbedTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateWatermarkEmbedTaskResponse) SetStatusCode(v int32) *CreateWatermarkEmbedTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWatermarkEmbedTaskResponse) SetBody(v *CreateWatermarkEmbedTaskResponseBody) *CreateWatermarkEmbedTaskResponse {
	s.Body = v
	return s
}

func (s *CreateWatermarkEmbedTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
