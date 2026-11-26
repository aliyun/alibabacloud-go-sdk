// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddScriptHotWordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHotWordId(v string) *AddScriptHotWordResponseBody
	GetHotWordId() *string
	SetRequestId(v string) *AddScriptHotWordResponseBody
	GetRequestId() *string
}

type AddScriptHotWordResponseBody struct {
	HotWordId *string `json:"hotWordId,omitempty" xml:"hotWordId,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddScriptHotWordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddScriptHotWordResponseBody) GoString() string {
	return s.String()
}

func (s *AddScriptHotWordResponseBody) GetHotWordId() *string {
	return s.HotWordId
}

func (s *AddScriptHotWordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddScriptHotWordResponseBody) SetHotWordId(v string) *AddScriptHotWordResponseBody {
	s.HotWordId = &v
	return s
}

func (s *AddScriptHotWordResponseBody) SetRequestId(v string) *AddScriptHotWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddScriptHotWordResponseBody) Validate() error {
	return dara.Validate(s)
}
