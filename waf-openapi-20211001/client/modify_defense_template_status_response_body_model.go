// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseTemplateStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDefenseTemplateStatusResponseBody
	GetRequestId() *string
}

type ModifyDefenseTemplateStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 58007AE3-65D9-57BA-ABB4-1A544015FB50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseTemplateStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseTemplateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDefenseTemplateStatusResponseBody) SetRequestId(v string) *ModifyDefenseTemplateStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDefenseTemplateStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
