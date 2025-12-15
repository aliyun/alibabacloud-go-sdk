// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachNodesResponse
	GetStatusCode() *int32
	SetBody(v *AttachNodesResponseBody) *AttachNodesResponse
	GetBody() *AttachNodesResponseBody
}

type AttachNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachNodesResponse) GoString() string {
	return s.String()
}

func (s *AttachNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachNodesResponse) GetBody() *AttachNodesResponseBody {
	return s.Body
}

func (s *AttachNodesResponse) SetHeaders(v map[string]*string) *AttachNodesResponse {
	s.Headers = v
	return s
}

func (s *AttachNodesResponse) SetStatusCode(v int32) *AttachNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachNodesResponse) SetBody(v *AttachNodesResponseBody) *AttachNodesResponse {
	s.Body = v
	return s
}

func (s *AttachNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
