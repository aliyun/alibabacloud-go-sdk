// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeAIDBClusterModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResizeAIDBClusterModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResizeAIDBClusterModelResponse
	GetStatusCode() *int32
	SetBody(v *ResizeAIDBClusterModelResponseBody) *ResizeAIDBClusterModelResponse
	GetBody() *ResizeAIDBClusterModelResponseBody
}

type ResizeAIDBClusterModelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeAIDBClusterModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeAIDBClusterModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ResizeAIDBClusterModelResponse) GoString() string {
	return s.String()
}

func (s *ResizeAIDBClusterModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResizeAIDBClusterModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResizeAIDBClusterModelResponse) GetBody() *ResizeAIDBClusterModelResponseBody {
	return s.Body
}

func (s *ResizeAIDBClusterModelResponse) SetHeaders(v map[string]*string) *ResizeAIDBClusterModelResponse {
	s.Headers = v
	return s
}

func (s *ResizeAIDBClusterModelResponse) SetStatusCode(v int32) *ResizeAIDBClusterModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeAIDBClusterModelResponse) SetBody(v *ResizeAIDBClusterModelResponseBody) *ResizeAIDBClusterModelResponse {
	s.Body = v
	return s
}

func (s *ResizeAIDBClusterModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
