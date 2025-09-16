// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIndexPartitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIndexPartitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIndexPartitionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIndexPartitionResponseBody) *ModifyIndexPartitionResponse
	GetBody() *ModifyIndexPartitionResponseBody
}

type ModifyIndexPartitionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIndexPartitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIndexPartitionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexPartitionResponse) GoString() string {
	return s.String()
}

func (s *ModifyIndexPartitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIndexPartitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIndexPartitionResponse) GetBody() *ModifyIndexPartitionResponseBody {
	return s.Body
}

func (s *ModifyIndexPartitionResponse) SetHeaders(v map[string]*string) *ModifyIndexPartitionResponse {
	s.Headers = v
	return s
}

func (s *ModifyIndexPartitionResponse) SetStatusCode(v int32) *ModifyIndexPartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIndexPartitionResponse) SetBody(v *ModifyIndexPartitionResponseBody) *ModifyIndexPartitionResponse {
	s.Body = v
	return s
}

func (s *ModifyIndexPartitionResponse) Validate() error {
	return dara.Validate(s)
}
