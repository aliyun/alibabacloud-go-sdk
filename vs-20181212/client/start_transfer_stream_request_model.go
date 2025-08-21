// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTransferStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StartTransferStreamRequest
	GetId() *string
	SetOwnerId(v int64) *StartTransferStreamRequest
	GetOwnerId() *int64
	SetTranscode(v string) *StartTransferStreamRequest
	GetTranscode() *string
	SetUrl(v string) *StartTransferStreamRequest
	GetUrl() *string
}

type StartTransferStreamRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 323*****998-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// sd
	Transcode *string `json:"Transcode,omitempty" xml:"Transcode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s StartTransferStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTransferStreamRequest) GoString() string {
	return s.String()
}

func (s *StartTransferStreamRequest) GetId() *string {
	return s.Id
}

func (s *StartTransferStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartTransferStreamRequest) GetTranscode() *string {
	return s.Transcode
}

func (s *StartTransferStreamRequest) GetUrl() *string {
	return s.Url
}

func (s *StartTransferStreamRequest) SetId(v string) *StartTransferStreamRequest {
	s.Id = &v
	return s
}

func (s *StartTransferStreamRequest) SetOwnerId(v int64) *StartTransferStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StartTransferStreamRequest) SetTranscode(v string) *StartTransferStreamRequest {
	s.Transcode = &v
	return s
}

func (s *StartTransferStreamRequest) SetUrl(v string) *StartTransferStreamRequest {
	s.Url = &v
	return s
}

func (s *StartTransferStreamRequest) Validate() error {
	return dara.Validate(s)
}
