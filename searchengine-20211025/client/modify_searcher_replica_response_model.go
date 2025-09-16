// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySearcherReplicaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySearcherReplicaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySearcherReplicaResponse
	GetStatusCode() *int32
	SetBody(v *ModifySearcherReplicaResponseBody) *ModifySearcherReplicaResponse
	GetBody() *ModifySearcherReplicaResponseBody
}

type ModifySearcherReplicaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySearcherReplicaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySearcherReplicaResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySearcherReplicaResponse) GoString() string {
	return s.String()
}

func (s *ModifySearcherReplicaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySearcherReplicaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySearcherReplicaResponse) GetBody() *ModifySearcherReplicaResponseBody {
	return s.Body
}

func (s *ModifySearcherReplicaResponse) SetHeaders(v map[string]*string) *ModifySearcherReplicaResponse {
	s.Headers = v
	return s
}

func (s *ModifySearcherReplicaResponse) SetStatusCode(v int32) *ModifySearcherReplicaResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySearcherReplicaResponse) SetBody(v *ModifySearcherReplicaResponseBody) *ModifySearcherReplicaResponse {
	s.Body = v
	return s
}

func (s *ModifySearcherReplicaResponse) Validate() error {
	return dara.Validate(s)
}
