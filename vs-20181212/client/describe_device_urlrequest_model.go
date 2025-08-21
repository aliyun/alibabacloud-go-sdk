// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceURLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuth(v bool) *DescribeDeviceURLRequest
	GetAuth() *bool
	SetExpire(v int64) *DescribeDeviceURLRequest
	GetExpire() *int64
	SetId(v string) *DescribeDeviceURLRequest
	GetId() *string
	SetMode(v string) *DescribeDeviceURLRequest
	GetMode() *string
	SetOutProtocol(v string) *DescribeDeviceURLRequest
	GetOutProtocol() *string
	SetOwnerId(v int64) *DescribeDeviceURLRequest
	GetOwnerId() *int64
	SetStream(v string) *DescribeDeviceURLRequest
	GetStream() *string
	SetType(v string) *DescribeDeviceURLRequest
	GetType() *string
}

type DescribeDeviceURLRequest struct {
	Auth *bool `json:"Auth,omitempty" xml:"Auth,omitempty"`
	// example:
	//
	// 3600
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// push
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rtmp
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// live001
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDeviceURLRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceURLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceURLRequest) GetAuth() *bool {
	return s.Auth
}

func (s *DescribeDeviceURLRequest) GetExpire() *int64 {
	return s.Expire
}

func (s *DescribeDeviceURLRequest) GetId() *string {
	return s.Id
}

func (s *DescribeDeviceURLRequest) GetMode() *string {
	return s.Mode
}

func (s *DescribeDeviceURLRequest) GetOutProtocol() *string {
	return s.OutProtocol
}

func (s *DescribeDeviceURLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDeviceURLRequest) GetStream() *string {
	return s.Stream
}

func (s *DescribeDeviceURLRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDeviceURLRequest) SetAuth(v bool) *DescribeDeviceURLRequest {
	s.Auth = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetExpire(v int64) *DescribeDeviceURLRequest {
	s.Expire = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetId(v string) *DescribeDeviceURLRequest {
	s.Id = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetMode(v string) *DescribeDeviceURLRequest {
	s.Mode = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetOutProtocol(v string) *DescribeDeviceURLRequest {
	s.OutProtocol = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetOwnerId(v int64) *DescribeDeviceURLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetStream(v string) *DescribeDeviceURLRequest {
	s.Stream = &v
	return s
}

func (s *DescribeDeviceURLRequest) SetType(v string) *DescribeDeviceURLRequest {
	s.Type = &v
	return s
}

func (s *DescribeDeviceURLRequest) Validate() error {
	return dara.Validate(s)
}
