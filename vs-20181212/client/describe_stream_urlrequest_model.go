// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamURLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuth(v bool) *DescribeStreamURLRequest
	GetAuth() *bool
	SetAuthKey(v string) *DescribeStreamURLRequest
	GetAuthKey() *string
	SetEndTime(v int64) *DescribeStreamURLRequest
	GetEndTime() *int64
	SetExpire(v int64) *DescribeStreamURLRequest
	GetExpire() *int64
	SetId(v string) *DescribeStreamURLRequest
	GetId() *string
	SetOutProtocol(v string) *DescribeStreamURLRequest
	GetOutProtocol() *string
	SetOwnerId(v int64) *DescribeStreamURLRequest
	GetOwnerId() *int64
	SetStartTime(v int64) *DescribeStreamURLRequest
	GetStartTime() *int64
	SetTranscode(v string) *DescribeStreamURLRequest
	GetTranscode() *string
	SetType(v string) *DescribeStreamURLRequest
	GetType() *string
}

type DescribeStreamURLRequest struct {
	// Specifies whether to generate a signed URL. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	Auth *bool `json:"Auth,omitempty" xml:"Auth,omitempty"`
	// The primary key associated with the playback domain name. This key is used to generate the authentication URL.
	//
	// > Call the [DescribeVsDomainConfigs](https://next.api.aliyun.com/document/vs/2018-12-12/DescribeVsDomainConfigs) operation to query the \\`AuthKey\\` information.
	//
	// example:
	//
	// ocs*****ace
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The end time. This parameter applies to \\`vod\\` streams.<br>
	//
	// A UNIX timestamp. Unit: seconds.<br>
	//
	// example:
	//
	// 1571649499
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time-to-live (TTL) of the URL. Unit: seconds.
	//
	// example:
	//
	// 3600
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The stream ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 323*****997-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The playback protocol for the stream. Valid values:
	//
	// - rtmp
	//
	// - flv
	//
	// - hls
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start time. This parameter applies to \\`vod\\` streams.<br>
	//
	// A UNIX timestamp. Unit: seconds.<br>
	//
	// example:
	//
	// 1571639499
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the transcoding rule. This parameter is valid only after a transcoding template is attached.
	//
	// example:
	//
	// sd
	Transcode *string `json:"Transcode,omitempty" xml:"Transcode,omitempty"`
	// The type of the stream. The default value is \\`live\\`. Valid values:
	//
	// - \\`live\\`: a live stream.
	//
	// - \\`vod\\`: a video-on-demand (VOD) stream, such as a historical stream from a Network Video Recorder (NVR).
	//
	// example:
	//
	// live
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeStreamURLRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamURLRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamURLRequest) GetAuth() *bool {
	return s.Auth
}

func (s *DescribeStreamURLRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *DescribeStreamURLRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeStreamURLRequest) GetExpire() *int64 {
	return s.Expire
}

func (s *DescribeStreamURLRequest) GetId() *string {
	return s.Id
}

func (s *DescribeStreamURLRequest) GetOutProtocol() *string {
	return s.OutProtocol
}

func (s *DescribeStreamURLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeStreamURLRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeStreamURLRequest) GetTranscode() *string {
	return s.Transcode
}

func (s *DescribeStreamURLRequest) GetType() *string {
	return s.Type
}

func (s *DescribeStreamURLRequest) SetAuth(v bool) *DescribeStreamURLRequest {
	s.Auth = &v
	return s
}

func (s *DescribeStreamURLRequest) SetAuthKey(v string) *DescribeStreamURLRequest {
	s.AuthKey = &v
	return s
}

func (s *DescribeStreamURLRequest) SetEndTime(v int64) *DescribeStreamURLRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeStreamURLRequest) SetExpire(v int64) *DescribeStreamURLRequest {
	s.Expire = &v
	return s
}

func (s *DescribeStreamURLRequest) SetId(v string) *DescribeStreamURLRequest {
	s.Id = &v
	return s
}

func (s *DescribeStreamURLRequest) SetOutProtocol(v string) *DescribeStreamURLRequest {
	s.OutProtocol = &v
	return s
}

func (s *DescribeStreamURLRequest) SetOwnerId(v int64) *DescribeStreamURLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStreamURLRequest) SetStartTime(v int64) *DescribeStreamURLRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeStreamURLRequest) SetTranscode(v string) *DescribeStreamURLRequest {
	s.Transcode = &v
	return s
}

func (s *DescribeStreamURLRequest) SetType(v string) *DescribeStreamURLRequest {
	s.Type = &v
	return s
}

func (s *DescribeStreamURLRequest) Validate() error {
	return dara.Validate(s)
}
