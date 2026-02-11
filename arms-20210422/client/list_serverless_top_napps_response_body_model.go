// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerlessTopNAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListServerlessTopNAppsResponseBodyData) *ListServerlessTopNAppsResponseBody
	GetData() []*ListServerlessTopNAppsResponseBodyData
	SetRequestId(v string) *ListServerlessTopNAppsResponseBody
	GetRequestId() *string
}

type ListServerlessTopNAppsResponseBody struct {
	Data      []*ListServerlessTopNAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListServerlessTopNAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServerlessTopNAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerlessTopNAppsResponseBody) GetData() []*ListServerlessTopNAppsResponseBodyData {
	return s.Data
}

func (s *ListServerlessTopNAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServerlessTopNAppsResponseBody) SetData(v []*ListServerlessTopNAppsResponseBodyData) *ListServerlessTopNAppsResponseBody {
	s.Data = v
	return s
}

func (s *ListServerlessTopNAppsResponseBody) SetRequestId(v string) *ListServerlessTopNAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerlessTopNAppsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServerlessTopNAppsResponseBodyData struct {
	Count  *int32   `json:"Count,omitempty" xml:"Count,omitempty"`
	Error  *int32   `json:"Error,omitempty" xml:"Error,omitempty"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Pid    *string  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Region *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	Rt     *float32 `json:"Rt,omitempty" xml:"Rt,omitempty"`
}

func (s ListServerlessTopNAppsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListServerlessTopNAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServerlessTopNAppsResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *ListServerlessTopNAppsResponseBodyData) GetError() *int32 {
	return s.Error
}

func (s *ListServerlessTopNAppsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListServerlessTopNAppsResponseBodyData) GetPid() *string {
	return s.Pid
}

func (s *ListServerlessTopNAppsResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *ListServerlessTopNAppsResponseBodyData) GetRt() *float32 {
	return s.Rt
}

func (s *ListServerlessTopNAppsResponseBodyData) SetCount(v int32) *ListServerlessTopNAppsResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListServerlessTopNAppsResponseBodyData) SetError(v int32) *ListServerlessTopNAppsResponseBodyData {
	s.Error = &v
	return s
}

func (s *ListServerlessTopNAppsResponseBodyData) SetName(v string) *ListServerlessTopNAppsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListServerlessTopNAppsResponseBodyData) SetPid(v string) *ListServerlessTopNAppsResponseBodyData {
	s.Pid = &v
	return s
}

func (s *ListServerlessTopNAppsResponseBodyData) SetRegion(v string) *ListServerlessTopNAppsResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListServerlessTopNAppsResponseBodyData) SetRt(v float32) *ListServerlessTopNAppsResponseBodyData {
	s.Rt = &v
	return s
}

func (s *ListServerlessTopNAppsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
