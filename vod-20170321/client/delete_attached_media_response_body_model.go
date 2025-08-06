// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttachedMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistMediaIds(v []*string) *DeleteAttachedMediaResponseBody
	GetNonExistMediaIds() []*string
	SetRequestId(v string) *DeleteAttachedMediaResponseBody
	GetRequestId() *string
}

type DeleteAttachedMediaResponseBody struct {
	// The IDs of the auxiliary media assets that failed to be deleted.
	NonExistMediaIds []*string `json:"NonExistMediaIds,omitempty" xml:"NonExistMediaIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAttachedMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttachedMediaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAttachedMediaResponseBody) GetNonExistMediaIds() []*string {
	return s.NonExistMediaIds
}

func (s *DeleteAttachedMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAttachedMediaResponseBody) SetNonExistMediaIds(v []*string) *DeleteAttachedMediaResponseBody {
	s.NonExistMediaIds = v
	return s
}

func (s *DeleteAttachedMediaResponseBody) SetRequestId(v string) *DeleteAttachedMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAttachedMediaResponseBody) Validate() error {
	return dara.Validate(s)
}
