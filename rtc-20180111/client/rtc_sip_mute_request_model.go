// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRtcSipMuteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RtcSipMuteRequest
	GetAppId() *string
	SetChannelId(v string) *RtcSipMuteRequest
	GetChannelId() *string
	SetOperations(v []*RtcSipMuteRequestOperations) *RtcSipMuteRequest
	GetOperations() []*RtcSipMuteRequestOperations
}

type RtcSipMuteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	Operations []*RtcSipMuteRequestOperations `json:"Operations,omitempty" xml:"Operations,omitempty" type:"Repeated"`
}

func (s RtcSipMuteRequest) String() string {
	return dara.Prettify(s)
}

func (s RtcSipMuteRequest) GoString() string {
	return s.String()
}

func (s *RtcSipMuteRequest) GetAppId() *string {
	return s.AppId
}

func (s *RtcSipMuteRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *RtcSipMuteRequest) GetOperations() []*RtcSipMuteRequestOperations {
	return s.Operations
}

func (s *RtcSipMuteRequest) SetAppId(v string) *RtcSipMuteRequest {
	s.AppId = &v
	return s
}

func (s *RtcSipMuteRequest) SetChannelId(v string) *RtcSipMuteRequest {
	s.ChannelId = &v
	return s
}

func (s *RtcSipMuteRequest) SetOperations(v []*RtcSipMuteRequestOperations) *RtcSipMuteRequest {
	s.Operations = v
	return s
}

func (s *RtcSipMuteRequest) Validate() error {
	if s.Operations != nil {
		for _, item := range s.Operations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RtcSipMuteRequestOperations struct {
	// This parameter is required.
	Ids []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// replace
	Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12122121
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /media/audio/status
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	Value *RtcSipMuteRequestOperationsValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s RtcSipMuteRequestOperations) String() string {
	return dara.Prettify(s)
}

func (s RtcSipMuteRequestOperations) GoString() string {
	return s.String()
}

func (s *RtcSipMuteRequestOperations) GetIds() []*string {
	return s.Ids
}

func (s *RtcSipMuteRequestOperations) GetOp() *string {
	return s.Op
}

func (s *RtcSipMuteRequestOperations) GetOperationId() *string {
	return s.OperationId
}

func (s *RtcSipMuteRequestOperations) GetPath() *string {
	return s.Path
}

func (s *RtcSipMuteRequestOperations) GetValue() *RtcSipMuteRequestOperationsValue {
	return s.Value
}

func (s *RtcSipMuteRequestOperations) SetIds(v []*string) *RtcSipMuteRequestOperations {
	s.Ids = v
	return s
}

func (s *RtcSipMuteRequestOperations) SetOp(v string) *RtcSipMuteRequestOperations {
	s.Op = &v
	return s
}

func (s *RtcSipMuteRequestOperations) SetOperationId(v string) *RtcSipMuteRequestOperations {
	s.OperationId = &v
	return s
}

func (s *RtcSipMuteRequestOperations) SetPath(v string) *RtcSipMuteRequestOperations {
	s.Path = &v
	return s
}

func (s *RtcSipMuteRequestOperations) SetValue(v *RtcSipMuteRequestOperationsValue) *RtcSipMuteRequestOperations {
	s.Value = v
	return s
}

func (s *RtcSipMuteRequestOperations) Validate() error {
	if s.Value != nil {
		if err := s.Value.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RtcSipMuteRequestOperationsValue struct {
	// This parameter is required.
	//
	// example:
	//
	// inactive
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RtcSipMuteRequestOperationsValue) String() string {
	return dara.Prettify(s)
}

func (s RtcSipMuteRequestOperationsValue) GoString() string {
	return s.String()
}

func (s *RtcSipMuteRequestOperationsValue) GetStatus() *string {
	return s.Status
}

func (s *RtcSipMuteRequestOperationsValue) SetStatus(v string) *RtcSipMuteRequestOperationsValue {
	s.Status = &v
	return s
}

func (s *RtcSipMuteRequestOperationsValue) Validate() error {
	return dara.Validate(s)
}
