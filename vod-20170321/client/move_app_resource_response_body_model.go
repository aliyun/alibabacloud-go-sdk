// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveAppResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedResourceIds(v []*string) *MoveAppResourceResponseBody
	GetFailedResourceIds() []*string
	SetNonExistResourceIds(v []*string) *MoveAppResourceResponseBody
	GetNonExistResourceIds() []*string
	SetRequestId(v string) *MoveAppResourceResponseBody
	GetRequestId() *string
}

type MoveAppResourceResponseBody struct {
	// The IDs of the resources that failed to be migrated.
	FailedResourceIds []*string `json:"FailedResourceIds,omitempty" xml:"FailedResourceIds,omitempty" type:"Repeated"`
	// The IDs of the resources that were not found.
	NonExistResourceIds []*string `json:"NonExistResourceIds,omitempty" xml:"NonExistResourceIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveAppResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveAppResourceResponseBody) GoString() string {
	return s.String()
}

func (s *MoveAppResourceResponseBody) GetFailedResourceIds() []*string {
	return s.FailedResourceIds
}

func (s *MoveAppResourceResponseBody) GetNonExistResourceIds() []*string {
	return s.NonExistResourceIds
}

func (s *MoveAppResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveAppResourceResponseBody) SetFailedResourceIds(v []*string) *MoveAppResourceResponseBody {
	s.FailedResourceIds = v
	return s
}

func (s *MoveAppResourceResponseBody) SetNonExistResourceIds(v []*string) *MoveAppResourceResponseBody {
	s.NonExistResourceIds = v
	return s
}

func (s *MoveAppResourceResponseBody) SetRequestId(v string) *MoveAppResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveAppResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
