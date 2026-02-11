// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDashboardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDashboardVos(v []*ListDashboardsResponseBodyDashboardVos) *ListDashboardsResponseBody
	GetDashboardVos() []*ListDashboardsResponseBodyDashboardVos
	SetRequestId(v string) *ListDashboardsResponseBody
	GetRequestId() *string
}

type ListDashboardsResponseBody struct {
	DashboardVos []*ListDashboardsResponseBodyDashboardVos `json:"DashboardVos,omitempty" xml:"DashboardVos,omitempty" type:"Repeated"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDashboardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDashboardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponseBody) GetDashboardVos() []*ListDashboardsResponseBodyDashboardVos {
	return s.DashboardVos
}

func (s *ListDashboardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDashboardsResponseBody) SetDashboardVos(v []*ListDashboardsResponseBodyDashboardVos) *ListDashboardsResponseBody {
	s.DashboardVos = v
	return s
}

func (s *ListDashboardsResponseBody) SetRequestId(v string) *ListDashboardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDashboardsResponseBody) Validate() error {
	if s.DashboardVos != nil {
		for _, item := range s.DashboardVos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDashboardsResponseBodyDashboardVos struct {
	DashboardType  *string   `json:"DashboardType,omitempty" xml:"DashboardType,omitempty"`
	Exporter       *string   `json:"Exporter,omitempty" xml:"Exporter,omitempty"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	IsArmsExporter *bool     `json:"IsArmsExporter,omitempty" xml:"IsArmsExporter,omitempty"`
	Kind           *string   `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedUpdate     *bool     `json:"NeedUpdate,omitempty" xml:"NeedUpdate,omitempty"`
	Tags           []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Time           *string   `json:"Time,omitempty" xml:"Time,omitempty"`
	Title          *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	Type           *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid            *string   `json:"Uid,omitempty" xml:"Uid,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Version        *string   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListDashboardsResponseBodyDashboardVos) String() string {
	return dara.Prettify(s)
}

func (s ListDashboardsResponseBodyDashboardVos) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponseBodyDashboardVos) GetDashboardType() *string {
	return s.DashboardType
}

func (s *ListDashboardsResponseBodyDashboardVos) GetExporter() *string {
	return s.Exporter
}

func (s *ListDashboardsResponseBodyDashboardVos) GetId() *string {
	return s.Id
}

func (s *ListDashboardsResponseBodyDashboardVos) GetIsArmsExporter() *bool {
	return s.IsArmsExporter
}

func (s *ListDashboardsResponseBodyDashboardVos) GetKind() *string {
	return s.Kind
}

func (s *ListDashboardsResponseBodyDashboardVos) GetName() *string {
	return s.Name
}

func (s *ListDashboardsResponseBodyDashboardVos) GetNeedUpdate() *bool {
	return s.NeedUpdate
}

func (s *ListDashboardsResponseBodyDashboardVos) GetTags() []*string {
	return s.Tags
}

func (s *ListDashboardsResponseBodyDashboardVos) GetTime() *string {
	return s.Time
}

func (s *ListDashboardsResponseBodyDashboardVos) GetTitle() *string {
	return s.Title
}

func (s *ListDashboardsResponseBodyDashboardVos) GetType() *string {
	return s.Type
}

func (s *ListDashboardsResponseBodyDashboardVos) GetUid() *string {
	return s.Uid
}

func (s *ListDashboardsResponseBodyDashboardVos) GetUrl() *string {
	return s.Url
}

func (s *ListDashboardsResponseBodyDashboardVos) GetVersion() *string {
	return s.Version
}

func (s *ListDashboardsResponseBodyDashboardVos) SetDashboardType(v string) *ListDashboardsResponseBodyDashboardVos {
	s.DashboardType = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetExporter(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Exporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetId(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Id = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetIsArmsExporter(v bool) *ListDashboardsResponseBodyDashboardVos {
	s.IsArmsExporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetKind(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Kind = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetName(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Name = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetNeedUpdate(v bool) *ListDashboardsResponseBodyDashboardVos {
	s.NeedUpdate = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTags(v []*string) *ListDashboardsResponseBodyDashboardVos {
	s.Tags = v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTime(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Time = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTitle(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Title = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetType(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Type = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetUid(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Uid = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetUrl(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Url = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetVersion(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Version = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) Validate() error {
	return dara.Validate(s)
}
