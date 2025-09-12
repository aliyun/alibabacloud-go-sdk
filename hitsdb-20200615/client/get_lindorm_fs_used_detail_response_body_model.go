// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormFsUsedDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetLindormFsUsedDetailResponseBody
	GetAccessDeniedDetail() *string
	SetFsCapacity(v string) *GetLindormFsUsedDetailResponseBody
	GetFsCapacity() *string
	SetFsCapacityCold(v string) *GetLindormFsUsedDetailResponseBody
	GetFsCapacityCold() *string
	SetFsCapacityHot(v string) *GetLindormFsUsedDetailResponseBody
	GetFsCapacityHot() *string
	SetFsUsedCold(v string) *GetLindormFsUsedDetailResponseBody
	GetFsUsedCold() *string
	SetFsUsedColdOnLindormSearch(v string) *GetLindormFsUsedDetailResponseBody
	GetFsUsedColdOnLindormSearch() *string
	SetFsUsedColdOnLindormTSDB(v string) *GetLindormFsUsedDetailResponseBody
	GetFsUsedColdOnLindormTSDB() *string
	SetFsUsedColdOnLindormTable(v string) *GetLindormFsUsedDetailResponseBody
	GetFsUsedColdOnLindormTable() *string
	SetFsUsedHot(v string) *GetLindormFsUsedDetailResponseBody
	GetFsUsedHot() *string
	SetFsUsedHotOnLindormSearch(v string) *GetLindormFsUsedDetailResponseBody
	GetFsUsedHotOnLindormSearch() *string
	SetFsUsedHotOnLindormTSDB(v string) *GetLindormFsUsedDetailResponseBody
	GetFsUsedHotOnLindormTSDB() *string
	SetFsUsedHotOnLindormTable(v string) *GetLindormFsUsedDetailResponseBody
	GetFsUsedHotOnLindormTable() *string
	SetFsUsedOnLindormSearch(v string) *GetLindormFsUsedDetailResponseBody
	GetFsUsedOnLindormSearch() *string
	SetFsUsedOnLindormTSDB(v string) *GetLindormFsUsedDetailResponseBody
	GetFsUsedOnLindormTSDB() *string
	SetFsUsedOnLindormTable(v string) *GetLindormFsUsedDetailResponseBody
	GetFsUsedOnLindormTable() *string
	SetFsUsedOnLindormTableData(v string) *GetLindormFsUsedDetailResponseBody
	GetFsUsedOnLindormTableData() *string
	SetFsUsedOnLindormTableWAL(v string) *GetLindormFsUsedDetailResponseBody
	GetFsUsedOnLindormTableWAL() *string
	SetLStorageUsageList(v []*GetLindormFsUsedDetailResponseBodyLStorageUsageList) *GetLindormFsUsedDetailResponseBody
	GetLStorageUsageList() []*GetLindormFsUsedDetailResponseBodyLStorageUsageList
	SetRequestId(v string) *GetLindormFsUsedDetailResponseBody
	GetRequestId() *string
	SetValid(v string) *GetLindormFsUsedDetailResponseBody
	GetValid() *string
}

