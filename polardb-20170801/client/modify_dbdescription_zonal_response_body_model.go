// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBDescriptionZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBDescriptionZonalResponseBody
	GetRequestId() *string
}

type ModifyDBDescriptionZonalResponseBody struct {
	// example:
	//
	// EB07CFF0-D8A4-5C76-AED7-D00E26FC2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBDescriptionZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBDescriptionZonalResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBDescriptionZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBDescriptionZonalResponseBody) SetRequestId(v string) *ModifyDBDescriptionZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBDescriptionZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
