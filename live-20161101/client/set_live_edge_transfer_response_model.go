// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveEdgeTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLiveEdgeTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLiveEdgeTransferResponse
	GetStatusCode() *int32
	SetBody(v *SetLiveEdgeTransferResponseBody) *SetLiveEdgeTransferResponse
	GetBody() *SetLiveEdgeTransferResponseBody
}

type SetLiveEdgeTransferResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLiveEdgeTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLiveEdgeTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLiveEdgeTransferResponse) GoString() string {
	return s.String()
}

func (s *SetLiveEdgeTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLiveEdgeTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLiveEdgeTransferResponse) GetBody() *SetLiveEdgeTransferResponseBody {
	return s.Body
}

func (s *SetLiveEdgeTransferResponse) SetHeaders(v map[string]*string) *SetLiveEdgeTransferResponse {
	s.Headers = v
	return s
}

func (s *SetLiveEdgeTransferResponse) SetStatusCode(v int32) *SetLiveEdgeTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLiveEdgeTransferResponse) SetBody(v *SetLiveEdgeTransferResponseBody) *SetLiveEdgeTransferResponse {
	s.Body = v
	return s
}

func (s *SetLiveEdgeTransferResponse) Validate() error {
	return dara.Validate(s)
}
