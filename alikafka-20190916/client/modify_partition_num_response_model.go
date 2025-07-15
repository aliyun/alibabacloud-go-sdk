// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPartitionNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPartitionNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPartitionNumResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPartitionNumResponseBody) *ModifyPartitionNumResponse
	GetBody() *ModifyPartitionNumResponseBody
}

type ModifyPartitionNumResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPartitionNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPartitionNumResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPartitionNumResponse) GoString() string {
	return s.String()
}

func (s *ModifyPartitionNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPartitionNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPartitionNumResponse) GetBody() *ModifyPartitionNumResponseBody {
	return s.Body
}

func (s *ModifyPartitionNumResponse) SetHeaders(v map[string]*string) *ModifyPartitionNumResponse {
	s.Headers = v
	return s
}

func (s *ModifyPartitionNumResponse) SetStatusCode(v int32) *ModifyPartitionNumResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPartitionNumResponse) SetBody(v *ModifyPartitionNumResponseBody) *ModifyPartitionNumResponse {
	s.Body = v
	return s
}

func (s *ModifyPartitionNumResponse) Validate() error {
	return dara.Validate(s)
}
