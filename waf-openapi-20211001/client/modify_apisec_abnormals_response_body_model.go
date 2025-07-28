// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecAbnormalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApisecAbnormalsResponseBody
	GetRequestId() *string
}

type ModifyApisecAbnormalsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C1823E96-EF4B-5BD2-9E02-1D18****3ED8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApisecAbnormalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecAbnormalsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApisecAbnormalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApisecAbnormalsResponseBody) SetRequestId(v string) *ModifyApisecAbnormalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApisecAbnormalsResponseBody) Validate() error {
	return dara.Validate(s)
}
