// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDefenseTemplateResponseBody
	GetRequestId() *string
}

type ModifyDefenseTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1A68C85D-7467-5BB1-9A7E-2E8A5D96D88A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDefenseTemplateResponseBody) SetRequestId(v string) *ModifyDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDefenseTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
