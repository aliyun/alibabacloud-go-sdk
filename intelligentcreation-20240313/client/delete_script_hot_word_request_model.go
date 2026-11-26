// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptHotWordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotWordId(v string) *DeleteScriptHotWordRequest
	GetHotWordId() *string
}

type DeleteScriptHotWordRequest struct {
	HotWordId *string `json:"hotWordId,omitempty" xml:"hotWordId,omitempty"`
}

func (s DeleteScriptHotWordRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptHotWordRequest) GoString() string {
	return s.String()
}

func (s *DeleteScriptHotWordRequest) GetHotWordId() *string {
	return s.HotWordId
}

func (s *DeleteScriptHotWordRequest) SetHotWordId(v string) *DeleteScriptHotWordRequest {
	s.HotWordId = &v
	return s
}

func (s *DeleteScriptHotWordRequest) Validate() error {
	return dara.Validate(s)
}
