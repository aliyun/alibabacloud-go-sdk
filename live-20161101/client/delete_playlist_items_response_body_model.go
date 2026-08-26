// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlaylistItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProgramId(v string) *DeletePlaylistItemsResponseBody
	GetProgramId() *string
	SetRequestId(v string) *DeletePlaylistItemsResponseBody
	GetRequestId() *string
}

type DeletePlaylistItemsResponseBody struct {
	// The ID of the playlist. You can use this ID as a request parameter to query playlist items, edit a playlist, delete a playlist, start a playlist, or stop a playlist.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePlaylistItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePlaylistItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePlaylistItemsResponseBody) GetProgramId() *string {
	return s.ProgramId
}

func (s *DeletePlaylistItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePlaylistItemsResponseBody) SetProgramId(v string) *DeletePlaylistItemsResponseBody {
	s.ProgramId = &v
	return s
}

func (s *DeletePlaylistItemsResponseBody) SetRequestId(v string) *DeletePlaylistItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePlaylistItemsResponseBody) Validate() error {
	return dara.Validate(s)
}
