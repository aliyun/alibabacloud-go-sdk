// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlaylistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProgramList(v []*ListPlaylistResponseBodyProgramList) *ListPlaylistResponseBody
	GetProgramList() []*ListPlaylistResponseBodyProgramList
	SetRequestId(v string) *ListPlaylistResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListPlaylistResponseBody
	GetTotal() *int32
}

type ListPlaylistResponseBody struct {
	// The episode lists.
	ProgramList []*ListPlaylistResponseBodyProgramList `json:"ProgramList,omitempty" xml:"ProgramList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of episode lists returned.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListPlaylistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPlaylistResponseBody) GoString() string {
	return s.String()
}

func (s *ListPlaylistResponseBody) GetProgramList() []*ListPlaylistResponseBodyProgramList {
	return s.ProgramList
}

func (s *ListPlaylistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPlaylistResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListPlaylistResponseBody) SetProgramList(v []*ListPlaylistResponseBodyProgramList) *ListPlaylistResponseBody {
	s.ProgramList = v
	return s
}

func (s *ListPlaylistResponseBody) SetRequestId(v string) *ListPlaylistResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPlaylistResponseBody) SetTotal(v int32) *ListPlaylistResponseBody {
	s.Total = &v
	return s
}

func (s *ListPlaylistResponseBody) Validate() error {
	if s.ProgramList != nil {
		for _, item := range s.ProgramList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPlaylistResponseBodyProgramList struct {
	// The ID of the production studio to which the episode list belongs. You can use the ID as a request parameter in the API operation that is used to add a layout to the virtual studio, delete a layout in the virtual studio, modify a layout of the virtual studio, or query layouts of the virtual studio.
	//
	// example:
	//
	// casdfasdfasfdasdflkasjdflaj****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the episode list.
	//
	// example:
	//
	// c09f3d63-eacf-4fbf-bd48-a07a6ba7****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The name of the episode list.
	//
	// example:
	//
	// playlist1
	ProgramName *string `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	// The number of playbacks after the first playback is complete. Valid values:
	//
	// 	- **0*	- (default): specifies that the episode list is played only once.
	//
	// 	- **-1**: specifies that the episode list is played in loop mode.
	//
	// 	- Positive integer: specifies the number of times the episode list repeats after the first playback is complete.
	//
	// example:
	//
	// 1
	RepeatNumber *int32 `json:"RepeatNumber,omitempty" xml:"RepeatNumber,omitempty"`
	// The status of the episode list. Valid values:
	//
	// 	- **0**: stopped
	//
	// 	- **1**: running
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPlaylistResponseBodyProgramList) String() string {
	return dara.Prettify(s)
}

func (s ListPlaylistResponseBodyProgramList) GoString() string {
	return s.String()
}

func (s *ListPlaylistResponseBodyProgramList) GetCasterId() *string {
	return s.CasterId
}

func (s *ListPlaylistResponseBodyProgramList) GetDomainName() *string {
	return s.DomainName
}

func (s *ListPlaylistResponseBodyProgramList) GetProgramId() *string {
	return s.ProgramId
}

func (s *ListPlaylistResponseBodyProgramList) GetProgramName() *string {
	return s.ProgramName
}

func (s *ListPlaylistResponseBodyProgramList) GetRepeatNumber() *int32 {
	return s.RepeatNumber
}

func (s *ListPlaylistResponseBodyProgramList) GetStatus() *int32 {
	return s.Status
}

func (s *ListPlaylistResponseBodyProgramList) SetCasterId(v string) *ListPlaylistResponseBodyProgramList {
	s.CasterId = &v
	return s
}

func (s *ListPlaylistResponseBodyProgramList) SetDomainName(v string) *ListPlaylistResponseBodyProgramList {
	s.DomainName = &v
	return s
}

func (s *ListPlaylistResponseBodyProgramList) SetProgramId(v string) *ListPlaylistResponseBodyProgramList {
	s.ProgramId = &v
	return s
}

func (s *ListPlaylistResponseBodyProgramList) SetProgramName(v string) *ListPlaylistResponseBodyProgramList {
	s.ProgramName = &v
	return s
}

func (s *ListPlaylistResponseBodyProgramList) SetRepeatNumber(v int32) *ListPlaylistResponseBodyProgramList {
	s.RepeatNumber = &v
	return s
}

func (s *ListPlaylistResponseBodyProgramList) SetStatus(v int32) *ListPlaylistResponseBodyProgramList {
	s.Status = &v
	return s
}

func (s *ListPlaylistResponseBodyProgramList) Validate() error {
	return dara.Validate(s)
}
