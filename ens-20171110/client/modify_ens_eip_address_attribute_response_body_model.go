// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEnsEipAddressAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEnsEipAddressAttributeResponseBody
	GetRequestId() *string
}

type ModifyEnsEipAddressAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEnsEipAddressAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEnsEipAddressAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEnsEipAddressAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEnsEipAddressAttributeResponseBody) SetRequestId(v string) *ModifyEnsEipAddressAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEnsEipAddressAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
