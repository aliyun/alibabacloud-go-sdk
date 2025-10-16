// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAclCheckResponseBody
	GetRequestId() *string
	SetTaskIds(v []*string) *CreateAclCheckResponseBody
	GetTaskIds() []*string
}

type CreateAclCheckResponseBody struct {
	// example:
	//
	// 4FB718F0-CC04-5A12-B17B-188CFC3F****
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskIds   []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s CreateAclCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAclCheckResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAclCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAclCheckResponseBody) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *CreateAclCheckResponseBody) SetRequestId(v string) *CreateAclCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAclCheckResponseBody) SetTaskIds(v []*string) *CreateAclCheckResponseBody {
	s.TaskIds = v
	return s
}

func (s *CreateAclCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
