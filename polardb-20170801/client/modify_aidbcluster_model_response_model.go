// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAIDBClusterModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAIDBClusterModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAIDBClusterModelResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAIDBClusterModelResponseBody) *ModifyAIDBClusterModelResponse
	GetBody() *ModifyAIDBClusterModelResponseBody
}

type ModifyAIDBClusterModelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAIDBClusterModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAIDBClusterModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAIDBClusterModelResponse) GoString() string {
	return s.String()
}

func (s *ModifyAIDBClusterModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAIDBClusterModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAIDBClusterModelResponse) GetBody() *ModifyAIDBClusterModelResponseBody {
	return s.Body
}

func (s *ModifyAIDBClusterModelResponse) SetHeaders(v map[string]*string) *ModifyAIDBClusterModelResponse {
	s.Headers = v
	return s
}

func (s *ModifyAIDBClusterModelResponse) SetStatusCode(v int32) *ModifyAIDBClusterModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAIDBClusterModelResponse) SetBody(v *ModifyAIDBClusterModelResponseBody) *ModifyAIDBClusterModelResponse {
	s.Body = v
	return s
}

func (s *ModifyAIDBClusterModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
