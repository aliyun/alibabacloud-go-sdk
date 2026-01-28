// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBridgeWebCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BridgeWebCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BridgeWebCallResponse
	GetStatusCode() *int32
	SetBody(v *BridgeWebCallResponseBody) *BridgeWebCallResponse
	GetBody() *BridgeWebCallResponseBody
}

type BridgeWebCallResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BridgeWebCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BridgeWebCallResponse) String() string {
	return dara.Prettify(s)
}

func (s BridgeWebCallResponse) GoString() string {
	return s.String()
}

func (s *BridgeWebCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BridgeWebCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BridgeWebCallResponse) GetBody() *BridgeWebCallResponseBody {
	return s.Body
}

func (s *BridgeWebCallResponse) SetHeaders(v map[string]*string) *BridgeWebCallResponse {
	s.Headers = v
	return s
}

func (s *BridgeWebCallResponse) SetStatusCode(v int32) *BridgeWebCallResponse {
	s.StatusCode = &v
	return s
}

func (s *BridgeWebCallResponse) SetBody(v *BridgeWebCallResponseBody) *BridgeWebCallResponse {
	s.Body = v
	return s
}

func (s *BridgeWebCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
