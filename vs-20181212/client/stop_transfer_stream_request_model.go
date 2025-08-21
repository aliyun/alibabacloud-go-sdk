// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTransferStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StopTransferStreamRequest
	GetId() *string
	SetOwnerId(v int64) *StopTransferStreamRequest
	GetOwnerId() *int64
	SetTranscode(v string) *StopTransferStreamRequest
	GetTranscode() *string
}

type StopTransferStreamRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 323434****83423432
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// sd
	Transcode *string `json:"Transcode,omitempty" xml:"Transcode,omitempty"`
}

func (s StopTransferStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTransferStreamRequest) GoString() string {
	return s.String()
}

func (s *StopTransferStreamRequest) GetId() *string {
	return s.Id
}

func (s *StopTransferStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopTransferStreamRequest) GetTranscode() *string {
	return s.Transcode
}

func (s *StopTransferStreamRequest) SetId(v string) *StopTransferStreamRequest {
	s.Id = &v
	return s
}

func (s *StopTransferStreamRequest) SetOwnerId(v int64) *StopTransferStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StopTransferStreamRequest) SetTranscode(v string) *StopTransferStreamRequest {
	s.Transcode = &v
	return s
}

func (s *StopTransferStreamRequest) Validate() error {
	return dara.Validate(s)
}
