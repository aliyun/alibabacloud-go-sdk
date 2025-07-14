// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *CancelCollectionRequest
	GetUserId() *string
	SetWorksId(v string) *CancelCollectionRequest
	GetWorksId() *string
}

type CancelCollectionRequest struct {
	// The ID of the favorite user. The user ID is the UserID of the Quick BI, not the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 121344444790****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the work to cancel the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5d6ae4e7-cede-43cd-b4d3-d2fd442a9202
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s CancelCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCollectionRequest) GoString() string {
	return s.String()
}

func (s *CancelCollectionRequest) GetUserId() *string {
	return s.UserId
}

func (s *CancelCollectionRequest) GetWorksId() *string {
	return s.WorksId
}

func (s *CancelCollectionRequest) SetUserId(v string) *CancelCollectionRequest {
	s.UserId = &v
	return s
}

func (s *CancelCollectionRequest) SetWorksId(v string) *CancelCollectionRequest {
	s.WorksId = &v
	return s
}

func (s *CancelCollectionRequest) Validate() error {
	return dara.Validate(s)
}
