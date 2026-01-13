// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetRecallManagementJobResponseBody
	GetEndTime() *string
	SetLog(v string) *GetRecallManagementJobResponseBody
	GetLog() *string
	SetRecallManagementJobId(v string) *GetRecallManagementJobResponseBody
	GetRecallManagementJobId() *string
	SetRecallManagerTableInfo(v *GetRecallManagementJobResponseBodyRecallManagerTableInfo) *GetRecallManagementJobResponseBody
	GetRecallManagerTableInfo() *GetRecallManagementJobResponseBodyRecallManagerTableInfo
	SetRequestId(v string) *GetRecallManagementJobResponseBody
	GetRequestId() *string
	SetStartTime(v string) *GetRecallManagementJobResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetRecallManagementJobResponseBody
	GetStatus() *string
}

type GetRecallManagementJobResponseBody struct {
	EndTime                *string                                                   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Log                    *string                                                   `json:"Log,omitempty" xml:"Log,omitempty"`
	RecallManagementJobId  *string                                                   `json:"RecallManagementJobId,omitempty" xml:"RecallManagementJobId,omitempty"`
	RecallManagerTableInfo *GetRecallManagementJobResponseBodyRecallManagerTableInfo `json:"RecallManagerTableInfo,omitempty" xml:"RecallManagerTableInfo,omitempty" type:"Struct"`
	RequestId              *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime              *string                                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                 *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRecallManagementJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecallManagementJobResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetRecallManagementJobResponseBody) GetLog() *string {
	return s.Log
}

func (s *GetRecallManagementJobResponseBody) GetRecallManagementJobId() *string {
	return s.RecallManagementJobId
}

func (s *GetRecallManagementJobResponseBody) GetRecallManagerTableInfo() *GetRecallManagementJobResponseBodyRecallManagerTableInfo {
	return s.RecallManagerTableInfo
}

func (s *GetRecallManagementJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecallManagementJobResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetRecallManagementJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetRecallManagementJobResponseBody) SetEndTime(v string) *GetRecallManagementJobResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetRecallManagementJobResponseBody) SetLog(v string) *GetRecallManagementJobResponseBody {
	s.Log = &v
	return s
}

func (s *GetRecallManagementJobResponseBody) SetRecallManagementJobId(v string) *GetRecallManagementJobResponseBody {
	s.RecallManagementJobId = &v
	return s
}

func (s *GetRecallManagementJobResponseBody) SetRecallManagerTableInfo(v *GetRecallManagementJobResponseBodyRecallManagerTableInfo) *GetRecallManagementJobResponseBody {
	s.RecallManagerTableInfo = v
	return s
}

func (s *GetRecallManagementJobResponseBody) SetRequestId(v string) *GetRecallManagementJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecallManagementJobResponseBody) SetStartTime(v string) *GetRecallManagementJobResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetRecallManagementJobResponseBody) SetStatus(v string) *GetRecallManagementJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetRecallManagementJobResponseBody) Validate() error {
	if s.RecallManagerTableInfo != nil {
		if err := s.RecallManagerTableInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRecallManagementJobResponseBodyRecallManagerTableInfo struct {
	DataVersion                 *string `json:"DataVersion,omitempty" xml:"DataVersion,omitempty"`
	RecallManagerTableVersionId *string `json:"RecallManagerTableVersionId,omitempty" xml:"RecallManagerTableVersionId,omitempty"`
	SourceTableDataSize         *string `json:"SourceTableDataSize,omitempty" xml:"SourceTableDataSize,omitempty"`
	SourceTableRowCount         *string `json:"SourceTableRowCount,omitempty" xml:"SourceTableRowCount,omitempty"`
}

func (s GetRecallManagementJobResponseBodyRecallManagerTableInfo) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementJobResponseBodyRecallManagerTableInfo) GoString() string {
	return s.String()
}

func (s *GetRecallManagementJobResponseBodyRecallManagerTableInfo) GetDataVersion() *string {
	return s.DataVersion
}

func (s *GetRecallManagementJobResponseBodyRecallManagerTableInfo) GetRecallManagerTableVersionId() *string {
	return s.RecallManagerTableVersionId
}

func (s *GetRecallManagementJobResponseBodyRecallManagerTableInfo) GetSourceTableDataSize() *string {
	return s.SourceTableDataSize
}

func (s *GetRecallManagementJobResponseBodyRecallManagerTableInfo) GetSourceTableRowCount() *string {
	return s.SourceTableRowCount
}

func (s *GetRecallManagementJobResponseBodyRecallManagerTableInfo) SetDataVersion(v string) *GetRecallManagementJobResponseBodyRecallManagerTableInfo {
	s.DataVersion = &v
	return s
}

func (s *GetRecallManagementJobResponseBodyRecallManagerTableInfo) SetRecallManagerTableVersionId(v string) *GetRecallManagementJobResponseBodyRecallManagerTableInfo {
	s.RecallManagerTableVersionId = &v
	return s
}

func (s *GetRecallManagementJobResponseBodyRecallManagerTableInfo) SetSourceTableDataSize(v string) *GetRecallManagementJobResponseBodyRecallManagerTableInfo {
	s.SourceTableDataSize = &v
	return s
}

func (s *GetRecallManagementJobResponseBodyRecallManagerTableInfo) SetSourceTableRowCount(v string) *GetRecallManagementJobResponseBodyRecallManagerTableInfo {
	s.SourceTableRowCount = &v
	return s
}

func (s *GetRecallManagementJobResponseBodyRecallManagerTableInfo) Validate() error {
	return dara.Validate(s)
}
