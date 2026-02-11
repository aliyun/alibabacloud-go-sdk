// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRetcodeAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRetcodeAppsResponseBody
	GetRequestId() *string
	SetRetcodeApps(v []*ListRetcodeAppsResponseBodyRetcodeApps) *ListRetcodeAppsResponseBody
	GetRetcodeApps() []*ListRetcodeAppsResponseBodyRetcodeApps
}

type ListRetcodeAppsResponseBody struct {
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetcodeApps []*ListRetcodeAppsResponseBodyRetcodeApps `json:"RetcodeApps,omitempty" xml:"RetcodeApps,omitempty" type:"Repeated"`
}

func (s ListRetcodeAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRetcodeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRetcodeAppsResponseBody) GetRetcodeApps() []*ListRetcodeAppsResponseBodyRetcodeApps {
	return s.RetcodeApps
}

func (s *ListRetcodeAppsResponseBody) SetRequestId(v string) *ListRetcodeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRetcodeAppsResponseBody) SetRetcodeApps(v []*ListRetcodeAppsResponseBodyRetcodeApps) *ListRetcodeAppsResponseBody {
	s.RetcodeApps = v
	return s
}

func (s *ListRetcodeAppsResponseBody) Validate() error {
	if s.RetcodeApps != nil {
		for _, item := range s.RetcodeApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRetcodeAppsResponseBodyRetcodeApps struct {
	AppId   *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Pid     *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s ListRetcodeAppsResponseBodyRetcodeApps) String() string {
	return dara.Prettify(s)
}

func (s ListRetcodeAppsResponseBodyRetcodeApps) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) GetAppId() *int64 {
	return s.AppId
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) GetAppName() *string {
	return s.AppName
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) GetPid() *string {
	return s.Pid
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetAppId(v int64) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.AppId = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetAppName(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.AppName = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetPid(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.Pid = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) Validate() error {
	return dara.Validate(s)
}
