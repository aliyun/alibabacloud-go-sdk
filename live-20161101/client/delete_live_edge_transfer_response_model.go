// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveEdgeTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveEdgeTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveEdgeTransferResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveEdgeTransferResponseBody) *DeleteLiveEdgeTransferResponse
	GetBody() *DeleteLiveEdgeTransferResponseBody
}

type DeleteLiveEdgeTransferResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveEdgeTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveEdgeTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveEdgeTransferResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveEdgeTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveEdgeTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveEdgeTransferResponse) GetBody() *DeleteLiveEdgeTransferResponseBody {
	return s.Body
}

func (s *DeleteLiveEdgeTransferResponse) SetHeaders(v map[string]*string) *DeleteLiveEdgeTransferResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveEdgeTransferResponse) SetStatusCode(v int32) *DeleteLiveEdgeTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveEdgeTransferResponse) SetBody(v *DeleteLiveEdgeTransferResponseBody) *DeleteLiveEdgeTransferResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveEdgeTransferResponse) Validate() error {
	return dara.Validate(s)
}
