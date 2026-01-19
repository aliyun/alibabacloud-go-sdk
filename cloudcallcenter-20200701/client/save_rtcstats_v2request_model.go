// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveRTCStatsV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *SaveRTCStatsV2Request
	GetCallId() *string
	SetGeneralInfo(v string) *SaveRTCStatsV2Request
	GetGeneralInfo() *string
	SetGoogAddress(v string) *SaveRTCStatsV2Request
	GetGoogAddress() *string
	SetInstanceId(v string) *SaveRTCStatsV2Request
	GetInstanceId() *string
	SetReceiverReport(v string) *SaveRTCStatsV2Request
	GetReceiverReport() *string
	SetSenderReport(v string) *SaveRTCStatsV2Request
	GetSenderReport() *string
}

type SaveRTCStatsV2Request struct {
	// This parameter is required.
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	GeneralInfo *string `json:"GeneralInfo,omitempty" xml:"GeneralInfo,omitempty"`
	// This parameter is required.
	GoogAddress *string `json:"GoogAddress,omitempty" xml:"GoogAddress,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	ReceiverReport *string `json:"ReceiverReport,omitempty" xml:"ReceiverReport,omitempty"`
	// This parameter is required.
	SenderReport *string `json:"SenderReport,omitempty" xml:"SenderReport,omitempty"`
}

func (s SaveRTCStatsV2Request) String() string {
	return dara.Prettify(s)
}

func (s SaveRTCStatsV2Request) GoString() string {
	return s.String()
}

func (s *SaveRTCStatsV2Request) GetCallId() *string {
	return s.CallId
}

func (s *SaveRTCStatsV2Request) GetGeneralInfo() *string {
	return s.GeneralInfo
}

func (s *SaveRTCStatsV2Request) GetGoogAddress() *string {
	return s.GoogAddress
}

func (s *SaveRTCStatsV2Request) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveRTCStatsV2Request) GetReceiverReport() *string {
	return s.ReceiverReport
}

func (s *SaveRTCStatsV2Request) GetSenderReport() *string {
	return s.SenderReport
}

func (s *SaveRTCStatsV2Request) SetCallId(v string) *SaveRTCStatsV2Request {
	s.CallId = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetGeneralInfo(v string) *SaveRTCStatsV2Request {
	s.GeneralInfo = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetGoogAddress(v string) *SaveRTCStatsV2Request {
	s.GoogAddress = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetInstanceId(v string) *SaveRTCStatsV2Request {
	s.InstanceId = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetReceiverReport(v string) *SaveRTCStatsV2Request {
	s.ReceiverReport = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetSenderReport(v string) *SaveRTCStatsV2Request {
	s.SenderReport = &v
	return s
}

func (s *SaveRTCStatsV2Request) Validate() error {
	return dara.Validate(s)
}
