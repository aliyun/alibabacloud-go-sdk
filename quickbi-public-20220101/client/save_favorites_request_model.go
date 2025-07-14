// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveFavoritesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *SaveFavoritesRequest
	GetUserId() *string
	SetWorksId(v string) *SaveFavoritesRequest
	GetWorksId() *string
}

type SaveFavoritesRequest struct {
	// The user ID of the person who favorites the work. This user ID is the UserID of Quick BI, not the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 121344444790****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the work being favorited.
	//
	// This parameter is required.
	//
	// example:
	//
	// d23e84a1-82a0-4292-bfdb-521306c3****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s SaveFavoritesRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveFavoritesRequest) GoString() string {
	return s.String()
}

func (s *SaveFavoritesRequest) GetUserId() *string {
	return s.UserId
}

func (s *SaveFavoritesRequest) GetWorksId() *string {
	return s.WorksId
}

func (s *SaveFavoritesRequest) SetUserId(v string) *SaveFavoritesRequest {
	s.UserId = &v
	return s
}

func (s *SaveFavoritesRequest) SetWorksId(v string) *SaveFavoritesRequest {
	s.WorksId = &v
	return s
}

func (s *SaveFavoritesRequest) Validate() error {
	return dara.Validate(s)
}
