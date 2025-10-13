// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProbe interface {
	dara.Model
	String() string
	GoString() string
	SetFailureThreshold(v int32) *Probe
	GetFailureThreshold() *int32
	SetInitialDelaySeconds(v int32) *Probe
	GetInitialDelaySeconds() *int32
	SetPeriodSeconds(v int32) *Probe
	GetPeriodSeconds() *int32
	SetProbeHandler(v *ProbeProbeHandler) *Probe
	GetProbeHandler() *ProbeProbeHandler
	SetTimeoutSeconds(v int32) *Probe
	GetTimeoutSeconds() *int32
}

type Probe struct {
	FailureThreshold    *int32             `json:"failureThreshold,omitempty" xml:"failureThreshold,omitempty"`
	InitialDelaySeconds *int32             `json:"initialDelaySeconds,omitempty" xml:"initialDelaySeconds,omitempty"`
	PeriodSeconds       *int32             `json:"periodSeconds,omitempty" xml:"periodSeconds,omitempty"`
	ProbeHandler        *ProbeProbeHandler `json:"probeHandler,omitempty" xml:"probeHandler,omitempty" type:"Struct"`
	TimeoutSeconds      *int32             `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s Probe) String() string {
	return dara.Prettify(s)
}

func (s Probe) GoString() string {
	return s.String()
}

func (s *Probe) GetFailureThreshold() *int32 {
	return s.FailureThreshold
}

func (s *Probe) GetInitialDelaySeconds() *int32 {
	return s.InitialDelaySeconds
}

func (s *Probe) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *Probe) GetProbeHandler() *ProbeProbeHandler {
	return s.ProbeHandler
}

func (s *Probe) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *Probe) SetFailureThreshold(v int32) *Probe {
	s.FailureThreshold = &v
	return s
}

func (s *Probe) SetInitialDelaySeconds(v int32) *Probe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *Probe) SetPeriodSeconds(v int32) *Probe {
	s.PeriodSeconds = &v
	return s
}

func (s *Probe) SetProbeHandler(v *ProbeProbeHandler) *Probe {
	s.ProbeHandler = v
	return s
}

func (s *Probe) SetTimeoutSeconds(v int32) *Probe {
	s.TimeoutSeconds = &v
	return s
}

func (s *Probe) Validate() error {
	if s.ProbeHandler != nil {
		if err := s.ProbeHandler.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProbeProbeHandler struct {
	HttpGet   *ProbeProbeHandlerHttpGet   `json:"httpGet,omitempty" xml:"httpGet,omitempty" type:"Struct"`
	TcpSocket *ProbeProbeHandlerTcpSocket `json:"tcpSocket,omitempty" xml:"tcpSocket,omitempty" type:"Struct"`
}

func (s ProbeProbeHandler) String() string {
	return dara.Prettify(s)
}

func (s ProbeProbeHandler) GoString() string {
	return s.String()
}

func (s *ProbeProbeHandler) GetHttpGet() *ProbeProbeHandlerHttpGet {
	return s.HttpGet
}

func (s *ProbeProbeHandler) GetTcpSocket() *ProbeProbeHandlerTcpSocket {
	return s.TcpSocket
}

func (s *ProbeProbeHandler) SetHttpGet(v *ProbeProbeHandlerHttpGet) *ProbeProbeHandler {
	s.HttpGet = v
	return s
}

func (s *ProbeProbeHandler) SetTcpSocket(v *ProbeProbeHandlerTcpSocket) *ProbeProbeHandler {
	s.TcpSocket = v
	return s
}

func (s *ProbeProbeHandler) Validate() error {
	if s.HttpGet != nil {
		if err := s.HttpGet.Validate(); err != nil {
			return err
		}
	}
	if s.TcpSocket != nil {
		if err := s.TcpSocket.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProbeProbeHandlerHttpGet struct {
	HttpHeaders []*ProbeProbeHandlerHttpGetHttpHeaders `json:"httpHeaders,omitempty" xml:"httpHeaders,omitempty" type:"Repeated"`
	Path        *string                                `json:"path,omitempty" xml:"path,omitempty"`
	Port        *int32                                 `json:"port,omitempty" xml:"port,omitempty"`
}

func (s ProbeProbeHandlerHttpGet) String() string {
	return dara.Prettify(s)
}

func (s ProbeProbeHandlerHttpGet) GoString() string {
	return s.String()
}

func (s *ProbeProbeHandlerHttpGet) GetHttpHeaders() []*ProbeProbeHandlerHttpGetHttpHeaders {
	return s.HttpHeaders
}

func (s *ProbeProbeHandlerHttpGet) GetPath() *string {
	return s.Path
}

func (s *ProbeProbeHandlerHttpGet) GetPort() *int32 {
	return s.Port
}

func (s *ProbeProbeHandlerHttpGet) SetHttpHeaders(v []*ProbeProbeHandlerHttpGetHttpHeaders) *ProbeProbeHandlerHttpGet {
	s.HttpHeaders = v
	return s
}

func (s *ProbeProbeHandlerHttpGet) SetPath(v string) *ProbeProbeHandlerHttpGet {
	s.Path = &v
	return s
}

func (s *ProbeProbeHandlerHttpGet) SetPort(v int32) *ProbeProbeHandlerHttpGet {
	s.Port = &v
	return s
}

func (s *ProbeProbeHandlerHttpGet) Validate() error {
	if s.HttpHeaders != nil {
		for _, item := range s.HttpHeaders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ProbeProbeHandlerHttpGetHttpHeaders struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ProbeProbeHandlerHttpGetHttpHeaders) String() string {
	return dara.Prettify(s)
}

func (s ProbeProbeHandlerHttpGetHttpHeaders) GoString() string {
	return s.String()
}

func (s *ProbeProbeHandlerHttpGetHttpHeaders) GetName() *string {
	return s.Name
}

func (s *ProbeProbeHandlerHttpGetHttpHeaders) GetValue() *string {
	return s.Value
}

func (s *ProbeProbeHandlerHttpGetHttpHeaders) SetName(v string) *ProbeProbeHandlerHttpGetHttpHeaders {
	s.Name = &v
	return s
}

func (s *ProbeProbeHandlerHttpGetHttpHeaders) SetValue(v string) *ProbeProbeHandlerHttpGetHttpHeaders {
	s.Value = &v
	return s
}

func (s *ProbeProbeHandlerHttpGetHttpHeaders) Validate() error {
	return dara.Validate(s)
}

type ProbeProbeHandlerTcpSocket struct {
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
}

func (s ProbeProbeHandlerTcpSocket) String() string {
	return dara.Prettify(s)
}

func (s ProbeProbeHandlerTcpSocket) GoString() string {
	return s.String()
}

func (s *ProbeProbeHandlerTcpSocket) GetPort() *int32 {
	return s.Port
}

func (s *ProbeProbeHandlerTcpSocket) SetPort(v int32) *ProbeProbeHandlerTcpSocket {
	s.Port = &v
	return s
}

func (s *ProbeProbeHandlerTcpSocket) Validate() error {
	return dara.Validate(s)
}
