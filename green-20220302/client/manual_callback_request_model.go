// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *ManualCallbackRequest
	GetChannel() *string
	SetChecksum(v string) *ManualCallbackRequest
	GetChecksum() *string
	SetCode(v string) *ManualCallbackRequest
	GetCode() *string
	SetData(v string) *ManualCallbackRequest
	GetData() *string
	SetMsg(v string) *ManualCallbackRequest
	GetMsg() *string
	SetReqId(v string) *ManualCallbackRequest
	GetReqId() *string
}

type ManualCallbackRequest struct {
	// Channel field
	//
	// example:
	//
	// ant
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// Checksum.
	//
	// example:
	//
	// abc
	Checksum *string `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	// Code value
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	//
	// example:
	//
	// {\\"Result\\": [{\\"Confidence\\": 100.0, \\"CustomizedHit\\": [{\\"KeyWords\\": u\\"\\u4fdd\\u969c,\\u6700\\u5927,\\u9ad8\\u7ea7\\", \\"LibName\\": u\\"\\u4f18\\u8def\\u654f\\u611f\\u8bcd\\"}], \\"Label\\": \\"customized\\"}]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Message information
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Platform request ID, used for troubleshooting assistance
	//
	// example:
	//
	// B0963D30-BAB4-562F-9ED0-7A23AEC51C7C
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s ManualCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s ManualCallbackRequest) GoString() string {
	return s.String()
}

func (s *ManualCallbackRequest) GetChannel() *string {
	return s.Channel
}

func (s *ManualCallbackRequest) GetChecksum() *string {
	return s.Checksum
}

func (s *ManualCallbackRequest) GetCode() *string {
	return s.Code
}

func (s *ManualCallbackRequest) GetData() *string {
	return s.Data
}

func (s *ManualCallbackRequest) GetMsg() *string {
	return s.Msg
}

func (s *ManualCallbackRequest) GetReqId() *string {
	return s.ReqId
}

func (s *ManualCallbackRequest) SetChannel(v string) *ManualCallbackRequest {
	s.Channel = &v
	return s
}

func (s *ManualCallbackRequest) SetChecksum(v string) *ManualCallbackRequest {
	s.Checksum = &v
	return s
}

func (s *ManualCallbackRequest) SetCode(v string) *ManualCallbackRequest {
	s.Code = &v
	return s
}

func (s *ManualCallbackRequest) SetData(v string) *ManualCallbackRequest {
	s.Data = &v
	return s
}

func (s *ManualCallbackRequest) SetMsg(v string) *ManualCallbackRequest {
	s.Msg = &v
	return s
}

func (s *ManualCallbackRequest) SetReqId(v string) *ManualCallbackRequest {
	s.ReqId = &v
	return s
}

func (s *ManualCallbackRequest) Validate() error {
	return dara.Validate(s)
}
