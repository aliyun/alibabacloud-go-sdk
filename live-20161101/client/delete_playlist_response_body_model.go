// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlaylistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProgramId(v string) *DeletePlaylistResponseBody
	GetProgramId() *string
	SetRequestId(v string) *DeletePlaylistResponseBody
	GetRequestId() *string
}

type DeletePlaylistResponseBody struct {
	// The playlist ID. You can use this ID as a request parameter to query, start, or stop the playlist.
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

func (s DeletePlaylistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePlaylistResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePlaylistResponseBody) GetProgramId() *string {
	return s.ProgramId
}

func (s *DeletePlaylistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePlaylistResponseBody) SetProgramId(v string) *DeletePlaylistResponseBody {
	s.ProgramId = &v
	return s
}

func (s *DeletePlaylistResponseBody) SetRequestId(v string) *DeletePlaylistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePlaylistResponseBody) Validate() error {
	return dara.Validate(s)
}
