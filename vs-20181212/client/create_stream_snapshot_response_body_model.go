// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFormat(v string) *CreateStreamSnapshotResponseBody
	GetFormat() *string
	SetHeight(v int64) *CreateStreamSnapshotResponseBody
	GetHeight() *int64
	SetId(v string) *CreateStreamSnapshotResponseBody
	GetId() *string
	SetOssBucket(v string) *CreateStreamSnapshotResponseBody
	GetOssBucket() *string
	SetOssEndpoint(v string) *CreateStreamSnapshotResponseBody
	GetOssEndpoint() *string
	SetOssObject(v string) *CreateStreamSnapshotResponseBody
	GetOssObject() *string
	SetRequestId(v string) *CreateStreamSnapshotResponseBody
	GetRequestId() *string
	SetTimestamp(v int64) *CreateStreamSnapshotResponseBody
	GetTimestamp() *int64
	SetUrl(v string) *CreateStreamSnapshotResponseBody
	GetUrl() *string
	SetWidth(v int64) *CreateStreamSnapshotResponseBody
	GetWidth() *int64
}

type CreateStreamSnapshotResponseBody struct {
	// example:
	//
	// jpg
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// 720
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 323*****997-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// examplebucket
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// example:
	//
	// oss-cn-qingdao.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// example:
	//
	// photos/live/340200*****100049/ondemand-1639126601767.jpg
	OssObject *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1639126601767
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// http://examplebucket.oss-*****.aliyuncs.com/photos/live/340200*****100049/ondemand-1639126601767.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// 1280
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CreateStreamSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStreamSnapshotResponseBody) GetFormat() *string {
	return s.Format
}

func (s *CreateStreamSnapshotResponseBody) GetHeight() *int64 {
	return s.Height
}

func (s *CreateStreamSnapshotResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateStreamSnapshotResponseBody) GetOssBucket() *string {
	return s.OssBucket
}

func (s *CreateStreamSnapshotResponseBody) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *CreateStreamSnapshotResponseBody) GetOssObject() *string {
	return s.OssObject
}

func (s *CreateStreamSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStreamSnapshotResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CreateStreamSnapshotResponseBody) GetUrl() *string {
	return s.Url
}

func (s *CreateStreamSnapshotResponseBody) GetWidth() *int64 {
	return s.Width
}

func (s *CreateStreamSnapshotResponseBody) SetFormat(v string) *CreateStreamSnapshotResponseBody {
	s.Format = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetHeight(v int64) *CreateStreamSnapshotResponseBody {
	s.Height = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetId(v string) *CreateStreamSnapshotResponseBody {
	s.Id = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetOssBucket(v string) *CreateStreamSnapshotResponseBody {
	s.OssBucket = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetOssEndpoint(v string) *CreateStreamSnapshotResponseBody {
	s.OssEndpoint = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetOssObject(v string) *CreateStreamSnapshotResponseBody {
	s.OssObject = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetRequestId(v string) *CreateStreamSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetTimestamp(v int64) *CreateStreamSnapshotResponseBody {
	s.Timestamp = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetUrl(v string) *CreateStreamSnapshotResponseBody {
	s.Url = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) SetWidth(v int64) *CreateStreamSnapshotResponseBody {
	s.Width = &v
	return s
}

func (s *CreateStreamSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
