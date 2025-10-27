// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSummaryAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSummaryAppsResponseBody
	GetRequestId() *string
	SetSummaryAppInfos(v *ListSummaryAppsResponseBodySummaryAppInfos) *ListSummaryAppsResponseBody
	GetSummaryAppInfos() *ListSummaryAppsResponseBodySummaryAppInfos
}

type ListSummaryAppsResponseBody struct {
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId       *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SummaryAppInfos *ListSummaryAppsResponseBodySummaryAppInfos `json:"SummaryAppInfos,omitempty" xml:"SummaryAppInfos,omitempty" type:"Struct"`
}

func (s ListSummaryAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSummaryAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSummaryAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSummaryAppsResponseBody) GetSummaryAppInfos() *ListSummaryAppsResponseBodySummaryAppInfos {
	return s.SummaryAppInfos
}

func (s *ListSummaryAppsResponseBody) SetRequestId(v string) *ListSummaryAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSummaryAppsResponseBody) SetSummaryAppInfos(v *ListSummaryAppsResponseBodySummaryAppInfos) *ListSummaryAppsResponseBody {
	s.SummaryAppInfos = v
	return s
}

func (s *ListSummaryAppsResponseBody) Validate() error {
	if s.SummaryAppInfos != nil {
		if err := s.SummaryAppInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSummaryAppsResponseBodySummaryAppInfos struct {
	SummaryAppInfo []*ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo `json:"SummaryAppInfo,omitempty" xml:"SummaryAppInfo,omitempty" type:"Repeated"`
}

func (s ListSummaryAppsResponseBodySummaryAppInfos) String() string {
	return dara.Prettify(s)
}

func (s ListSummaryAppsResponseBodySummaryAppInfos) GoString() string {
	return s.String()
}

func (s *ListSummaryAppsResponseBodySummaryAppInfos) GetSummaryAppInfo() []*ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo {
	return s.SummaryAppInfo
}

func (s *ListSummaryAppsResponseBodySummaryAppInfos) SetSummaryAppInfo(v []*ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo) *ListSummaryAppsResponseBodySummaryAppInfos {
	s.SummaryAppInfo = v
	return s
}

func (s *ListSummaryAppsResponseBodySummaryAppInfos) Validate() error {
	if s.SummaryAppInfo != nil {
		for _, item := range s.SummaryAppInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo struct {
	// example:
	//
	// 23****07
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// abc
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo) GoString() string {
	return s.String()
}

func (s *ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo) GetAppKey() *int64 {
	return s.AppKey
}

func (s *ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo) GetAppName() *string {
	return s.AppName
}

func (s *ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo) SetAppKey(v int64) *ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo {
	s.AppKey = &v
	return s
}

func (s *ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo) SetAppName(v string) *ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo {
	s.AppName = &v
	return s
}

func (s *ListSummaryAppsResponseBodySummaryAppInfosSummaryAppInfo) Validate() error {
	return dara.Validate(s)
}
