// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *JoinResourceGroupResponseBody
	GetRequestId() *string
}

type JoinResourceGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JoinResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s JoinResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *JoinResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *JoinResourceGroupResponseBody) SetRequestId(v string) *JoinResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
