// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLastUpgradeRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetLastUpgradeRecordResponseBodyData) *GetLastUpgradeRecordResponseBody
	GetData() *GetLastUpgradeRecordResponseBodyData
	SetRequestId(v string) *GetLastUpgradeRecordResponseBody
	GetRequestId() *string
}

type GetLastUpgradeRecordResponseBody struct {
	Data *GetLastUpgradeRecordResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 819A7F0F-2951-540F-BD94-6A41ECF0281F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetLastUpgradeRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLastUpgradeRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetLastUpgradeRecordResponseBody) GetData() *GetLastUpgradeRecordResponseBodyData {
	return s.Data
}

func (s *GetLastUpgradeRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLastUpgradeRecordResponseBody) SetData(v *GetLastUpgradeRecordResponseBodyData) *GetLastUpgradeRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetLastUpgradeRecordResponseBody) SetRequestId(v string) *GetLastUpgradeRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLastUpgradeRecordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLastUpgradeRecordResponseBodyData struct {
	// example:
	//
	// r3.2.11
	FromVersion *string `json:"fromVersion,omitempty" xml:"fromVersion,omitempty"`
	// example:
	//
	// 2025-09-24 17:29:57
	StartReadonlyTime *string `json:"startReadonlyTime,omitempty" xml:"startReadonlyTime,omitempty"`
	// example:
	//
	// 2025-09-24 18:23:22
	StopReadonlyTime *string `json:"stopReadonlyTime,omitempty" xml:"stopReadonlyTime,omitempty"`
	// example:
	//
	// r4.0.2
	ToVersion *string `json:"toVersion,omitempty" xml:"toVersion,omitempty"`
}

func (s GetLastUpgradeRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLastUpgradeRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLastUpgradeRecordResponseBodyData) GetFromVersion() *string {
	return s.FromVersion
}

func (s *GetLastUpgradeRecordResponseBodyData) GetStartReadonlyTime() *string {
	return s.StartReadonlyTime
}

func (s *GetLastUpgradeRecordResponseBodyData) GetStopReadonlyTime() *string {
	return s.StopReadonlyTime
}

func (s *GetLastUpgradeRecordResponseBodyData) GetToVersion() *string {
	return s.ToVersion
}

func (s *GetLastUpgradeRecordResponseBodyData) SetFromVersion(v string) *GetLastUpgradeRecordResponseBodyData {
	s.FromVersion = &v
	return s
}

func (s *GetLastUpgradeRecordResponseBodyData) SetStartReadonlyTime(v string) *GetLastUpgradeRecordResponseBodyData {
	s.StartReadonlyTime = &v
	return s
}

func (s *GetLastUpgradeRecordResponseBodyData) SetStopReadonlyTime(v string) *GetLastUpgradeRecordResponseBodyData {
	s.StopReadonlyTime = &v
	return s
}

func (s *GetLastUpgradeRecordResponseBodyData) SetToVersion(v string) *GetLastUpgradeRecordResponseBodyData {
	s.ToVersion = &v
	return s
}

func (s *GetLastUpgradeRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
