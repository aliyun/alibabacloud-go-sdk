// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceEngineListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetLindormV2InstanceEngineListResponseBody
	GetAccessDeniedDetail() *string
	SetEngineList(v []*GetLindormV2InstanceEngineListResponseBodyEngineList) *GetLindormV2InstanceEngineListResponseBody
	GetEngineList() []*GetLindormV2InstanceEngineListResponseBodyEngineList
	SetInstanceId(v string) *GetLindormV2InstanceEngineListResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *GetLindormV2InstanceEngineListResponseBody
	GetRequestId() *string
}

type GetLindormV2InstanceEngineListResponseBody struct {
	AccessDeniedDetail *string                                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	EngineList         []*GetLindormV2InstanceEngineListResponseBodyEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	InstanceId         *string                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLindormV2InstanceEngineListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceEngineListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceEngineListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetLindormV2InstanceEngineListResponseBody) GetEngineList() []*GetLindormV2InstanceEngineListResponseBodyEngineList {
	return s.EngineList
}

func (s *GetLindormV2InstanceEngineListResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormV2InstanceEngineListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLindormV2InstanceEngineListResponseBody) SetAccessDeniedDetail(v string) *GetLindormV2InstanceEngineListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBody) SetEngineList(v []*GetLindormV2InstanceEngineListResponseBodyEngineList) *GetLindormV2InstanceEngineListResponseBody {
	s.EngineList = v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBody) SetInstanceId(v string) *GetLindormV2InstanceEngineListResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBody) SetRequestId(v string) *GetLindormV2InstanceEngineListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceEngineListResponseBodyEngineList struct {
	EngineType  *string                                                            `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	NetInfoList []*GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList `json:"NetInfoList,omitempty" xml:"NetInfoList,omitempty" type:"Repeated"`
}

func (s GetLindormV2InstanceEngineListResponseBodyEngineList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceEngineListResponseBodyEngineList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineList) GetEngineType() *string {
	return s.EngineType
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineList) GetNetInfoList() []*GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList {
	return s.NetInfoList
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineList) SetEngineType(v string) *GetLindormV2InstanceEngineListResponseBodyEngineList {
	s.EngineType = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineList) SetNetInfoList(v []*GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) *GetLindormV2InstanceEngineListResponseBodyEngineList {
	s.NetInfoList = v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineList) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList struct {
	AccessType       *int32  `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	NetType          *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) GetAccessType() *int32 {
	return s.AccessType
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) GetNetType() *string {
	return s.NetType
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) GetPort() *int32 {
	return s.Port
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) SetAccessType(v int32) *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList {
	s.AccessType = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) SetConnectionString(v string) *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList {
	s.ConnectionString = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) SetNetType(v string) *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList {
	s.NetType = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) SetPort(v int32) *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList {
	s.Port = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) Validate() error {
	return dara.Validate(s)
}
