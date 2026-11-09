// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateGroupIdResponseBody
	GetRequestId() *string
}

type CreateGroupIdResponseBody struct {
	// The request ID. This is a common parameter. Each request ID is unique.
	//
	// example:
	//
	// 2C7D722D-0F3D-4415-A9CD-A464D82C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupIdResponseBody) SetRequestId(v string) *CreateGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupIdResponseBody) Validate() error {
	return dara.Validate(s)
}
