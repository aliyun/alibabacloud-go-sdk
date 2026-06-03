// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAnnotationMissionTagInfoListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAnnotationMissionTagInfoListResponseBody
	GetCode() *string
	SetData(v *GetAnnotationMissionTagInfoListResponseBodyData) *GetAnnotationMissionTagInfoListResponseBody
	GetData() *GetAnnotationMissionTagInfoListResponseBodyData
	SetHttpStatusCode(v int32) *GetAnnotationMissionTagInfoListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAnnotationMissionTagInfoListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAnnotationMissionTagInfoListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAnnotationMissionTagInfoListResponseBody
	GetSuccess() *bool
}

type GetAnnotationMissionTagInfoListResponseBody struct {
	// example:
	//
	// OK
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAnnotationMissionTagInfoListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// CDR \\"job-efbaeefc-4d45-4e79-83f7-b33b0769c969\\" doesn\\"t exists.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAnnotationMissionTagInfoListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAnnotationMissionTagInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAnnotationMissionTagInfoListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAnnotationMissionTagInfoListResponseBody) GetData() *GetAnnotationMissionTagInfoListResponseBodyData {
	return s.Data
}

func (s *GetAnnotationMissionTagInfoListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAnnotationMissionTagInfoListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAnnotationMissionTagInfoListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAnnotationMissionTagInfoListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAnnotationMissionTagInfoListResponseBody) SetCode(v string) *GetAnnotationMissionTagInfoListResponseBody {
	s.Code = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBody) SetData(v *GetAnnotationMissionTagInfoListResponseBodyData) *GetAnnotationMissionTagInfoListResponseBody {
	s.Data = v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBody) SetHttpStatusCode(v int32) *GetAnnotationMissionTagInfoListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBody) SetMessage(v string) *GetAnnotationMissionTagInfoListResponseBody {
	s.Message = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBody) SetRequestId(v string) *GetAnnotationMissionTagInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBody) SetSuccess(v bool) *GetAnnotationMissionTagInfoListResponseBody {
	s.Success = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAnnotationMissionTagInfoListResponseBodyData struct {
	AnnotationMissionTagInfoList []*GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList `json:"AnnotationMissionTagInfoList,omitempty" xml:"AnnotationMissionTagInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// CDR \\"job-efbaeefc-4d45-4e79-83f7-b33b0769c969\\" doesn\\"t exists.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAnnotationMissionTagInfoListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAnnotationMissionTagInfoListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAnnotationMissionTagInfoListResponseBodyData) GetAnnotationMissionTagInfoList() []*GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList {
	return s.AnnotationMissionTagInfoList
}

func (s *GetAnnotationMissionTagInfoListResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetAnnotationMissionTagInfoListResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *GetAnnotationMissionTagInfoListResponseBodyData) SetAnnotationMissionTagInfoList(v []*GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) *GetAnnotationMissionTagInfoListResponseBodyData {
	s.AnnotationMissionTagInfoList = v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBodyData) SetMessage(v string) *GetAnnotationMissionTagInfoListResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBodyData) SetSuccess(v bool) *GetAnnotationMissionTagInfoListResponseBodyData {
	s.Success = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBodyData) Validate() error {
	if s.AnnotationMissionTagInfoList != nil {
		for _, item := range s.AnnotationMissionTagInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList struct {
	AnnotationMissionTagInfoDescription *string `json:"AnnotationMissionTagInfoDescription,omitempty" xml:"AnnotationMissionTagInfoDescription,omitempty"`
	// example:
	//
	// e1ee87ea-ebad-4079-aebb-1c56a4ef0c06
	AnnotationMissionTagInfoId   *string `json:"AnnotationMissionTagInfoId,omitempty" xml:"AnnotationMissionTagInfoId,omitempty"`
	AnnotationMissionTagInfoName *string `json:"AnnotationMissionTagInfoName,omitempty" xml:"AnnotationMissionTagInfoName,omitempty"`
	// example:
	//
	// false
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// example:
	//
	// 3c3b8d1d-deff-48d9-9318-addc80ae5b1e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// DING_ORG_dingbd9daecdb7aaed3bffe93478753d9884
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) GoString() string {
	return s.String()
}

func (s *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) GetAnnotationMissionTagInfoDescription() *string {
	return s.AnnotationMissionTagInfoDescription
}

func (s *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) GetAnnotationMissionTagInfoId() *string {
	return s.AnnotationMissionTagInfoId
}

func (s *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) GetAnnotationMissionTagInfoName() *string {
	return s.AnnotationMissionTagInfoName
}

func (s *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) GetDelete() *bool {
	return s.Delete
}

func (s *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) SetAnnotationMissionTagInfoDescription(v string) *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList {
	s.AnnotationMissionTagInfoDescription = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) SetAnnotationMissionTagInfoId(v string) *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList {
	s.AnnotationMissionTagInfoId = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) SetAnnotationMissionTagInfoName(v string) *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList {
	s.AnnotationMissionTagInfoName = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) SetDelete(v bool) *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList {
	s.Delete = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) SetInstanceId(v string) *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList {
	s.InstanceId = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) SetTenantId(v string) *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList {
	s.TenantId = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListResponseBodyDataAnnotationMissionTagInfoList) Validate() error {
	return dara.Validate(s)
}
