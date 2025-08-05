// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMemberAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceMemberAttributesResponseBody
	GetRequestId() *string
}

type ModifyInstanceMemberAttributesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AB74E747-BA5C-492C-87DD-CEA67FCFFFE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceMemberAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMemberAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMemberAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceMemberAttributesResponseBody) SetRequestId(v string) *ModifyInstanceMemberAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceMemberAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}
