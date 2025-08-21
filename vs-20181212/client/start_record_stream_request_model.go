// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRecordStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *StartRecordStreamRequest
	GetApp() *string
	SetId(v string) *StartRecordStreamRequest
	GetId() *string
	SetName(v string) *StartRecordStreamRequest
	GetName() *string
	SetOwnerId(v int64) *StartRecordStreamRequest
	GetOwnerId() *int64
	SetPlayDomain(v string) *StartRecordStreamRequest
	GetPlayDomain() *string
}

type StartRecordStreamRequest struct {
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
	// 310000*****000002
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// example.com
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
}

func (s StartRecordStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRecordStreamRequest) GoString() string {
	return s.String()
}

func (s *StartRecordStreamRequest) GetApp() *string {
	return s.App
}

func (s *StartRecordStreamRequest) GetId() *string {
	return s.Id
}

func (s *StartRecordStreamRequest) GetName() *string {
	return s.Name
}

func (s *StartRecordStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartRecordStreamRequest) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *StartRecordStreamRequest) SetApp(v string) *StartRecordStreamRequest {
	s.App = &v
	return s
}

func (s *StartRecordStreamRequest) SetId(v string) *StartRecordStreamRequest {
	s.Id = &v
	return s
}

func (s *StartRecordStreamRequest) SetName(v string) *StartRecordStreamRequest {
	s.Name = &v
	return s
}

func (s *StartRecordStreamRequest) SetOwnerId(v int64) *StartRecordStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *StartRecordStreamRequest) SetPlayDomain(v string) *StartRecordStreamRequest {
	s.PlayDomain = &v
	return s
}

func (s *StartRecordStreamRequest) Validate() error {
	return dara.Validate(s)
}
