// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMinorVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceMinorVersionResponseBody
	GetRequestId() *string
}

type ModifyInstanceMinorVersionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 28761557-0B33-41DF-AEEB-322DFF96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceMinorVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMinorVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceMinorVersionResponseBody) SetRequestId(v string) *ModifyInstanceMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceMinorVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
