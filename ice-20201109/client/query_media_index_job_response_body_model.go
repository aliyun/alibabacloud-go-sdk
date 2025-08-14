// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaIndexJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryMediaIndexJobResponseBody
	GetCode() *string
	SetIndexJobInfoList(v []*QueryMediaIndexJobResponseBodyIndexJobInfoList) *QueryMediaIndexJobResponseBody
	GetIndexJobInfoList() []*QueryMediaIndexJobResponseBodyIndexJobInfoList
	SetRequestId(v string) *QueryMediaIndexJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *QueryMediaIndexJobResponseBody
	GetSuccess() *string
}

type QueryMediaIndexJobResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The indexing jobs enabled for the media asset.
	IndexJobInfoList []*QueryMediaIndexJobResponseBodyIndexJobInfoList `json:"IndexJobInfoList,omitempty" xml:"IndexJobInfoList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4E84BE44-58A7-****-****-FBEBEA16EF94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMediaIndexJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaIndexJobResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMediaIndexJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryMediaIndexJobResponseBody) GetIndexJobInfoList() []*QueryMediaIndexJobResponseBodyIndexJobInfoList {
	return s.IndexJobInfoList
}

func (s *QueryMediaIndexJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMediaIndexJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *QueryMediaIndexJobResponseBody) SetCode(v string) *QueryMediaIndexJobResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMediaIndexJobResponseBody) SetIndexJobInfoList(v []*QueryMediaIndexJobResponseBodyIndexJobInfoList) *QueryMediaIndexJobResponseBody {
	s.IndexJobInfoList = v
	return s
}

func (s *QueryMediaIndexJobResponseBody) SetRequestId(v string) *QueryMediaIndexJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMediaIndexJobResponseBody) SetSuccess(v string) *QueryMediaIndexJobResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMediaIndexJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMediaIndexJobResponseBodyIndexJobInfoList struct {
	// The end time of the indexing job.
	//
	// example:
	//
	// 2023-11-21 11:33:51
	GmtFinish *string `json:"GmtFinish,omitempty" xml:"GmtFinish,omitempty"`
	// The time when the index job was submitted.
	//
	// example:
	//
	// 2023-11-21 11:33:50
	GmtSubmit *string `json:"GmtSubmit,omitempty" xml:"GmtSubmit,omitempty"`
	// The index type. Valid values:
	//
	// 	- mm: large visual model.
	//
	// 	- face: face recognition.
	//
	// 	- aiLabel: smart tagging.
	//
	// example:
	//
	// mm
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// The job status. Valid values:
	//
	// 	- Running
	//
	// 	- Success
	//
	// 	- Fail
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryMediaIndexJobResponseBodyIndexJobInfoList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaIndexJobResponseBodyIndexJobInfoList) GoString() string {
	return s.String()
}

func (s *QueryMediaIndexJobResponseBodyIndexJobInfoList) GetGmtFinish() *string {
	return s.GmtFinish
}

func (s *QueryMediaIndexJobResponseBodyIndexJobInfoList) GetGmtSubmit() *string {
	return s.GmtSubmit
}

func (s *QueryMediaIndexJobResponseBodyIndexJobInfoList) GetIndexType() *string {
	return s.IndexType
}

func (s *QueryMediaIndexJobResponseBodyIndexJobInfoList) GetStatus() *string {
	return s.Status
}

func (s *QueryMediaIndexJobResponseBodyIndexJobInfoList) SetGmtFinish(v string) *QueryMediaIndexJobResponseBodyIndexJobInfoList {
	s.GmtFinish = &v
	return s
}

func (s *QueryMediaIndexJobResponseBodyIndexJobInfoList) SetGmtSubmit(v string) *QueryMediaIndexJobResponseBodyIndexJobInfoList {
	s.GmtSubmit = &v
	return s
}

func (s *QueryMediaIndexJobResponseBodyIndexJobInfoList) SetIndexType(v string) *QueryMediaIndexJobResponseBodyIndexJobInfoList {
	s.IndexType = &v
	return s
}

func (s *QueryMediaIndexJobResponseBodyIndexJobInfoList) SetStatus(v string) *QueryMediaIndexJobResponseBodyIndexJobInfoList {
	s.Status = &v
	return s
}

func (s *QueryMediaIndexJobResponseBodyIndexJobInfoList) Validate() error {
	return dara.Validate(s)
}
