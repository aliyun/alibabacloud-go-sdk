// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateScriptResponseBody
	GetRequestId() *string
	SetScriptId(v string) *CreateScriptResponseBody
	GetScriptId() *string
}

type CreateScriptResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The script ID.
	//
	// example:
	//
	// cs-d6d6bc841c0d415fb81808bc6d09****
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s CreateScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScriptResponseBody) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateScriptResponseBody) SetRequestId(v string) *CreateScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScriptResponseBody) SetScriptId(v string) *CreateScriptResponseBody {
	s.ScriptId = &v
	return s
}

func (s *CreateScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
