// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficMarkingPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTrafficMarkingPolicyResponseBody
	GetRequestId() *string
}

type DeleteTrafficMarkingPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5F1F3A57-A753-572B-8F71-4F964398C566
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrafficMarkingPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficMarkingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMarkingPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTrafficMarkingPolicyResponseBody) SetRequestId(v string) *DeleteTrafficMarkingPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
