// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAccessModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebAccessModeResponseBody
	GetRequestId() *string
}

type ModifyWebAccessModeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebAccessModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAccessModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebAccessModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebAccessModeResponseBody) SetRequestId(v string) *ModifyWebAccessModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebAccessModeResponseBody) Validate() error {
	return dara.Validate(s)
}
