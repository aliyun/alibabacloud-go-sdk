// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTrancodeSEIRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AddTrancodeSEIRequest
	GetAppName() *string
	SetDelay(v int32) *AddTrancodeSEIRequest
	GetDelay() *int32
	SetDomainName(v string) *AddTrancodeSEIRequest
	GetDomainName() *string
	SetOwnerId(v int64) *AddTrancodeSEIRequest
	GetOwnerId() *int64
	SetPattern(v string) *AddTrancodeSEIRequest
	GetPattern() *string
	SetRegionId(v string) *AddTrancodeSEIRequest
	GetRegionId() *string
	SetRepeat(v int32) *AddTrancodeSEIRequest
	GetRepeat() *int32
	SetStreamName(v string) *AddTrancodeSEIRequest
	GetStreamName() *string
	SetText(v string) *AddTrancodeSEIRequest
	GetText() *string
}

type AddTrancodeSEIRequest struct {
	// The AppName of the live stream. View AppNames on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The delay in milliseconds before the SEI is inserted after the command is received.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Delay *int32 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies the insertion pattern for the SEI.
	//
	// - **keyframe**: Inserts at every keyframe.
	//
	// - **frame**: Inserts at every single frame.
	//
	// This parameter is required.
	//
	// example:
	//
	// keyframe
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of times to repeat the insertion. A value of -1 means infinite repetitions.
	//
	// This parameter is required.
	//
	// example:
	//
	// -1
	Repeat *int32 `json:"Repeat,omitempty" xml:"Repeat,omitempty"`
	// The name of the live stream.
	//
	// > It must be the name of the source stream. This ensures that SEI is inserted into all transcoded streams.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The SEI text. Length limit: 4000 bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveSei****
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s AddTrancodeSEIRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTrancodeSEIRequest) GoString() string {
	return s.String()
}

func (s *AddTrancodeSEIRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddTrancodeSEIRequest) GetDelay() *int32 {
	return s.Delay
}

func (s *AddTrancodeSEIRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddTrancodeSEIRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddTrancodeSEIRequest) GetPattern() *string {
	return s.Pattern
}

func (s *AddTrancodeSEIRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddTrancodeSEIRequest) GetRepeat() *int32 {
	return s.Repeat
}

func (s *AddTrancodeSEIRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *AddTrancodeSEIRequest) GetText() *string {
	return s.Text
}

func (s *AddTrancodeSEIRequest) SetAppName(v string) *AddTrancodeSEIRequest {
	s.AppName = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetDelay(v int32) *AddTrancodeSEIRequest {
	s.Delay = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetDomainName(v string) *AddTrancodeSEIRequest {
	s.DomainName = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetOwnerId(v int64) *AddTrancodeSEIRequest {
	s.OwnerId = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetPattern(v string) *AddTrancodeSEIRequest {
	s.Pattern = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetRegionId(v string) *AddTrancodeSEIRequest {
	s.RegionId = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetRepeat(v int32) *AddTrancodeSEIRequest {
	s.Repeat = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetStreamName(v string) *AddTrancodeSEIRequest {
	s.StreamName = &v
	return s
}

func (s *AddTrancodeSEIRequest) SetText(v string) *AddTrancodeSEIRequest {
	s.Text = &v
	return s
}

func (s *AddTrancodeSEIRequest) Validate() error {
	return dara.Validate(s)
}
