// This file is auto-generated, don't edit it. Thanks.
package client

import (
	websocketutils "github.com/alibabacloud-go/darabonba-openapi/v2/websocketUtils"
	"github.com/alibabacloud-go/tea/dara"
)

type iRealtimeSpeechTranslationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetWebSocketClient(v *websocketutils.WebSocketClient) *RealtimeSpeechTranslationResponse
	GetWebSocketClient() *websocketutils.WebSocketClient
}

type RealtimeSpeechTranslationResponse struct {
	WebSocketClient *websocketutils.WebSocketClient `json:"webSocketClient,omitempty" xml:"webSocketClient,omitempty"`
}

func (s RealtimeSpeechTranslationResponse) String() string {
	return dara.Prettify(s)
}

func (s RealtimeSpeechTranslationResponse) GoString() string {
	return s.String()
}

func (s *RealtimeSpeechTranslationResponse) GetWebSocketClient() *websocketutils.WebSocketClient {
	return s.WebSocketClient
}

func (s *RealtimeSpeechTranslationResponse) SetWebSocketClient(v *websocketutils.WebSocketClient) *RealtimeSpeechTranslationResponse {
	s.WebSocketClient = v
	return s
}

func (s *RealtimeSpeechTranslationResponse) Validate() error {
	if s.WebSocketClient != nil {
		if err := s.WebSocketClient.Validate(); err != nil {
			return err
		}
	}
	return nil
}
