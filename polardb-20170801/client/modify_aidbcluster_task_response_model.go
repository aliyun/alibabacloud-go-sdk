// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAIDBClusterTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAIDBClusterTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAIDBClusterTaskResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAIDBClusterTaskResponseBody) *ModifyAIDBClusterTaskResponse
	GetBody() *ModifyAIDBClusterTaskResponseBody
}

type ModifyAIDBClusterTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAIDBClusterTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAIDBClusterTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAIDBClusterTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyAIDBClusterTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAIDBClusterTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAIDBClusterTaskResponse) GetBody() *ModifyAIDBClusterTaskResponseBody {
	return s.Body
}

func (s *ModifyAIDBClusterTaskResponse) SetHeaders(v map[string]*string) *ModifyAIDBClusterTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyAIDBClusterTaskResponse) SetStatusCode(v int32) *ModifyAIDBClusterTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAIDBClusterTaskResponse) SetBody(v *ModifyAIDBClusterTaskResponseBody) *ModifyAIDBClusterTaskResponse {
	s.Body = v
	return s
}

func (s *ModifyAIDBClusterTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
