// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsoleProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *ConsoleProxyRequest
	GetAppCode() *string
	SetInterfaceName(v string) *ConsoleProxyRequest
	GetInterfaceName() *string
	SetParamJson(v string) *ConsoleProxyRequest
	GetParamJson() *string
	SetRequestId(v string) *ConsoleProxyRequest
	GetRequestId() *string
	SetTeamHashId(v string) *ConsoleProxyRequest
	GetTeamHashId() *string
}

type ConsoleProxyRequest struct {
	AppCode       *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	InterfaceName *string `json:"interfaceName,omitempty" xml:"interfaceName,omitempty"`
	ParamJson     *string `json:"paramJson,omitempty" xml:"paramJson,omitempty"`
	RequestId     *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TeamHashId    *string `json:"teamHashId,omitempty" xml:"teamHashId,omitempty"`
}

func (s ConsoleProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s ConsoleProxyRequest) GoString() string {
	return s.String()
}

func (s *ConsoleProxyRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *ConsoleProxyRequest) GetInterfaceName() *string {
	return s.InterfaceName
}

func (s *ConsoleProxyRequest) GetParamJson() *string {
	return s.ParamJson
}

func (s *ConsoleProxyRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ConsoleProxyRequest) GetTeamHashId() *string {
	return s.TeamHashId
}

func (s *ConsoleProxyRequest) SetAppCode(v string) *ConsoleProxyRequest {
	s.AppCode = &v
	return s
}

func (s *ConsoleProxyRequest) SetInterfaceName(v string) *ConsoleProxyRequest {
	s.InterfaceName = &v
	return s
}

func (s *ConsoleProxyRequest) SetParamJson(v string) *ConsoleProxyRequest {
	s.ParamJson = &v
	return s
}

func (s *ConsoleProxyRequest) SetRequestId(v string) *ConsoleProxyRequest {
	s.RequestId = &v
	return s
}

func (s *ConsoleProxyRequest) SetTeamHashId(v string) *ConsoleProxyRequest {
	s.TeamHashId = &v
	return s
}

func (s *ConsoleProxyRequest) Validate() error {
	return dara.Validate(s)
}
