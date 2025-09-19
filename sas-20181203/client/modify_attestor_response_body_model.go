// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAttestorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAttestorResponseBody
	GetRequestId() *string
}

type ModifyAttestorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9E8B1D8F-DE1C-5421-81AA-**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAttestorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAttestorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAttestorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAttestorResponseBody) SetRequestId(v string) *ModifyAttestorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAttestorResponseBody) Validate() error {
	return dara.Validate(s)
}
