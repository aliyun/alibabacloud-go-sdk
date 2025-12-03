// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbiddenVideoIds(v []*string) *UpdateVideoInfosResponseBody
	GetForbiddenVideoIds() []*string
	SetNonExistReferenceIds(v []*string) *UpdateVideoInfosResponseBody
	GetNonExistReferenceIds() []*string
	SetNonExistVideoIds(v []*string) *UpdateVideoInfosResponseBody
	GetNonExistVideoIds() []*string
	SetRequestId(v string) *UpdateVideoInfosResponseBody
	GetRequestId() *string
}

type UpdateVideoInfosResponseBody struct {
	// The IDs of the videos that cannot be modified. Generally, videos cannot be modified if you do not have required [permissions](https://help.aliyun.com/document_detail/113600.html).
	ForbiddenVideoIds    []*string `json:"ForbiddenVideoIds,omitempty" xml:"ForbiddenVideoIds,omitempty" type:"Repeated"`
	NonExistReferenceIds []*string `json:"NonExistReferenceIds,omitempty" xml:"NonExistReferenceIds,omitempty" type:"Repeated"`
	// The IDs of the videos that do not exist.
	NonExistVideoIds []*string `json:"NonExistVideoIds,omitempty" xml:"NonExistVideoIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVideoInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoInfosResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoInfosResponseBody) GetForbiddenVideoIds() []*string {
	return s.ForbiddenVideoIds
}

func (s *UpdateVideoInfosResponseBody) GetNonExistReferenceIds() []*string {
	return s.NonExistReferenceIds
}

func (s *UpdateVideoInfosResponseBody) GetNonExistVideoIds() []*string {
	return s.NonExistVideoIds
}

func (s *UpdateVideoInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVideoInfosResponseBody) SetForbiddenVideoIds(v []*string) *UpdateVideoInfosResponseBody {
	s.ForbiddenVideoIds = v
	return s
}

func (s *UpdateVideoInfosResponseBody) SetNonExistReferenceIds(v []*string) *UpdateVideoInfosResponseBody {
	s.NonExistReferenceIds = v
	return s
}

func (s *UpdateVideoInfosResponseBody) SetNonExistVideoIds(v []*string) *UpdateVideoInfosResponseBody {
	s.NonExistVideoIds = v
	return s
}

func (s *UpdateVideoInfosResponseBody) SetRequestId(v string) *UpdateVideoInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVideoInfosResponseBody) Validate() error {
	return dara.Validate(s)
}
