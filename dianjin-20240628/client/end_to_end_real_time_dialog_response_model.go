// This file is auto-generated, don't edit it. Thanks.
package client

import (
  websocketutils "github.com/alibabacloud-go/darabonba-openapi/v2/websocketUtils"
  "github.com/alibabacloud-go/tea/dara"
)

type iEndToEndRealTimeDialogResponse interface {
  dara.Model
  String() string
  GoString() string
  SetWebSocketClient(v *websocketutils.WebSocketClient) *EndToEndRealTimeDialogResponse
  GetWebSocketClient() *websocketutils.WebSocketClient 
}

type EndToEndRealTimeDialogResponse struct {
  WebSocketClient *websocketutils.WebSocketClient `json:"webSocketClient,omitempty" xml:"webSocketClient,omitempty"`
}

func (s EndToEndRealTimeDialogResponse) String() string {
  return dara.Prettify(s)
}

func (s EndToEndRealTimeDialogResponse) GoString() string {
  return s.String()
}

func (s *EndToEndRealTimeDialogResponse) GetWebSocketClient() *websocketutils.WebSocketClient  {
  return s.WebSocketClient
}

func (s *EndToEndRealTimeDialogResponse) SetWebSocketClient(v *websocketutils.WebSocketClient) *EndToEndRealTimeDialogResponse {
  s.WebSocketClient = v
  return s
}

func (s *EndToEndRealTimeDialogResponse) Validate() error {
  if s.WebSocketClient != nil {
    if err := s.WebSocketClient.Validate(); err != nil {
      return err
    }
  }
  return nil
}

