// This file is auto-generated, don't edit it. Thanks.
package client

import (
	websocketutils "github.com/alibabacloud-go/darabonba-openapi/v2/websocketUtils"
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardWsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetWebSocketClient(v *websocketutils.WebSocketClient) *MultiModalGuardWsResponse
	GetWebSocketClient() *websocketutils.WebSocketClient
}

type MultiModalGuardWsResponse struct {
	WebSocketClient *websocketutils.WebSocketClient `json:"WebSocketClient,omitempty" xml:"WebSocketClient,omitempty"`
}

func (s MultiModalGuardWsResponse) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardWsResponse) GoString() string {
	return s.String()
}

func (s *MultiModalGuardWsResponse) GetWebSocketClient() *websocketutils.WebSocketClient {
	return s.WebSocketClient
}

func (s *MultiModalGuardWsResponse) SetWebSocketClient(v *websocketutils.WebSocketClient) *MultiModalGuardWsResponse {
	s.WebSocketClient = v
	return s
}

func (s *MultiModalGuardWsResponse) Validate() error {
	if s.WebSocketClient != nil {
		if err := s.WebSocketClient.Validate(); err != nil {
			return err
		}
	}
	return nil
}
