// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWatermarkExtractTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWatermarkExtractTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWatermarkExtractTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateWatermarkExtractTaskResponseBody) *CreateWatermarkExtractTaskResponse
	GetBody() *CreateWatermarkExtractTaskResponseBody
}

type CreateWatermarkExtractTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWatermarkExtractTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWatermarkExtractTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWatermarkExtractTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateWatermarkExtractTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWatermarkExtractTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWatermarkExtractTaskResponse) GetBody() *CreateWatermarkExtractTaskResponseBody {
	return s.Body
}

func (s *CreateWatermarkExtractTaskResponse) SetHeaders(v map[string]*string) *CreateWatermarkExtractTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateWatermarkExtractTaskResponse) SetStatusCode(v int32) *CreateWatermarkExtractTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWatermarkExtractTaskResponse) SetBody(v *CreateWatermarkExtractTaskResponseBody) *CreateWatermarkExtractTaskResponse {
	s.Body = v
	return s
}

func (s *CreateWatermarkExtractTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
