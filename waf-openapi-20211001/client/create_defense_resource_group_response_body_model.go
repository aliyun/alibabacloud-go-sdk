// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefenseResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDefenseResourceGroupResponseBody
	GetRequestId() *string
}

type CreateDefenseResourceGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5D69B341-4F97-58E7-9E16-1B17FE7A9E98
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDefenseResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDefenseResourceGroupResponseBody) SetRequestId(v string) *CreateDefenseResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDefenseResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
