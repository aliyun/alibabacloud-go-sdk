// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProbeHandler interface {
	dara.Model
	String() string
	GoString() string
	SetHttpGet(v *HTTPGetAction) *ProbeHandler
	GetHttpGet() *HTTPGetAction
	SetTcpSocket(v *TCPSocketAction) *ProbeHandler
	GetTcpSocket() *TCPSocketAction
}

type ProbeHandler struct {
	HttpGet   *HTTPGetAction   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty"`
	TcpSocket *TCPSocketAction `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty"`
}

func (s ProbeHandler) String() string {
	return dara.Prettify(s)
}

func (s ProbeHandler) GoString() string {
	return s.String()
}

func (s *ProbeHandler) GetHttpGet() *HTTPGetAction {
	return s.HttpGet
}

func (s *ProbeHandler) GetTcpSocket() *TCPSocketAction {
	return s.TcpSocket
}

func (s *ProbeHandler) SetHttpGet(v *HTTPGetAction) *ProbeHandler {
	s.HttpGet = v
	return s
}

func (s *ProbeHandler) SetTcpSocket(v *TCPSocketAction) *ProbeHandler {
	s.TcpSocket = v
	return s
}

func (s *ProbeHandler) Validate() error {
	return dara.Validate(s)
}
