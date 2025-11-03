// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeSpecBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNodeSpecBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNodeSpecBatchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNodeSpecBatchResponseBody) *ModifyNodeSpecBatchResponse
	GetBody() *ModifyNodeSpecBatchResponseBody
}

type ModifyNodeSpecBatchResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNodeSpecBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNodeSpecBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeSpecBatchResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodeSpecBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNodeSpecBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNodeSpecBatchResponse) GetBody() *ModifyNodeSpecBatchResponseBody {
	return s.Body
}

func (s *ModifyNodeSpecBatchResponse) SetHeaders(v map[string]*string) *ModifyNodeSpecBatchResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodeSpecBatchResponse) SetStatusCode(v int32) *ModifyNodeSpecBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodeSpecBatchResponse) SetBody(v *ModifyNodeSpecBatchResponseBody) *ModifyNodeSpecBatchResponse {
	s.Body = v
	return s
}

func (s *ModifyNodeSpecBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
