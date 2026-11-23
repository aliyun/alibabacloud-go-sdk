// This file is auto-generated, don't edit it. Thanks.
package client

import (
	websocketutils "github.com/alibabacloud-go/darabonba-openapi/v2/websocketUtils"
	"github.com/alibabacloud-go/tea/dara"
)

type iStreamVoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetWebSocketClient(v *websocketutils.WebSocketClient) *StreamVoiceResponse
	GetWebSocketClient() *websocketutils.WebSocketClient
}

type StreamVoiceResponse struct {
	WebSocketClient *websocketutils.WebSocketClient `json:"webSocketClient,omitempty" xml:"webSocketClient,omitempty"`
}

func (s StreamVoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s StreamVoiceResponse) GoString() string {
	return s.String()
}

func (s *StreamVoiceResponse) GetWebSocketClient() *websocketutils.WebSocketClient {
	return s.WebSocketClient
}

func (s *StreamVoiceResponse) SetWebSocketClient(v *websocketutils.WebSocketClient) *StreamVoiceResponse {
	s.WebSocketClient = v
	return s
}

func (s *StreamVoiceResponse) Validate() error {
	if s.WebSocketClient != nil {
		if err := s.WebSocketClient.Validate(); err != nil {
			return err
		}
	}
	return nil
}
