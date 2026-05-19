// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectClientEventDashboardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFileProtectClientEventDashboardResponseBodyData) *GetFileProtectClientEventDashboardResponseBody
	GetData() *GetFileProtectClientEventDashboardResponseBodyData
	SetRequestId(v string) *GetFileProtectClientEventDashboardResponseBody
	GetRequestId() *string
}

type GetFileProtectClientEventDashboardResponseBody struct {
	Data *GetFileProtectClientEventDashboardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// F5CF78A7-30AA-59DB-847F-13EE3AE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileProtectClientEventDashboardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientEventDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientEventDashboardResponseBody) GetData() *GetFileProtectClientEventDashboardResponseBodyData {
	return s.Data
}

func (s *GetFileProtectClientEventDashboardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileProtectClientEventDashboardResponseBody) SetData(v *GetFileProtectClientEventDashboardResponseBodyData) *GetFileProtectClientEventDashboardResponseBody {
	s.Data = v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBody) SetRequestId(v string) *GetFileProtectClientEventDashboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileProtectClientEventDashboardResponseBodyData struct {
	FilePathStats []*GetFileProtectClientEventDashboardResponseBodyDataFilePathStats `json:"FilePathStats,omitempty" xml:"FilePathStats,omitempty" type:"Repeated"`
	FileTypeStats []*GetFileProtectClientEventDashboardResponseBodyDataFileTypeStats `json:"FileTypeStats,omitempty" xml:"FileTypeStats,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	OneDayFileChangeCount *int32                                                                `json:"OneDayFileChangeCount,omitempty" xml:"OneDayFileChangeCount,omitempty"`
	ProcessNameStats      []*GetFileProtectClientEventDashboardResponseBodyDataProcessNameStats `json:"ProcessNameStats,omitempty" xml:"ProcessNameStats,omitempty" type:"Repeated"`
	// example:
	//
	// 33
	RecentFileChangeCount *int32 `json:"RecentFileChangeCount,omitempty" xml:"RecentFileChangeCount,omitempty"`
	// example:
	//
	// 4
	ServerCount *int32 `json:"ServerCount,omitempty" xml:"ServerCount,omitempty"`
}

func (s GetFileProtectClientEventDashboardResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientEventDashboardResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientEventDashboardResponseBodyData) GetFilePathStats() []*GetFileProtectClientEventDashboardResponseBodyDataFilePathStats {
	return s.FilePathStats
}

func (s *GetFileProtectClientEventDashboardResponseBodyData) GetFileTypeStats() []*GetFileProtectClientEventDashboardResponseBodyDataFileTypeStats {
	return s.FileTypeStats
}

func (s *GetFileProtectClientEventDashboardResponseBodyData) GetOneDayFileChangeCount() *int32 {
	return s.OneDayFileChangeCount
}

func (s *GetFileProtectClientEventDashboardResponseBodyData) GetProcessNameStats() []*GetFileProtectClientEventDashboardResponseBodyDataProcessNameStats {
	return s.ProcessNameStats
}

func (s *GetFileProtectClientEventDashboardResponseBodyData) GetRecentFileChangeCount() *int32 {
	return s.RecentFileChangeCount
}

func (s *GetFileProtectClientEventDashboardResponseBodyData) GetServerCount() *int32 {
	return s.ServerCount
}

func (s *GetFileProtectClientEventDashboardResponseBodyData) SetFilePathStats(v []*GetFileProtectClientEventDashboardResponseBodyDataFilePathStats) *GetFileProtectClientEventDashboardResponseBodyData {
	s.FilePathStats = v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBodyData) SetFileTypeStats(v []*GetFileProtectClientEventDashboardResponseBodyDataFileTypeStats) *GetFileProtectClientEventDashboardResponseBodyData {
	s.FileTypeStats = v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBodyData) SetOneDayFileChangeCount(v int32) *GetFileProtectClientEventDashboardResponseBodyData {
	s.OneDayFileChangeCount = &v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBodyData) SetProcessNameStats(v []*GetFileProtectClientEventDashboardResponseBodyDataProcessNameStats) *GetFileProtectClientEventDashboardResponseBodyData {
	s.ProcessNameStats = v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBodyData) SetRecentFileChangeCount(v int32) *GetFileProtectClientEventDashboardResponseBodyData {
	s.RecentFileChangeCount = &v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBodyData) SetServerCount(v int32) *GetFileProtectClientEventDashboardResponseBodyData {
	s.ServerCount = &v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBodyData) Validate() error {
	if s.FilePathStats != nil {
		for _, item := range s.FilePathStats {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FileTypeStats != nil {
		for _, item := range s.FileTypeStats {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ProcessNameStats != nil {
		for _, item := range s.ProcessNameStats {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFileProtectClientEventDashboardResponseBodyDataFilePathStats struct {
	// example:
	//
	// /usr/a
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 10
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
}

func (s GetFileProtectClientEventDashboardResponseBodyDataFilePathStats) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientEventDashboardResponseBodyDataFilePathStats) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataFilePathStats) GetKey() *string {
	return s.Key
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataFilePathStats) GetNum() *int64 {
	return s.Num
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataFilePathStats) SetKey(v string) *GetFileProtectClientEventDashboardResponseBodyDataFilePathStats {
	s.Key = &v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataFilePathStats) SetNum(v int64) *GetFileProtectClientEventDashboardResponseBodyDataFilePathStats {
	s.Num = &v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataFilePathStats) Validate() error {
	return dara.Validate(s)
}

type GetFileProtectClientEventDashboardResponseBodyDataFileTypeStats struct {
	// example:
	//
	// txt
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 60
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
}

func (s GetFileProtectClientEventDashboardResponseBodyDataFileTypeStats) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientEventDashboardResponseBodyDataFileTypeStats) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataFileTypeStats) GetKey() *string {
	return s.Key
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataFileTypeStats) GetNum() *int64 {
	return s.Num
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataFileTypeStats) SetKey(v string) *GetFileProtectClientEventDashboardResponseBodyDataFileTypeStats {
	s.Key = &v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataFileTypeStats) SetNum(v int64) *GetFileProtectClientEventDashboardResponseBodyDataFileTypeStats {
	s.Num = &v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataFileTypeStats) Validate() error {
	return dara.Validate(s)
}

type GetFileProtectClientEventDashboardResponseBodyDataProcessNameStats struct {
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 50
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
}

func (s GetFileProtectClientEventDashboardResponseBodyDataProcessNameStats) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientEventDashboardResponseBodyDataProcessNameStats) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataProcessNameStats) GetKey() *string {
	return s.Key
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataProcessNameStats) GetNum() *int64 {
	return s.Num
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataProcessNameStats) SetKey(v string) *GetFileProtectClientEventDashboardResponseBodyDataProcessNameStats {
	s.Key = &v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataProcessNameStats) SetNum(v int64) *GetFileProtectClientEventDashboardResponseBodyDataProcessNameStats {
	s.Num = &v
	return s
}

func (s *GetFileProtectClientEventDashboardResponseBodyDataProcessNameStats) Validate() error {
	return dara.Validate(s)
}
