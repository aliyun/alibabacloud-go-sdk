// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpire(v int64) *GetArtifactDownloadUrlResponseBody
	GetExpire() *int64
	SetRequestId(v string) *GetArtifactDownloadUrlResponseBody
	GetRequestId() *string
	SetUrl(v string) *GetArtifactDownloadUrlResponseBody
	GetUrl() *string
}

type GetArtifactDownloadUrlResponseBody struct {
	// example:
	//
	// 1770000000
	Expire *int64 `json:"expire,omitempty" xml:"expire,omitempty"`
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-1234567890AB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com/agents/123/sample-agent/home/starops/reports/summary.pdf?response-content-disposition=attachment&Expires=1770000000&OSSAccessKeyId=LTAI******&Signature=******
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetArtifactDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactDownloadUrlResponseBody) GetExpire() *int64 {
	return s.Expire
}

func (s *GetArtifactDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetArtifactDownloadUrlResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GetArtifactDownloadUrlResponseBody) SetExpire(v int64) *GetArtifactDownloadUrlResponseBody {
	s.Expire = &v
	return s
}

func (s *GetArtifactDownloadUrlResponseBody) SetRequestId(v string) *GetArtifactDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactDownloadUrlResponseBody) SetUrl(v string) *GetArtifactDownloadUrlResponseBody {
	s.Url = &v
	return s
}

func (s *GetArtifactDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
