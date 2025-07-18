// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicDisposalProcessesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDynamicDisposalProcessesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDynamicDisposalProcessesResponse
	GetStatusCode() *int32
	SetBody(v *ListDynamicDisposalProcessesResponseBody) *ListDynamicDisposalProcessesResponse
	GetBody() *ListDynamicDisposalProcessesResponseBody
}

type ListDynamicDisposalProcessesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDynamicDisposalProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDynamicDisposalProcessesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicDisposalProcessesResponse) GoString() string {
	return s.String()
}

func (s *ListDynamicDisposalProcessesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDynamicDisposalProcessesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDynamicDisposalProcessesResponse) GetBody() *ListDynamicDisposalProcessesResponseBody {
	return s.Body
}

func (s *ListDynamicDisposalProcessesResponse) SetHeaders(v map[string]*string) *ListDynamicDisposalProcessesResponse {
	s.Headers = v
	return s
}

func (s *ListDynamicDisposalProcessesResponse) SetStatusCode(v int32) *ListDynamicDisposalProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDynamicDisposalProcessesResponse) SetBody(v *ListDynamicDisposalProcessesResponseBody) *ListDynamicDisposalProcessesResponse {
	s.Body = v
	return s
}

func (s *ListDynamicDisposalProcessesResponse) Validate() error {
	return dara.Validate(s)
}
