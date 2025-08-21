// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransferableNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTransferableNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTransferableNodesResponse
	GetStatusCode() *int32
	SetBody(v *GetTransferableNodesResponseBody) *GetTransferableNodesResponse
	GetBody() *GetTransferableNodesResponseBody
}

type GetTransferableNodesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTransferableNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTransferableNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTransferableNodesResponse) GoString() string {
	return s.String()
}

func (s *GetTransferableNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTransferableNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTransferableNodesResponse) GetBody() *GetTransferableNodesResponseBody {
	return s.Body
}

func (s *GetTransferableNodesResponse) SetHeaders(v map[string]*string) *GetTransferableNodesResponse {
	s.Headers = v
	return s
}

func (s *GetTransferableNodesResponse) SetStatusCode(v int32) *GetTransferableNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTransferableNodesResponse) SetBody(v *GetTransferableNodesResponseBody) *GetTransferableNodesResponse {
	s.Body = v
	return s
}

func (s *GetTransferableNodesResponse) Validate() error {
	return dara.Validate(s)
}