type GetLindormFsUsedDetailResponseBody struct {
	// The detailed reason why the access was denied.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The total storage space of the cluster. Unit: bytes.
	//
	// example:
	//
	// 85899345920
	FsCapacity *string `json:"FsCapacity,omitempty" xml:"FsCapacity,omitempty"`
	// The cold storage space of the cluster. Unit: bytes.
	//
	// example:
	//
	// 85899345920
	FsCapacityCold *string `json:"FsCapacityCold,omitempty" xml:"FsCapacityCold,omitempty"`
	// The hot storage space of the cluster. Unit: bytes.
	//
	// example:
	//
	// 85899345920
	FsCapacityHot *string `json:"FsCapacityHot,omitempty" xml:"FsCapacityHot,omitempty"`
	// The cold storage usage of the cluster. Unit: bytes.
	//
	// example:
	//
	// 33269
	FsUsedCold *string `json:"FsUsedCold,omitempty" xml:"FsUsedCold,omitempty"`
	// The cold storage usage of the table data of the search engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	FsUsedColdOnLindormSearch *string `json:"FsUsedColdOnLindormSearch,omitempty" xml:"FsUsedColdOnLindormSearch,omitempty"`
	// The cold storage usage of the table data of the time series engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	FsUsedColdOnLindormTSDB *string `json:"FsUsedColdOnLindormTSDB,omitempty" xml:"FsUsedColdOnLindormTSDB,omitempty"`
	// The cold storage usage of the table data of the wide table engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	FsUsedColdOnLindormTable *string `json:"FsUsedColdOnLindormTable,omitempty" xml:"FsUsedColdOnLindormTable,omitempty"`
	// The hot storage usage of the cluster. Unit: bytes.
	//
	// example:
	//
	// 33269
	FsUsedHot *string `json:"FsUsedHot,omitempty" xml:"FsUsedHot,omitempty"`
	// The hot storage usage of the table data of the search engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	FsUsedHotOnLindormSearch *string `json:"FsUsedHotOnLindormSearch,omitempty" xml:"FsUsedHotOnLindormSearch,omitempty"`
	// The hot storage usage of the table data of the time series engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	FsUsedHotOnLindormTSDB *string `json:"FsUsedHotOnLindormTSDB,omitempty" xml:"FsUsedHotOnLindormTSDB,omitempty"`
	// The hot storage usage of the table data of the wide table engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	FsUsedHotOnLindormTable *string `json:"FsUsedHotOnLindormTable,omitempty" xml:"FsUsedHotOnLindormTable,omitempty"`
	// The storage usage of the search engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	FsUsedOnLindormSearch *string `json:"FsUsedOnLindormSearch,omitempty" xml:"FsUsedOnLindormSearch,omitempty"`
	// The storage usage of the time series engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	FsUsedOnLindormTSDB *string `json:"FsUsedOnLindormTSDB,omitempty" xml:"FsUsedOnLindormTSDB,omitempty"`
	// The space usage of the wide table engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	FsUsedOnLindormTable *string `json:"FsUsedOnLindormTable,omitempty" xml:"FsUsedOnLindormTable,omitempty"`
	// The storage usage of the table data of the wide table engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	FsUsedOnLindormTableData *string `json:"FsUsedOnLindormTableData,omitempty" xml:"FsUsedOnLindormTableData,omitempty"`
	// The storage usage of the log data of the wide table engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	FsUsedOnLindormTableWAL *string `json:"FsUsedOnLindormTableWAL,omitempty" xml:"FsUsedOnLindormTableWAL,omitempty"`
	// If the version of the underlying storage engine is 4.1.9 or later, the storage usage values returned for the LStorageUsageList parameter prevail. Storage details are returned based on the storage type.
	LStorageUsageList []*GetLindormFsUsedDetailResponseBodyLStorageUsageList `json:"LStorageUsageList,omitempty" xml:"LStorageUsageList,omitempty" type:"Repeated"`
	// The request ID. Each request has a unique ID. You can use the request ID to locate and troubleshoot issues.
	//
	// example:
	//
	// 4F23D50C-400C-592C-9486-9D1E10179065
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the return value is valid. Valid values: true and false. If a value of false is returned, you must provide the request ID for troubleshooting.
	//
	// example:
	//
	// true
	Valid *string `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetLindormFsUsedDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLindormFsUsedDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormFsUsedDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsCapacity() *string {
	return s.FsCapacity
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsCapacityCold() *string {
	return s.FsCapacityCold
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsCapacityHot() *string {
	return s.FsCapacityHot
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsUsedCold() *string {
	return s.FsUsedCold
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsUsedColdOnLindormSearch() *string {
	return s.FsUsedColdOnLindormSearch
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsUsedColdOnLindormTSDB() *string {
	return s.FsUsedColdOnLindormTSDB
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsUsedColdOnLindormTable() *string {
	return s.FsUsedColdOnLindormTable
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsUsedHot() *string {
	return s.FsUsedHot
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsUsedHotOnLindormSearch() *string {
	return s.FsUsedHotOnLindormSearch
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsUsedHotOnLindormTSDB() *string {
	return s.FsUsedHotOnLindormTSDB
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsUsedHotOnLindormTable() *string {
	return s.FsUsedHotOnLindormTable
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsUsedOnLindormSearch() *string {
	return s.FsUsedOnLindormSearch
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsUsedOnLindormTSDB() *string {
	return s.FsUsedOnLindormTSDB
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsUsedOnLindormTable() *string {
	return s.FsUsedOnLindormTable
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsUsedOnLindormTableData() *string {
	return s.FsUsedOnLindormTableData
}

func (s *GetLindormFsUsedDetailResponseBody) GetFsUsedOnLindormTableWAL() *string {
	return s.FsUsedOnLindormTableWAL
}

func (s *GetLindormFsUsedDetailResponseBody) GetLStorageUsageList() []*GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	return s.LStorageUsageList
}

func (s *GetLindormFsUsedDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLindormFsUsedDetailResponseBody) GetValid() *string {
	return s.Valid
}

func (s *GetLindormFsUsedDetailResponseBody) SetAccessDeniedDetail(v string) *GetLindormFsUsedDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsCapacity(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsCapacity = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsCapacityCold(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsCapacityCold = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsCapacityHot(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsCapacityHot = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedCold(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedCold = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedColdOnLindormSearch(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedColdOnLindormSearch = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedColdOnLindormTSDB(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedColdOnLindormTSDB = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedColdOnLindormTable(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedColdOnLindormTable = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedHot(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedHot = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedHotOnLindormSearch(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedHotOnLindormSearch = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedHotOnLindormTSDB(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedHotOnLindormTSDB = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedHotOnLindormTable(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedHotOnLindormTable = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedOnLindormSearch(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedOnLindormSearch = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedOnLindormTSDB(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedOnLindormTSDB = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedOnLindormTable(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedOnLindormTable = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedOnLindormTableData(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedOnLindormTableData = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedOnLindormTableWAL(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedOnLindormTableWAL = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetLStorageUsageList(v []*GetLindormFsUsedDetailResponseBodyLStorageUsageList) *GetLindormFsUsedDetailResponseBody {
	s.LStorageUsageList = v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetRequestId(v string) *GetLindormFsUsedDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetValid(v string) *GetLindormFsUsedDetailResponseBody {
	s.Valid = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLindormFsUsedDetailResponseBodyLStorageUsageList struct {
	// The total storage capacity. Unit: bytes.
	//
	// example:
	//
	// 85899345920
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The storage type of the cluster. Valid values:
	//
	// 	- StandardCloudStorage
	//
	// 	- PerformanceCloudStorage
	//
	// 	- CapacityCloudStorage
	//
	// 	- LocalSsdStorage
	//
	// 	- LocalHddStorage
	//
	// 	- LocalEbsStorage
	//
	// example:
	//
	// StandardCloudStorage
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The storage usage. Unit: bytes.
	//
	// example:
	//
	// 33269
	Used                *string `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedLindormColumn3  *string `json:"UsedLindormColumn3,omitempty" xml:"UsedLindormColumn3,omitempty"`
	UsedLindormMessage3 *string `json:"UsedLindormMessage3,omitempty" xml:"UsedLindormMessage3,omitempty"`
	// The storage usage of the search engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	UsedLindormSearch *string `json:"UsedLindormSearch,omitempty" xml:"UsedLindormSearch,omitempty"`
	// The storage usage of the compute engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	UsedLindormSpark *string `json:"UsedLindormSpark,omitempty" xml:"UsedLindormSpark,omitempty"`
	// The storage usage of the wide table engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	UsedLindormTable *string `json:"UsedLindormTable,omitempty" xml:"UsedLindormTable,omitempty"`
	// The storage usage of the time series engine. Unit: bytes.
	//
	// example:
	//
	// 33269
	UsedLindormTsdb    *string `json:"UsedLindormTsdb,omitempty" xml:"UsedLindormTsdb,omitempty"`
	UsedLindormVector3 *string `json:"UsedLindormVector3,omitempty" xml:"UsedLindormVector3,omitempty"`
	// The storage usage of other resources, such as logs and recycle bins. Unit: bytes.
	//
	// example:
	//
	// 33269
	UsedOther *string `json:"UsedOther,omitempty" xml:"UsedOther,omitempty"`
}

