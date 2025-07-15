// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbanLiveMessageGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbanLiveMessageGroupResponseBody
	GetRequestId() *string
}

type UnbanLiveMessageGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 84AF36BF-0B39-1F8A-A416-FAC7C484****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbanLiveMessageGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbanLiveMessageGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UnbanLiveMessageGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbanLiveMessageGroupResponseBody) SetRequestId(v string) *UnbanLiveMessageGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbanLiveMessageGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
