// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSearchLibResponseBody
	GetCode() *string
	SetRequestId(v string) *ListSearchLibResponseBody
	GetRequestId() *string
	SetSearchLibInfoList(v []*ListSearchLibResponseBodySearchLibInfoList) *ListSearchLibResponseBody
	GetSearchLibInfoList() []*ListSearchLibResponseBodySearchLibInfoList
	SetSuccess(v string) *ListSearchLibResponseBody
	GetSuccess() *string
	SetTotal(v int64) *ListSearchLibResponseBody
	GetTotal() *int64
}

type ListSearchLibResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Information about search libraries.
	SearchLibInfoList []*ListSearchLibResponseBodySearchLibInfoList `json:"SearchLibInfoList,omitempty" xml:"SearchLibInfoList,omitempty" type:"Repeated"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 8
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSearchLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSearchLibResponseBody) GoString() string {
	return s.String()
}

func (s *ListSearchLibResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSearchLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSearchLibResponseBody) GetSearchLibInfoList() []*ListSearchLibResponseBodySearchLibInfoList {
	return s.SearchLibInfoList
}

func (s *ListSearchLibResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListSearchLibResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListSearchLibResponseBody) SetCode(v string) *ListSearchLibResponseBody {
	s.Code = &v
	return s
}

func (s *ListSearchLibResponseBody) SetRequestId(v string) *ListSearchLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSearchLibResponseBody) SetSearchLibInfoList(v []*ListSearchLibResponseBodySearchLibInfoList) *ListSearchLibResponseBody {
	s.SearchLibInfoList = v
	return s
}

func (s *ListSearchLibResponseBody) SetSuccess(v string) *ListSearchLibResponseBody {
	s.Success = &v
	return s
}

func (s *ListSearchLibResponseBody) SetTotal(v int64) *ListSearchLibResponseBody {
	s.Total = &v
	return s
}

func (s *ListSearchLibResponseBody) Validate() error {
	if s.SearchLibInfoList != nil {
		for _, item := range s.SearchLibInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSearchLibResponseBodySearchLibInfoList struct {
	IndexInfo       []*ListSearchLibResponseBodySearchLibInfoListIndexInfo `json:"IndexInfo,omitempty" xml:"IndexInfo,omitempty" type:"Repeated"`
	SearchLibConfig *string                                                `json:"SearchLibConfig,omitempty" xml:"SearchLibConfig,omitempty"`
	// The search library.
	//
	// example:
	//
	// faceSearchLib
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// The status of the search library.
	//
	// 	- normal
	//
	// 	- deleting
	//
	// 	- deleteFail
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSearchLibResponseBodySearchLibInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListSearchLibResponseBodySearchLibInfoList) GoString() string {
	return s.String()
}

func (s *ListSearchLibResponseBodySearchLibInfoList) GetIndexInfo() []*ListSearchLibResponseBodySearchLibInfoListIndexInfo {
	return s.IndexInfo
}

func (s *ListSearchLibResponseBodySearchLibInfoList) GetSearchLibConfig() *string {
	return s.SearchLibConfig
}

func (s *ListSearchLibResponseBodySearchLibInfoList) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *ListSearchLibResponseBodySearchLibInfoList) GetStatus() *string {
	return s.Status
}

func (s *ListSearchLibResponseBodySearchLibInfoList) SetIndexInfo(v []*ListSearchLibResponseBodySearchLibInfoListIndexInfo) *ListSearchLibResponseBodySearchLibInfoList {
	s.IndexInfo = v
	return s
}

func (s *ListSearchLibResponseBodySearchLibInfoList) SetSearchLibConfig(v string) *ListSearchLibResponseBodySearchLibInfoList {
	s.SearchLibConfig = &v
	return s
}

func (s *ListSearchLibResponseBodySearchLibInfoList) SetSearchLibName(v string) *ListSearchLibResponseBodySearchLibInfoList {
	s.SearchLibName = &v
	return s
}

func (s *ListSearchLibResponseBodySearchLibInfoList) SetStatus(v string) *ListSearchLibResponseBodySearchLibInfoList {
	s.Status = &v
	return s
}

func (s *ListSearchLibResponseBodySearchLibInfoList) Validate() error {
	if s.IndexInfo != nil {
		for _, item := range s.IndexInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSearchLibResponseBodySearchLibInfoListIndexInfo struct {
	// example:
	//
	// Normal
	IndexReadiness *string `json:"IndexReadiness,omitempty" xml:"IndexReadiness,omitempty"`
	// example:
	//
	// Active
	IndexStatus *string `json:"IndexStatus,omitempty" xml:"IndexStatus,omitempty"`
	// example:
	//
	// face
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
}

func (s ListSearchLibResponseBodySearchLibInfoListIndexInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSearchLibResponseBodySearchLibInfoListIndexInfo) GoString() string {
	return s.String()
}

func (s *ListSearchLibResponseBodySearchLibInfoListIndexInfo) GetIndexReadiness() *string {
	return s.IndexReadiness
}

func (s *ListSearchLibResponseBodySearchLibInfoListIndexInfo) GetIndexStatus() *string {
	return s.IndexStatus
}

func (s *ListSearchLibResponseBodySearchLibInfoListIndexInfo) GetIndexType() *string {
	return s.IndexType
}

func (s *ListSearchLibResponseBodySearchLibInfoListIndexInfo) SetIndexReadiness(v string) *ListSearchLibResponseBodySearchLibInfoListIndexInfo {
	s.IndexReadiness = &v
	return s
}

func (s *ListSearchLibResponseBodySearchLibInfoListIndexInfo) SetIndexStatus(v string) *ListSearchLibResponseBodySearchLibInfoListIndexInfo {
	s.IndexStatus = &v
	return s
}

func (s *ListSearchLibResponseBodySearchLibInfoListIndexInfo) SetIndexType(v string) *ListSearchLibResponseBodySearchLibInfoListIndexInfo {
	s.IndexType = &v
	return s
}

func (s *ListSearchLibResponseBodySearchLibInfoListIndexInfo) Validate() error {
	return dara.Validate(s)
}
