// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMixStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteMixStreamRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteMixStreamRequest
	GetDomainName() *string
	SetMixStreamId(v string) *DeleteMixStreamRequest
	GetMixStreamId() *string
	SetOwnerId(v int64) *DeleteMixStreamRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteMixStreamRequest
	GetRegionId() *string
	SetStreamName(v string) *DeleteMixStreamRequest
	GetStreamName() *string
}

type DeleteMixStreamRequest struct {
	// The name of the application.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the stream mixing task. If the task was created by calling the [CreateMixStream](https://help.aliyun.com/document_detail/2848087.html) operation, check the value of the response parameter MixStreamId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 749b7594-86d6-37b1-513b-e1e19845****
	MixStreamId *string `json:"MixStreamId,omitempty" xml:"MixStreamId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the output stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteMixStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMixStreamRequest) GoString() string {
	return s.String()
}

func (s *DeleteMixStreamRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteMixStreamRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteMixStreamRequest) GetMixStreamId() *string {
	return s.MixStreamId
}

func (s *DeleteMixStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteMixStreamRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMixStreamRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DeleteMixStreamRequest) SetAppName(v string) *DeleteMixStreamRequest {
	s.AppName = &v
	return s
}

func (s *DeleteMixStreamRequest) SetDomainName(v string) *DeleteMixStreamRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteMixStreamRequest) SetMixStreamId(v string) *DeleteMixStreamRequest {
	s.MixStreamId = &v
	return s
}

func (s *DeleteMixStreamRequest) SetOwnerId(v int64) *DeleteMixStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMixStreamRequest) SetRegionId(v string) *DeleteMixStreamRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMixStreamRequest) SetStreamName(v string) *DeleteMixStreamRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteMixStreamRequest) Validate() error {
	return dara.Validate(s)
}
