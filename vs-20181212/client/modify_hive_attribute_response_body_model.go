// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHiveAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHiveAttributeResponseBody
	GetRequestId() *string
}

type ModifyHiveAttributeResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHiveAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHiveAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHiveAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHiveAttributeResponseBody) SetRequestId(v string) *ModifyHiveAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHiveAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
