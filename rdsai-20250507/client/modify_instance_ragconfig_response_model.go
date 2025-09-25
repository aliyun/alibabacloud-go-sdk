// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceRAGConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceRAGConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceRAGConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceRAGConfigResponseBody) *ModifyInstanceRAGConfigResponse
	GetBody() *ModifyInstanceRAGConfigResponseBody
}

type ModifyInstanceRAGConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceRAGConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceRAGConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRAGConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRAGConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceRAGConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceRAGConfigResponse) GetBody() *ModifyInstanceRAGConfigResponseBody {
	return s.Body
}

func (s *ModifyInstanceRAGConfigResponse) SetHeaders(v map[string]*string) *ModifyInstanceRAGConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceRAGConfigResponse) SetStatusCode(v int32) *ModifyInstanceRAGConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceRAGConfigResponse) SetBody(v *ModifyInstanceRAGConfigResponseBody) *ModifyInstanceRAGConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceRAGConfigResponse) Validate() error {
	return dara.Validate(s)
}
