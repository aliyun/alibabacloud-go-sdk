// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLiveMessageGroupResponseBody
	GetRequestId() *string
}

type ModifyLiveMessageGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 84AF36BF-0B39-1F8A-A416-FAC7C484****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLiveMessageGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLiveMessageGroupResponseBody) SetRequestId(v string) *ModifyLiveMessageGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLiveMessageGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
