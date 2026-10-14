// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeENSECKServiceRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InitializeENSECKServiceRoleResponseBody
	GetRequestId() *string
}

type InitializeENSECKServiceRoleResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F3B261DD-3858-4D3C-877D-303ADF374600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeENSECKServiceRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitializeENSECKServiceRoleResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeENSECKServiceRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitializeENSECKServiceRoleResponseBody) SetRequestId(v string) *InitializeENSECKServiceRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeENSECKServiceRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
