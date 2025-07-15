// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAndroidInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAndroidInstanceGroupResponseBody
	GetRequestId() *string
}

type DeleteAndroidInstanceGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CB95E410-FD1D-53C5-9F7D-93CC44D7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAndroidInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAndroidInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAndroidInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAndroidInstanceGroupResponseBody) SetRequestId(v string) *DeleteAndroidInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAndroidInstanceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
