// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectClientEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFileProtectClientEventResponseBody
	GetRequestId() *string
}

type UpdateFileProtectClientEventResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 69BFFCDE-37D6-5A49-A8BC-BB03AC83****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFileProtectClientEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectClientEventResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectClientEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileProtectClientEventResponseBody) SetRequestId(v string) *UpdateFileProtectClientEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileProtectClientEventResponseBody) Validate() error {
	return dara.Validate(s)
}
