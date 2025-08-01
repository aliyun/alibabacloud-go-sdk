// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBegEnd(v int64) *GetHotspotInstanceListRequest
	GetBegEnd() *int64
	SetBegStart(v int64) *GetHotspotInstanceListRequest
	GetBegStart() *int64
	SetTable(v string) *GetHotspotInstanceListRequest
	GetTable() *string
}

type GetHotspotInstanceListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1725413947000
	BegEnd *int64 `json:"beg_end,omitempty" xml:"beg_end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725413647000
	BegStart *int64 `json:"beg_start,omitempty" xml:"beg_start,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// prof_on
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetHotspotInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotInstanceListRequest) GetBegEnd() *int64 {
	return s.BegEnd
}

func (s *GetHotspotInstanceListRequest) GetBegStart() *int64 {
	return s.BegStart
}

func (s *GetHotspotInstanceListRequest) GetTable() *string {
	return s.Table
}

func (s *GetHotspotInstanceListRequest) SetBegEnd(v int64) *GetHotspotInstanceListRequest {
	s.BegEnd = &v
	return s
}

func (s *GetHotspotInstanceListRequest) SetBegStart(v int64) *GetHotspotInstanceListRequest {
	s.BegStart = &v
	return s
}

func (s *GetHotspotInstanceListRequest) SetTable(v string) *GetHotspotInstanceListRequest {
	s.Table = &v
	return s
}

func (s *GetHotspotInstanceListRequest) Validate() error {
	return dara.Validate(s)
}
