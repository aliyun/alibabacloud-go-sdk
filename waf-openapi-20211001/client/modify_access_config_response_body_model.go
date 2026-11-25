// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccessConfigResponseBody
	GetRequestId() *string
}

type ModifyAccessConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccessConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccessConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccessConfigResponseBody) SetRequestId(v string) *ModifyAccessConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccessConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
