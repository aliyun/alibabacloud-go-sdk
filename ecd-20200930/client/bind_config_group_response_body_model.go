// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindConfigGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *BindConfigGroupResponseBody
	GetGroupId() *string
	SetRequestId(v string) *BindConfigGroupResponseBody
	GetRequestId() *string
}

type BindConfigGroupResponseBody struct {
	// The ID of the configuration group.
	//
	// example:
	//
	// ccg-0chlk9b65lj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E54EB497-D7B7-5F04-B744-D8DFA7B******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindConfigGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindConfigGroupResponseBody) GoString() string {
	return s.String()
}

func (s *BindConfigGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *BindConfigGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindConfigGroupResponseBody) SetGroupId(v string) *BindConfigGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *BindConfigGroupResponseBody) SetRequestId(v string) *BindConfigGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindConfigGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
