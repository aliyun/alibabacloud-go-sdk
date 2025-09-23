// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeSccModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBNodeSccModeResponseBody
	GetRequestId() *string
}

type ModifyDBNodeSccModeResponseBody struct {
	// example:
	//
	// E2FDB684-751D-424D-98B9-704BEA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBNodeSccModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeSccModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeSccModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBNodeSccModeResponseBody) SetRequestId(v string) *ModifyDBNodeSccModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBNodeSccModeResponseBody) Validate() error {
	return dara.Validate(s)
}