func (s GetLindormFsUsedDetailResponseBodyLStorageUsageList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormFsUsedDetailResponseBodyLStorageUsageList) GoString() string {
	return s.String()
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) GetCapacity() *string {
	return s.Capacity
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) GetDiskType() *string {
	return s.DiskType
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) GetUsed() *string {
	return s.Used
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) GetUsedLindormColumn3() *string {
	return s.UsedLindormColumn3
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) GetUsedLindormMessage3() *string {
	return s.UsedLindormMessage3
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) GetUsedLindormSearch() *string {
	return s.UsedLindormSearch
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) GetUsedLindormSpark() *string {
	return s.UsedLindormSpark
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) GetUsedLindormTable() *string {
	return s.UsedLindormTable
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) GetUsedLindormTsdb() *string {
	return s.UsedLindormTsdb
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) GetUsedLindormVector3() *string {
	return s.UsedLindormVector3
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) GetUsedOther() *string {
	return s.UsedOther
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetCapacity(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.Capacity = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetDiskType(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.DiskType = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsed(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.Used = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsedLindormColumn3(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.UsedLindormColumn3 = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsedLindormMessage3(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.UsedLindormMessage3 = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsedLindormSearch(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.UsedLindormSearch = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsedLindormSpark(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.UsedLindormSpark = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsedLindormTable(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.UsedLindormTable = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsedLindormTsdb(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.UsedLindormTsdb = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsedLindormVector3(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.UsedLindormVector3 = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsedOther(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.UsedOther = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) Validate() error {
	return dara.Validate(s)
}
