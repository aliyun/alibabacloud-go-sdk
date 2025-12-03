// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbiddenVideoIds(v []*string) *DeleteVideoResponseBody
	GetForbiddenVideoIds() []*string
	SetNonExistReferenceIds(v []*string) *DeleteVideoResponseBody
	GetNonExistReferenceIds() []*string
	SetNonExistVideoIds(v []*string) *DeleteVideoResponseBody
	GetNonExistVideoIds() []*string
	SetRequestId(v string) *DeleteVideoResponseBody
	GetRequestId() *string
}

type DeleteVideoResponseBody struct {
	// The IDs of the videos that cannot be deleted.
	//
	// > Generally, videos cannot be deleted if you do not have the required [permissions](https://help.aliyun.com/document_detail/113600.html).
	ForbiddenVideoIds    []*string `json:"ForbiddenVideoIds,omitempty" xml:"ForbiddenVideoIds,omitempty" type:"Repeated"`
	NonExistReferenceIds []*string `json:"NonExistReferenceIds,omitempty" xml:"NonExistReferenceIds,omitempty" type:"Repeated"`
	// The IDs of the videos that do not exist.
	//
	// > If the list of videos to be deleted contains one or more videos that do not exist, the IDs of these non-existing videos are returned. If none of the videos in the list exists, a 404 error is returned.
	NonExistVideoIds []*string `json:"NonExistVideoIds,omitempty" xml:"NonExistVideoIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// E4EBD2BF-5EB0-4476-8829-9D94E1B1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVideoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVideoResponseBody) GetForbiddenVideoIds() []*string {
	return s.ForbiddenVideoIds
}

func (s *DeleteVideoResponseBody) GetNonExistReferenceIds() []*string {
	return s.NonExistReferenceIds
}

func (s *DeleteVideoResponseBody) GetNonExistVideoIds() []*string {
	return s.NonExistVideoIds
}

func (s *DeleteVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVideoResponseBody) SetForbiddenVideoIds(v []*string) *DeleteVideoResponseBody {
	s.ForbiddenVideoIds = v
	return s
}

func (s *DeleteVideoResponseBody) SetNonExistReferenceIds(v []*string) *DeleteVideoResponseBody {
	s.NonExistReferenceIds = v
	return s
}

func (s *DeleteVideoResponseBody) SetNonExistVideoIds(v []*string) *DeleteVideoResponseBody {
	s.NonExistVideoIds = v
	return s
}

func (s *DeleteVideoResponseBody) SetRequestId(v string) *DeleteVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVideoResponseBody) Validate() error {
	return dara.Validate(s)
}
