// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAttachedMediaInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistMediaIds(v []*string) *UpdateAttachedMediaInfosResponseBody
	GetNonExistMediaIds() []*string
	SetRequestId(v string) *UpdateAttachedMediaInfosResponseBody
	GetRequestId() *string
}

type UpdateAttachedMediaInfosResponseBody struct {
	// The IDs of the auxiliary media assets that do not exist.
	NonExistMediaIds []*string `json:"NonExistMediaIds,omitempty" xml:"NonExistMediaIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4DF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAttachedMediaInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttachedMediaInfosResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAttachedMediaInfosResponseBody) GetNonExistMediaIds() []*string {
	return s.NonExistMediaIds
}

func (s *UpdateAttachedMediaInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAttachedMediaInfosResponseBody) SetNonExistMediaIds(v []*string) *UpdateAttachedMediaInfosResponseBody {
	s.NonExistMediaIds = v
	return s
}

func (s *UpdateAttachedMediaInfosResponseBody) SetRequestId(v string) *UpdateAttachedMediaInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAttachedMediaInfosResponseBody) Validate() error {
	return dara.Validate(s)
}
