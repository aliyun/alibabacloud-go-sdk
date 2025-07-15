// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverLiveMessageDeletedGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *RecoverLiveMessageDeletedGroupResponseBody
	GetGroupId() *string
	SetRequestId(v string) *RecoverLiveMessageDeletedGroupResponseBody
	GetRequestId() *string
}

type RecoverLiveMessageDeletedGroupResponseBody struct {
	// The ID of the group that was restored.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B95BE680-5A6A-1CAD-8AB1-09DFF5D6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecoverLiveMessageDeletedGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecoverLiveMessageDeletedGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverLiveMessageDeletedGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *RecoverLiveMessageDeletedGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecoverLiveMessageDeletedGroupResponseBody) SetGroupId(v string) *RecoverLiveMessageDeletedGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *RecoverLiveMessageDeletedGroupResponseBody) SetRequestId(v string) *RecoverLiveMessageDeletedGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverLiveMessageDeletedGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
