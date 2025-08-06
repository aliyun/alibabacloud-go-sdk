// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPersonalStorageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetPersonalStorageListResponseBody
	GetRequestId() *string
	SetStorageInfoList(v *GetPersonalStorageListResponseBodyStorageInfoList) *GetPersonalStorageListResponseBody
	GetStorageInfoList() *GetPersonalStorageListResponseBodyStorageInfoList
}

type GetPersonalStorageListResponseBody struct {
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageInfoList *GetPersonalStorageListResponseBodyStorageInfoList `json:"StorageInfoList,omitempty" xml:"StorageInfoList,omitempty" type:"Struct"`
}

func (s GetPersonalStorageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPersonalStorageListResponseBody) GoString() string {
	return s.String()
}

func (s *GetPersonalStorageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPersonalStorageListResponseBody) GetStorageInfoList() *GetPersonalStorageListResponseBodyStorageInfoList {
	return s.StorageInfoList
}

func (s *GetPersonalStorageListResponseBody) SetRequestId(v string) *GetPersonalStorageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPersonalStorageListResponseBody) SetStorageInfoList(v *GetPersonalStorageListResponseBodyStorageInfoList) *GetPersonalStorageListResponseBody {
	s.StorageInfoList = v
	return s
}

func (s *GetPersonalStorageListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPersonalStorageListResponseBodyStorageInfoList struct {
	StorageInfo []*GetPersonalStorageListResponseBodyStorageInfoListStorageInfo `json:"StorageInfo,omitempty" xml:"StorageInfo,omitempty" type:"Repeated"`
}

func (s GetPersonalStorageListResponseBodyStorageInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPersonalStorageListResponseBodyStorageInfoList) GoString() string {
	return s.String()
}

func (s *GetPersonalStorageListResponseBodyStorageInfoList) GetStorageInfo() []*GetPersonalStorageListResponseBodyStorageInfoListStorageInfo {
	return s.StorageInfo
}

func (s *GetPersonalStorageListResponseBodyStorageInfoList) SetStorageInfo(v []*GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) *GetPersonalStorageListResponseBodyStorageInfoList {
	s.StorageInfo = v
	return s
}

func (s *GetPersonalStorageListResponseBodyStorageInfoList) Validate() error {
	return dara.Validate(s)
}

type GetPersonalStorageListResponseBodyStorageInfoListStorageInfo struct {
	ExtranetEndpoint *string `json:"ExtranetEndpoint,omitempty" xml:"ExtranetEndpoint,omitempty"`
	GmtCreate        *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	Location         *string `json:"Location,omitempty" xml:"Location,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StorageACL       *string `json:"StorageACL,omitempty" xml:"StorageACL,omitempty"`
	StorageClass     *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) GoString() string {
	return s.String()
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) GetExtranetEndpoint() *string {
	return s.ExtranetEndpoint
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) GetIntranetEndpoint() *string {
	return s.IntranetEndpoint
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) GetLocation() *string {
	return s.Location
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) GetStorageACL() *string {
	return s.StorageACL
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) GetStorageClass() *string {
	return s.StorageClass
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) SetExtranetEndpoint(v string) *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo {
	s.ExtranetEndpoint = &v
	return s
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) SetGmtCreate(v string) *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) SetIntranetEndpoint(v string) *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo {
	s.IntranetEndpoint = &v
	return s
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) SetLocation(v string) *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo {
	s.Location = &v
	return s
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) SetOwnerId(v int64) *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo {
	s.OwnerId = &v
	return s
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) SetStorageACL(v string) *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo {
	s.StorageACL = &v
	return s
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) SetStorageClass(v string) *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo {
	s.StorageClass = &v
	return s
}

func (s *GetPersonalStorageListResponseBodyStorageInfoListStorageInfo) Validate() error {
	return dara.Validate(s)
}
