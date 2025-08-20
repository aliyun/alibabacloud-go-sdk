// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTemplatePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetTemplatePermissionResponseBody
	GetRequestId() *string
}

type SetTemplatePermissionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetTemplatePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetTemplatePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *SetTemplatePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetTemplatePermissionResponseBody) SetRequestId(v string) *SetTemplatePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetTemplatePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
