// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantHonorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedUserIds(v []*string) *GrantHonorResponseBody
	GetFailedUserIds() []*string
	SetRequestId(v string) *GrantHonorResponseBody
	GetRequestId() *string
	SetSuccessUserIds(v []*string) *GrantHonorResponseBody
	GetSuccessUserIds() []*string
}

type GrantHonorResponseBody struct {
	FailedUserIds []*string `json:"failedUserIds,omitempty" xml:"failedUserIds,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 45b4d029-ab94-4672-aa0f-bd79590374cb
	RequestId      *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SuccessUserIds []*string `json:"successUserIds,omitempty" xml:"successUserIds,omitempty" type:"Repeated"`
}

func (s GrantHonorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantHonorResponseBody) GoString() string {
	return s.String()
}

func (s *GrantHonorResponseBody) GetFailedUserIds() []*string {
	return s.FailedUserIds
}

func (s *GrantHonorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantHonorResponseBody) GetSuccessUserIds() []*string {
	return s.SuccessUserIds
}

func (s *GrantHonorResponseBody) SetFailedUserIds(v []*string) *GrantHonorResponseBody {
	s.FailedUserIds = v
	return s
}

func (s *GrantHonorResponseBody) SetRequestId(v string) *GrantHonorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantHonorResponseBody) SetSuccessUserIds(v []*string) *GrantHonorResponseBody {
	s.SuccessUserIds = v
	return s
}

func (s *GrantHonorResponseBody) Validate() error {
	return dara.Validate(s)
}
