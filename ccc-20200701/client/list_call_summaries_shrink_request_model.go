// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallSummariesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactIdListShrink(v string) *ListCallSummariesShrinkRequest
	GetContactIdListShrink() *string
	SetInstanceId(v string) *ListCallSummariesShrinkRequest
	GetInstanceId() *string
}

type ListCallSummariesShrinkRequest struct {
	ContactIdListShrink *string `json:"ContactIdList,omitempty" xml:"ContactIdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListCallSummariesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCallSummariesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCallSummariesShrinkRequest) GetContactIdListShrink() *string {
	return s.ContactIdListShrink
}

func (s *ListCallSummariesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCallSummariesShrinkRequest) SetContactIdListShrink(v string) *ListCallSummariesShrinkRequest {
	s.ContactIdListShrink = &v
	return s
}

func (s *ListCallSummariesShrinkRequest) SetInstanceId(v string) *ListCallSummariesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCallSummariesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
