// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncAppInstanceForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SyncAppInstanceForPartnerResponseBodyData) *SyncAppInstanceForPartnerResponseBody
	GetData() *SyncAppInstanceForPartnerResponseBodyData
	SetRequestId(v string) *SyncAppInstanceForPartnerResponseBody
	GetRequestId() *string
}

type SyncAppInstanceForPartnerResponseBody struct {
	Data *SyncAppInstanceForPartnerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncAppInstanceForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncAppInstanceForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *SyncAppInstanceForPartnerResponseBody) GetData() *SyncAppInstanceForPartnerResponseBodyData {
	return s.Data
}

func (s *SyncAppInstanceForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncAppInstanceForPartnerResponseBody) SetData(v *SyncAppInstanceForPartnerResponseBodyData) *SyncAppInstanceForPartnerResponseBody {
	s.Data = v
	return s
}

func (s *SyncAppInstanceForPartnerResponseBody) SetRequestId(v string) *SyncAppInstanceForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncAppInstanceForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}

type SyncAppInstanceForPartnerResponseBodyData struct {
	AppInstance *SyncAppInstanceForPartnerResponseBodyDataAppInstance `json:"AppInstance,omitempty" xml:"AppInstance,omitempty" type:"Struct"`
}

func (s SyncAppInstanceForPartnerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SyncAppInstanceForPartnerResponseBodyData) GoString() string {
	return s.String()
}

func (s *SyncAppInstanceForPartnerResponseBodyData) GetAppInstance() *SyncAppInstanceForPartnerResponseBodyDataAppInstance {
	return s.AppInstance
}

func (s *SyncAppInstanceForPartnerResponseBodyData) SetAppInstance(v *SyncAppInstanceForPartnerResponseBodyDataAppInstance) *SyncAppInstanceForPartnerResponseBodyData {
	s.AppInstance = v
	return s
}

func (s *SyncAppInstanceForPartnerResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type SyncAppInstanceForPartnerResponseBodyDataAppInstance struct {
	// example:
	//
	// WD20250626114752000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s SyncAppInstanceForPartnerResponseBodyDataAppInstance) String() string {
	return dara.Prettify(s)
}

func (s SyncAppInstanceForPartnerResponseBodyDataAppInstance) GoString() string {
	return s.String()
}

func (s *SyncAppInstanceForPartnerResponseBodyDataAppInstance) GetBizId() *string {
	return s.BizId
}

func (s *SyncAppInstanceForPartnerResponseBodyDataAppInstance) SetBizId(v string) *SyncAppInstanceForPartnerResponseBodyDataAppInstance {
	s.BizId = &v
	return s
}

func (s *SyncAppInstanceForPartnerResponseBodyDataAppInstance) Validate() error {
	return dara.Validate(s)
}
