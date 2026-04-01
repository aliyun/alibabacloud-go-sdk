// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPGHbaConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPGHbaConfigResponseBody
	GetRequestId() *string
}

type ModifyPGHbaConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 458E0781-C46C-55F5-A0E5-1DD284B28A3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPGHbaConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPGHbaConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPGHbaConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPGHbaConfigResponseBody) SetRequestId(v string) *ModifyPGHbaConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPGHbaConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
