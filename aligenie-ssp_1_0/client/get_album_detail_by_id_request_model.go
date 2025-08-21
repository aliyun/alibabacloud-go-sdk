// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlbumDetailByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlbumId(v string) *GetAlbumDetailByIdRequest
	GetAlbumId() *string
}

type GetAlbumDetailByIdRequest struct {
	// example:
	//
	// 51999575
	AlbumId *string `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
}

func (s GetAlbumDetailByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlbumDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *GetAlbumDetailByIdRequest) GetAlbumId() *string {
	return s.AlbumId
}

func (s *GetAlbumDetailByIdRequest) SetAlbumId(v string) *GetAlbumDetailByIdRequest {
	s.AlbumId = &v
	return s
}

func (s *GetAlbumDetailByIdRequest) Validate() error {
	return dara.Validate(s)
}
