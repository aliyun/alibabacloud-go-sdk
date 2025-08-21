// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBlackholeStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBlackholeStatusResponseBody
	GetRequestId() *string
}

type ModifyBlackholeStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBlackholeStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBlackholeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBlackholeStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBlackholeStatusResponseBody) SetRequestId(v string) *ModifyBlackholeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBlackholeStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
