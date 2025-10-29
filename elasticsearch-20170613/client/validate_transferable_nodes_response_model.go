// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateTransferableNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateTransferableNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateTransferableNodesResponse
	GetStatusCode() *int32
	SetBody(v *ValidateTransferableNodesResponseBody) *ValidateTransferableNodesResponse
	GetBody() *ValidateTransferableNodesResponseBody
}

type ValidateTransferableNodesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateTransferableNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateTransferableNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateTransferableNodesResponse) GoString() string {
	return s.String()
}

func (s *ValidateTransferableNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateTransferableNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateTransferableNodesResponse) GetBody() *ValidateTransferableNodesResponseBody {
	return s.Body
}

func (s *ValidateTransferableNodesResponse) SetHeaders(v map[string]*string) *ValidateTransferableNodesResponse {
	s.Headers = v
	return s
}

func (s *ValidateTransferableNodesResponse) SetStatusCode(v int32) *ValidateTransferableNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateTransferableNodesResponse) SetBody(v *ValidateTransferableNodesResponseBody) *ValidateTransferableNodesResponse {
	s.Body = v
	return s
}

func (s *ValidateTransferableNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
