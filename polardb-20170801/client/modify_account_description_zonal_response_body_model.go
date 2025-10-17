// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountDescriptionZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccountDescriptionZonalResponseBody
	GetRequestId() *string
}

type ModifyAccountDescriptionZonalResponseBody struct {
	// example:
	//
	// EB07CFF0-D8A4-5C76-AED7-D00E26FC2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountDescriptionZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountDescriptionZonalResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountDescriptionZonalResponseBody) SetRequestId(v string) *ModifyAccountDescriptionZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountDescriptionZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
