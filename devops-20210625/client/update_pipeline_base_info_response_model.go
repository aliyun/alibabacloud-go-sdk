// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineBaseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePipelineBaseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePipelineBaseInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePipelineBaseInfoResponseBody) *UpdatePipelineBaseInfoResponse
	GetBody() *UpdatePipelineBaseInfoResponseBody
}

type UpdatePipelineBaseInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePipelineBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePipelineBaseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelineBaseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePipelineBaseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePipelineBaseInfoResponse) GetBody() *UpdatePipelineBaseInfoResponseBody {
	return s.Body
}

func (s *UpdatePipelineBaseInfoResponse) SetHeaders(v map[string]*string) *UpdatePipelineBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelineBaseInfoResponse) SetStatusCode(v int32) *UpdatePipelineBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePipelineBaseInfoResponse) SetBody(v *UpdatePipelineBaseInfoResponseBody) *UpdatePipelineBaseInfoResponse {
	s.Body = v
	return s
}

func (s *UpdatePipelineBaseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
