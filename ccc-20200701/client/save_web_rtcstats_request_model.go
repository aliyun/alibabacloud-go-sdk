// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWebRTCStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *SaveWebRTCStatsRequest
	GetCallId() *string
	SetGeneralInfo(v string) *SaveWebRTCStatsRequest
	GetGeneralInfo() *string
	SetGoogAddress(v string) *SaveWebRTCStatsRequest
	GetGoogAddress() *string
	SetInstanceId(v string) *SaveWebRTCStatsRequest
	GetInstanceId() *string
	SetReceiverReport(v string) *SaveWebRTCStatsRequest
	GetReceiverReport() *string
	SetSenderReport(v string) *SaveWebRTCStatsRequest
	GetSenderReport() *string
}

type SaveWebRTCStatsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 257e73de-1ee8-123b-0b9a-00163e0a****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	GeneralInfo *string `json:"GeneralInfo,omitempty" xml:"GeneralInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["47.101.XX.XX","44368","47.104.XX.XX","37947"]
	GoogAddress *string `json:"GoogAddress,omitempty" xml:"GoogAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["1","1649328987","40","PCMU","383560","89","49","ssrc_1649328987_recv","2022-03-15T09:52:08.820Z","","0","20000"]
	ReceiverReport *string `json:"ReceiverReport,omitempty" xml:"ReceiverReport,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["11090","2055127460","0","2236","384592","ssrc_2055127460_send","2022-03-15T09:52:08.820Z","1"]
	SenderReport *string `json:"SenderReport,omitempty" xml:"SenderReport,omitempty"`
}

func (s SaveWebRTCStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveWebRTCStatsRequest) GoString() string {
	return s.String()
}

func (s *SaveWebRTCStatsRequest) GetCallId() *string {
	return s.CallId
}

func (s *SaveWebRTCStatsRequest) GetGeneralInfo() *string {
	return s.GeneralInfo
}

func (s *SaveWebRTCStatsRequest) GetGoogAddress() *string {
	return s.GoogAddress
}

func (s *SaveWebRTCStatsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveWebRTCStatsRequest) GetReceiverReport() *string {
	return s.ReceiverReport
}

func (s *SaveWebRTCStatsRequest) GetSenderReport() *string {
	return s.SenderReport
}

func (s *SaveWebRTCStatsRequest) SetCallId(v string) *SaveWebRTCStatsRequest {
	s.CallId = &v
	return s
}

func (s *SaveWebRTCStatsRequest) SetGeneralInfo(v string) *SaveWebRTCStatsRequest {
	s.GeneralInfo = &v
	return s
}

func (s *SaveWebRTCStatsRequest) SetGoogAddress(v string) *SaveWebRTCStatsRequest {
	s.GoogAddress = &v
	return s
}

func (s *SaveWebRTCStatsRequest) SetInstanceId(v string) *SaveWebRTCStatsRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveWebRTCStatsRequest) SetReceiverReport(v string) *SaveWebRTCStatsRequest {
	s.ReceiverReport = &v
	return s
}

func (s *SaveWebRTCStatsRequest) SetSenderReport(v string) *SaveWebRTCStatsRequest {
	s.SenderReport = &v
	return s
}

func (s *SaveWebRTCStatsRequest) Validate() error {
	return dara.Validate(s)
}
