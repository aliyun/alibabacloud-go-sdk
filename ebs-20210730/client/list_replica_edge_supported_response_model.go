// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReplicaEdgeSupportedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListReplicaEdgeSupportedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListReplicaEdgeSupportedResponse
	GetStatusCode() *int32
	SetBody(v *ListReplicaEdgeSupportedResponseBody) *ListReplicaEdgeSupportedResponse
	GetBody() *ListReplicaEdgeSupportedResponseBody
}

type ListReplicaEdgeSupportedResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReplicaEdgeSupportedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReplicaEdgeSupportedResponse) String() string {
	return dara.Prettify(s)
}

func (s ListReplicaEdgeSupportedResponse) GoString() string {
	return s.String()
}

func (s *ListReplicaEdgeSupportedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListReplicaEdgeSupportedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListReplicaEdgeSupportedResponse) GetBody() *ListReplicaEdgeSupportedResponseBody {
	return s.Body
}

func (s *ListReplicaEdgeSupportedResponse) SetHeaders(v map[string]*string) *ListReplicaEdgeSupportedResponse {
	s.Headers = v
	return s
}

func (s *ListReplicaEdgeSupportedResponse) SetStatusCode(v int32) *ListReplicaEdgeSupportedResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReplicaEdgeSupportedResponse) SetBody(v *ListReplicaEdgeSupportedResponseBody) *ListReplicaEdgeSupportedResponse {
	s.Body = v
	return s
}

func (s *ListReplicaEdgeSupportedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
