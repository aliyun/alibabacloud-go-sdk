// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarFsMappingAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPolarFsMappingAuthResponseBody
	GetRequestId() *string
}

type ModifyPolarFsMappingAuthResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPolarFsMappingAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarFsMappingAuthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolarFsMappingAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPolarFsMappingAuthResponseBody) SetRequestId(v string) *ModifyPolarFsMappingAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPolarFsMappingAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
