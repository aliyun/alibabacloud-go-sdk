// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAndroidInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAndroidInstanceResponseBody
	GetRequestId() *string
}

type ModifyAndroidInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E5138F7E-46B5-526A-8C99-82DEAE6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAndroidInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAndroidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAndroidInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAndroidInstanceResponseBody) SetRequestId(v string) *ModifyAndroidInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAndroidInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
