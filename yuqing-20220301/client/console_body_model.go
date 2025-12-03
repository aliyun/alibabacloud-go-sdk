// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsoleBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *ConsoleBody
	GetAppCode() *string
	SetInterfaceName(v string) *ConsoleBody
	GetInterfaceName() *string
	SetParamJson(v string) *ConsoleBody
	GetParamJson() *string
	SetRequestId(v string) *ConsoleBody
	GetRequestId() *string
	SetTeamHashId(v string) *ConsoleBody
	GetTeamHashId() *string
}

type ConsoleBody struct {
	AppCode       *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	InterfaceName *string `json:"interfaceName,omitempty" xml:"interfaceName,omitempty"`
	ParamJson     *string `json:"paramJson,omitempty" xml:"paramJson,omitempty"`
	RequestId     *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TeamHashId    *string `json:"teamHashId,omitempty" xml:"teamHashId,omitempty"`
}

func (s ConsoleBody) String() string {
	return dara.Prettify(s)
}

func (s ConsoleBody) GoString() string {
	return s.String()
}

func (s *ConsoleBody) GetAppCode() *string {
	return s.AppCode
}

func (s *ConsoleBody) GetInterfaceName() *string {
	return s.InterfaceName
}

func (s *ConsoleBody) GetParamJson() *string {
	return s.ParamJson
}

func (s *ConsoleBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConsoleBody) GetTeamHashId() *string {
	return s.TeamHashId
}

func (s *ConsoleBody) SetAppCode(v string) *ConsoleBody {
	s.AppCode = &v
	return s
}

func (s *ConsoleBody) SetInterfaceName(v string) *ConsoleBody {
	s.InterfaceName = &v
	return s
}

func (s *ConsoleBody) SetParamJson(v string) *ConsoleBody {
	s.ParamJson = &v
	return s
}

func (s *ConsoleBody) SetRequestId(v string) *ConsoleBody {
	s.RequestId = &v
	return s
}

func (s *ConsoleBody) SetTeamHashId(v string) *ConsoleBody {
	s.TeamHashId = &v
	return s
}

func (s *ConsoleBody) Validate() error {
	return dara.Validate(s)
}
