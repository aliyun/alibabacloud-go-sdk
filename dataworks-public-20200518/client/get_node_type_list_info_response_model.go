// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeTypeListInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeTypeListInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeTypeListInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeTypeListInfoResponseBody) *GetNodeTypeListInfoResponse
	GetBody() *GetNodeTypeListInfoResponseBody
}

type GetNodeTypeListInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeTypeListInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeTypeListInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeTypeListInfoResponse) GoString() string {
	return s.String()
}

func (s *GetNodeTypeListInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeTypeListInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeTypeListInfoResponse) GetBody() *GetNodeTypeListInfoResponseBody {
	return s.Body
}

func (s *GetNodeTypeListInfoResponse) SetHeaders(v map[string]*string) *GetNodeTypeListInfoResponse {
	s.Headers = v
	return s
}

func (s *GetNodeTypeListInfoResponse) SetStatusCode(v int32) *GetNodeTypeListInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeTypeListInfoResponse) SetBody(v *GetNodeTypeListInfoResponseBody) *GetNodeTypeListInfoResponse {
	s.Body = v
	return s
}

func (s *GetNodeTypeListInfoResponse) Validate() error {
	return dara.Validate(s)
}
