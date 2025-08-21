// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlbumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetAlbumRequest
	GetId() *int64
	SetType(v string) *GetAlbumRequest
	GetType() *string
}

type GetAlbumRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12343
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// song
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAlbumRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlbumRequest) GoString() string {
	return s.String()
}

func (s *GetAlbumRequest) GetId() *int64 {
	return s.Id
}

func (s *GetAlbumRequest) GetType() *string {
	return s.Type
}

func (s *GetAlbumRequest) SetId(v int64) *GetAlbumRequest {
	s.Id = &v
	return s
}

func (s *GetAlbumRequest) SetType(v string) *GetAlbumRequest {
	s.Type = &v
	return s
}

func (s *GetAlbumRequest) Validate() error {
	return dara.Validate(s)
}
