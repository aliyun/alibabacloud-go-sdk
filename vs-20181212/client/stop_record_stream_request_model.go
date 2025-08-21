// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRecordStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *StopRecordStreamRequest
	GetApp() *string
	SetId(v string) *StopRecordStreamRequest
	GetId() *string
	SetName(v string) *StopRecordStreamRequest
	GetName() *string
	SetOwnerId(v int64) *StopRecordStreamRequest
	GetOwnerId() *int64
	SetPlayDomain(v string) *StopRecordStreamRequest
	GetPlayDomain() *string
}

type StopRecordStreamRequest struct {
	// example:
	//
	// live
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// example:
	//
	// 323*****997-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// example.com
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
}

func (s StopRecordStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRecordStreamRequest) GoString() string {
	return s.String()
}

func (s *StopRecordStreamRequest) GetApp() *string {
	return s.App
}

func (s *StopRecordStreamRequest) GetId() *string {
	return s.Id
}

func (s *StopRecordStreamRequest) GetName() *string {
	return s.Name
}

func (s *StopRecordStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopRecordStreamRequest) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *StopRecordStreamRequest) SetApp(v string) *StopRecordStreamRequest {
	s.App = &v
	return s
}

func (s *StopRecordStreamRequest) SetId(v string) *StopRecordStreamRequest {
	s.Id = &v
	return s
}

func (s *StopRecordStreamRequest) SetName(v string) *StopRecordStreamRequest {
	s.Name = &v
	return s
}

func (s *StopRecordStreamRequest) SetOwnerId(v int64) *StopRecordStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StopRecordStreamRequest) SetPlayDomain(v string) *StopRecordStreamRequest {
	s.PlayDomain = &v
	return s
}

func (s *StopRecordStreamRequest) Validate() error {
	return dara.Validate(s)
}
