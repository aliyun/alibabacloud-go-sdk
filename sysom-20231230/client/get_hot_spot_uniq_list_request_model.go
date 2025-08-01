// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotSpotUniqListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBegEnd(v int64) *GetHotSpotUniqListRequest
	GetBegEnd() *int64
	SetBegStart(v int64) *GetHotSpotUniqListRequest
	GetBegStart() *int64
	SetInstance(v string) *GetHotSpotUniqListRequest
	GetInstance() *string
	SetPid(v int64) *GetHotSpotUniqListRequest
	GetPid() *int64
	SetTable(v string) *GetHotSpotUniqListRequest
	GetTable() *string
	SetUniq(v string) *GetHotSpotUniqListRequest
	GetUniq() *string
}

type GetHotSpotUniqListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1735010351000
	BegEnd *int64 `json:"beg_end,omitempty" xml:"beg_end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1735008551000
	BegStart *int64 `json:"beg_start,omitempty" xml:"beg_start,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-bp1g2i0siirefgwnnnvy
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 12345
	Pid *int64 `json:"pid,omitempty" xml:"pid,omitempty"`
	// example:
	//
	// prof_on
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
	// This parameter is required.
	Uniq *string `json:"uniq,omitempty" xml:"uniq,omitempty"`
}

func (s GetHotSpotUniqListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotSpotUniqListRequest) GoString() string {
	return s.String()
}

func (s *GetHotSpotUniqListRequest) GetBegEnd() *int64 {
	return s.BegEnd
}

func (s *GetHotSpotUniqListRequest) GetBegStart() *int64 {
	return s.BegStart
}

func (s *GetHotSpotUniqListRequest) GetInstance() *string {
	return s.Instance
}

func (s *GetHotSpotUniqListRequest) GetPid() *int64 {
	return s.Pid
}

func (s *GetHotSpotUniqListRequest) GetTable() *string {
	return s.Table
}

func (s *GetHotSpotUniqListRequest) GetUniq() *string {
	return s.Uniq
}

func (s *GetHotSpotUniqListRequest) SetBegEnd(v int64) *GetHotSpotUniqListRequest {
	s.BegEnd = &v
	return s
}

func (s *GetHotSpotUniqListRequest) SetBegStart(v int64) *GetHotSpotUniqListRequest {
	s.BegStart = &v
	return s
}

func (s *GetHotSpotUniqListRequest) SetInstance(v string) *GetHotSpotUniqListRequest {
	s.Instance = &v
	return s
}

func (s *GetHotSpotUniqListRequest) SetPid(v int64) *GetHotSpotUniqListRequest {
	s.Pid = &v
	return s
}

func (s *GetHotSpotUniqListRequest) SetTable(v string) *GetHotSpotUniqListRequest {
	s.Table = &v
	return s
}

func (s *GetHotSpotUniqListRequest) SetUniq(v string) *GetHotSpotUniqListRequest {
	s.Uniq = &v
	return s
}

func (s *GetHotSpotUniqListRequest) Validate() error {
	return dara.Validate(s)
}
