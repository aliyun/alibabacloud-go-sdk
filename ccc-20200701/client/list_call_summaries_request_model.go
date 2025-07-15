// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallSummariesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactIdList(v []*string) *ListCallSummariesRequest
	GetContactIdList() []*string
	SetInstanceId(v string) *ListCallSummariesRequest
	GetInstanceId() *string
}

type ListCallSummariesRequest struct {
	ContactIdList []*string `json:"ContactIdList,omitempty" xml:"ContactIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListCallSummariesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCallSummariesRequest) GoString() string {
	return s.String()
}

func (s *ListCallSummariesRequest) GetContactIdList() []*string {
	return s.ContactIdList
}

func (s *ListCallSummariesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCallSummariesRequest) SetContactIdList(v []*string) *ListCallSummariesRequest {
	s.ContactIdList = v
	return s
}

func (s *ListCallSummariesRequest) SetInstanceId(v string) *ListCallSummariesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCallSummariesRequest) Validate() error {
	return dara.Validate(s)
}
