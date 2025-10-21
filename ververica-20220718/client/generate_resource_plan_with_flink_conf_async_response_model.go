// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateResourcePlanWithFlinkConfAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateResourcePlanWithFlinkConfAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateResourcePlanWithFlinkConfAsyncResponse
	GetStatusCode() *int32
	SetBody(v *GenerateResourcePlanWithFlinkConfAsyncResponseBody) *GenerateResourcePlanWithFlinkConfAsyncResponse
	GetBody() *GenerateResourcePlanWithFlinkConfAsyncResponseBody
}

type GenerateResourcePlanWithFlinkConfAsyncResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateResourcePlanWithFlinkConfAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateResourcePlanWithFlinkConfAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateResourcePlanWithFlinkConfAsyncResponse) GoString() string {
	return s.String()
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponse) GetBody() *GenerateResourcePlanWithFlinkConfAsyncResponseBody {
	return s.Body
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponse) SetHeaders(v map[string]*string) *GenerateResourcePlanWithFlinkConfAsyncResponse {
	s.Headers = v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponse) SetStatusCode(v int32) *GenerateResourcePlanWithFlinkConfAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponse) SetBody(v *GenerateResourcePlanWithFlinkConfAsyncResponseBody) *GenerateResourcePlanWithFlinkConfAsyncResponse {
	s.Body = v
	return s
}

func (s *GenerateResourcePlanWithFlinkConfAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
