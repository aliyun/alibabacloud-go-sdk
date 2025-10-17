// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSaasPermissionGroupInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListSaasPermissionGroupInfosResponseBodyData) *ListSaasPermissionGroupInfosResponseBody
	GetData() []*ListSaasPermissionGroupInfosResponseBodyData
	SetRequestId(v string) *ListSaasPermissionGroupInfosResponseBody
	GetRequestId() *string
}

type ListSaasPermissionGroupInfosResponseBody struct {
	Data []*ListSaasPermissionGroupInfosResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 8AD9FA10-7780-5E12-B701-13C928524F32
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSaasPermissionGroupInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSaasPermissionGroupInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListSaasPermissionGroupInfosResponseBody) GetData() []*ListSaasPermissionGroupInfosResponseBodyData {
	return s.Data
}

func (s *ListSaasPermissionGroupInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSaasPermissionGroupInfosResponseBody) SetData(v []*ListSaasPermissionGroupInfosResponseBodyData) *ListSaasPermissionGroupInfosResponseBody {
	s.Data = v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBody) SetRequestId(v string) *ListSaasPermissionGroupInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBody) Validate() error {
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

type ListSaasPermissionGroupInfosResponseBodyData struct {
	// example:
	//
	// Release Center
	EnName  *string                                                `json:"EnName,omitempty" xml:"EnName,omitempty"`
	Name    *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	PgInfos []*ListSaasPermissionGroupInfosResponseBodyDataPgInfos `json:"PgInfos,omitempty" xml:"PgInfos,omitempty" type:"Repeated"`
	// example:
	//
	// FAQ
	SaasCode *string `json:"SaasCode,omitempty" xml:"SaasCode,omitempty"`
}

func (s ListSaasPermissionGroupInfosResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSaasPermissionGroupInfosResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSaasPermissionGroupInfosResponseBodyData) GetEnName() *string {
	return s.EnName
}

func (s *ListSaasPermissionGroupInfosResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListSaasPermissionGroupInfosResponseBodyData) GetPgInfos() []*ListSaasPermissionGroupInfosResponseBodyDataPgInfos {
	return s.PgInfos
}

func (s *ListSaasPermissionGroupInfosResponseBodyData) GetSaasCode() *string {
	return s.SaasCode
}

func (s *ListSaasPermissionGroupInfosResponseBodyData) SetEnName(v string) *ListSaasPermissionGroupInfosResponseBodyData {
	s.EnName = &v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBodyData) SetName(v string) *ListSaasPermissionGroupInfosResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBodyData) SetPgInfos(v []*ListSaasPermissionGroupInfosResponseBodyDataPgInfos) *ListSaasPermissionGroupInfosResponseBodyData {
	s.PgInfos = v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBodyData) SetSaasCode(v string) *ListSaasPermissionGroupInfosResponseBodyData {
	s.SaasCode = &v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBodyData) Validate() error {
	if s.PgInfos != nil {
		for _, item := range s.PgInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSaasPermissionGroupInfosResponseBodyDataPgInfos struct {
	// example:
	//
	// FAQ
	PgCode *string `json:"PgCode,omitempty" xml:"PgCode,omitempty"`
	// example:
	//
	// FAQ
	PgEnName *string `json:"PgEnName,omitempty" xml:"PgEnName,omitempty"`
	PgName   *string `json:"PgName,omitempty" xml:"PgName,omitempty"`
}

func (s ListSaasPermissionGroupInfosResponseBodyDataPgInfos) String() string {
	return dara.Prettify(s)
}

func (s ListSaasPermissionGroupInfosResponseBodyDataPgInfos) GoString() string {
	return s.String()
}

func (s *ListSaasPermissionGroupInfosResponseBodyDataPgInfos) GetPgCode() *string {
	return s.PgCode
}

func (s *ListSaasPermissionGroupInfosResponseBodyDataPgInfos) GetPgEnName() *string {
	return s.PgEnName
}

func (s *ListSaasPermissionGroupInfosResponseBodyDataPgInfos) GetPgName() *string {
	return s.PgName
}

func (s *ListSaasPermissionGroupInfosResponseBodyDataPgInfos) SetPgCode(v string) *ListSaasPermissionGroupInfosResponseBodyDataPgInfos {
	s.PgCode = &v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBodyDataPgInfos) SetPgEnName(v string) *ListSaasPermissionGroupInfosResponseBodyDataPgInfos {
	s.PgEnName = &v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBodyDataPgInfos) SetPgName(v string) *ListSaasPermissionGroupInfosResponseBodyDataPgInfos {
	s.PgName = &v
	return s
}

func (s *ListSaasPermissionGroupInfosResponseBodyDataPgInfos) Validate() error {
	return dara.Validate(s)
}
