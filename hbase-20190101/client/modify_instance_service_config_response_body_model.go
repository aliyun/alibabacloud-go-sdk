// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceServiceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceServiceConfigResponseBody
	GetRequestId() *string
}

type ModifyInstanceServiceConfigResponseBody struct {
	// example:
	//
	// F008B7AB-025D-4C20-AE12-047C8F8C3D97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceServiceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceServiceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceServiceConfigResponseBody) SetRequestId(v string) *ModifyInstanceServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceServiceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
