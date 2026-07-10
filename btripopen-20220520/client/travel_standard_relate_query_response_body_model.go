// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *TravelStandardRelateQueryResponseBody
	GetMessage() *string
	SetModule(v *TravelStandardRelateQueryResponseBodyModule) *TravelStandardRelateQueryResponseBody
	GetModule() *TravelStandardRelateQueryResponseBodyModule
	SetRequestId(v string) *TravelStandardRelateQueryResponseBody
	GetRequestId() *string
	SetResultCode(v int32) *TravelStandardRelateQueryResponseBody
	GetResultCode() *int32
	SetSuccess(v bool) *TravelStandardRelateQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TravelStandardRelateQueryResponseBody
	GetTraceId() *string
}

type TravelStandardRelateQueryResponseBody struct {
	Message    *string                                      `json:"message,omitempty" xml:"message,omitempty"`
	Module     *TravelStandardRelateQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId  *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultCode *int32                                       `json:"resultCode,omitempty" xml:"resultCode,omitempty"`
	Success    *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
	TraceId    *string                                      `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TravelStandardRelateQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TravelStandardRelateQueryResponseBody) GetModule() *TravelStandardRelateQueryResponseBodyModule {
	return s.Module
}

func (s *TravelStandardRelateQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TravelStandardRelateQueryResponseBody) GetResultCode() *int32 {
	return s.ResultCode
}

func (s *TravelStandardRelateQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TravelStandardRelateQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TravelStandardRelateQueryResponseBody) SetMessage(v string) *TravelStandardRelateQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TravelStandardRelateQueryResponseBody) SetModule(v *TravelStandardRelateQueryResponseBodyModule) *TravelStandardRelateQueryResponseBody {
	s.Module = v
	return s
}

func (s *TravelStandardRelateQueryResponseBody) SetRequestId(v string) *TravelStandardRelateQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TravelStandardRelateQueryResponseBody) SetResultCode(v int32) *TravelStandardRelateQueryResponseBody {
	s.ResultCode = &v
	return s
}

func (s *TravelStandardRelateQueryResponseBody) SetSuccess(v bool) *TravelStandardRelateQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TravelStandardRelateQueryResponseBody) SetTraceId(v string) *TravelStandardRelateQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TravelStandardRelateQueryResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TravelStandardRelateQueryResponseBodyModule struct {
	ReserveBindEntityList []*TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList `json:"reserve_bind_entity_list,omitempty" xml:"reserve_bind_entity_list,omitempty" type:"Repeated"`
	Total                 *int32                                                              `json:"total,omitempty" xml:"total,omitempty"`
}

func (s TravelStandardRelateQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateQueryResponseBodyModule) GetReserveBindEntityList() []*TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList {
	return s.ReserveBindEntityList
}

func (s *TravelStandardRelateQueryResponseBodyModule) GetTotal() *int32 {
	return s.Total
}

func (s *TravelStandardRelateQueryResponseBodyModule) SetReserveBindEntityList(v []*TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList) *TravelStandardRelateQueryResponseBodyModule {
	s.ReserveBindEntityList = v
	return s
}

func (s *TravelStandardRelateQueryResponseBodyModule) SetTotal(v int32) *TravelStandardRelateQueryResponseBodyModule {
	s.Total = &v
	return s
}

func (s *TravelStandardRelateQueryResponseBodyModule) Validate() error {
	if s.ReserveBindEntityList != nil {
		for _, item := range s.ReserveBindEntityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList struct {
	EntityId   *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	EntityName *string `json:"entity_name,omitempty" xml:"entity_name,omitempty"`
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
}

func (s TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList) GetEntityId() *string {
	return s.EntityId
}

func (s *TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList) GetEntityName() *string {
	return s.EntityName
}

func (s *TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList) GetEntityType() *string {
	return s.EntityType
}

func (s *TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList) SetEntityId(v string) *TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList {
	s.EntityId = &v
	return s
}

func (s *TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList) SetEntityName(v string) *TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList {
	s.EntityName = &v
	return s
}

func (s *TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList) SetEntityType(v string) *TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList {
	s.EntityType = &v
	return s
}

func (s *TravelStandardRelateQueryResponseBodyModuleReserveBindEntityList) Validate() error {
	return dara.Validate(s)
}
