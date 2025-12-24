// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescriberPython3ScriptLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescriberPython3ScriptLogsRequest
	GetLang() *string
	SetRequestUuid(v string) *DescriberPython3ScriptLogsRequest
	GetRequestUuid() *string
}

type DescriberPython3ScriptLogsRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID that is returned when the Python3 script is run.
	//
	// >  You can call the [RunPython3Script](~~RunPython3Script~~) operation to query the UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 69edc2b4-c95c-424f-9114-xxxxxxx
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
}

func (s DescriberPython3ScriptLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescriberPython3ScriptLogsRequest) GoString() string {
	return s.String()
}

func (s *DescriberPython3ScriptLogsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescriberPython3ScriptLogsRequest) GetRequestUuid() *string {
	return s.RequestUuid
}

func (s *DescriberPython3ScriptLogsRequest) SetLang(v string) *DescriberPython3ScriptLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescriberPython3ScriptLogsRequest) SetRequestUuid(v string) *DescriberPython3ScriptLogsRequest {
	s.RequestUuid = &v
	return s
}

func (s *DescriberPython3ScriptLogsRequest) Validate() error {
	return dara.Validate(s)
}
