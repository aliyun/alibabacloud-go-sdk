// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVersionDescResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVersionDescResponseBody
	GetRequestId() *string
}

type UpdateVersionDescResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 52DD50C2-C381-13BD-A269-5FAEEB848ACD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVersionDescResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVersionDescResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVersionDescResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVersionDescResponseBody) SetRequestId(v string) *UpdateVersionDescResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVersionDescResponseBody) Validate() error {
	return dara.Validate(s)
}
