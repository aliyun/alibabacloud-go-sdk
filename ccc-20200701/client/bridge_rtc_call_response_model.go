// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBridgeRtcCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BridgeRtcCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BridgeRtcCallResponse
	GetStatusCode() *int32
	SetBody(v *BridgeRtcCallResponseBody) *BridgeRtcCallResponse
	GetBody() *BridgeRtcCallResponseBody
}

type BridgeRtcCallResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BridgeRtcCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BridgeRtcCallResponse) String() string {
	return dara.Prettify(s)
}

func (s BridgeRtcCallResponse) GoString() string {
	return s.String()
}

func (s *BridgeRtcCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BridgeRtcCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BridgeRtcCallResponse) GetBody() *BridgeRtcCallResponseBody {
	return s.Body
}

func (s *BridgeRtcCallResponse) SetHeaders(v map[string]*string) *BridgeRtcCallResponse {
	s.Headers = v
	return s
}

func (s *BridgeRtcCallResponse) SetStatusCode(v int32) *BridgeRtcCallResponse {
	s.StatusCode = &v
	return s
}

func (s *BridgeRtcCallResponse) SetBody(v *BridgeRtcCallResponseBody) *BridgeRtcCallResponse {
	s.Body = v
	return s
}

func (s *BridgeRtcCallResponse) Validate() error {
	return dara.Validate(s)
}
