// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPlaylistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProgramId(v string) *StopPlaylistResponseBody
	GetProgramId() *string
	SetRequestId(v string) *StopPlaylistResponseBody
	GetRequestId() *string
}

type StopPlaylistResponseBody struct {
	// The ID of the episode list.
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

func (s StopPlaylistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopPlaylistResponseBody) GoString() string {
	return s.String()
}

func (s *StopPlaylistResponseBody) GetProgramId() *string {
	return s.ProgramId
}

func (s *StopPlaylistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopPlaylistResponseBody) SetProgramId(v string) *StopPlaylistResponseBody {
	s.ProgramId = &v
	return s
}

func (s *StopPlaylistResponseBody) SetRequestId(v string) *StopPlaylistResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopPlaylistResponseBody) Validate() error {
	return dara.Validate(s)
}
