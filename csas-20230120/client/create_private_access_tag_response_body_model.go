// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateAccessTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePrivateAccessTagResponseBody
	GetRequestId() *string
	SetTagId(v string) *CreatePrivateAccessTagResponseBody
	GetTagId() *string
}

type CreatePrivateAccessTagResponseBody struct {
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s CreatePrivateAccessTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrivateAccessTagResponseBody) GetTagId() *string {
	return s.TagId
}

func (s *CreatePrivateAccessTagResponseBody) SetRequestId(v string) *CreatePrivateAccessTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrivateAccessTagResponseBody) SetTagId(v string) *CreatePrivateAccessTagResponseBody {
	s.TagId = &v
	return s
}

func (s *CreatePrivateAccessTagResponseBody) Validate() error {
	return dara.Validate(s)
}
