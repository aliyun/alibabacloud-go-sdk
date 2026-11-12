// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTempDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetTempDownloadUrlRequest
	GetBizId() *string
	SetOssKey(v string) *GetTempDownloadUrlRequest
	GetOssKey() *string
}

type GetTempDownloadUrlRequest struct {
	// The business ID.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The OSS object key.
	//
	// example:
	//
	// backend/detection/objects/r-0008ujvfksltf5j45n81/task-000hckiuwyau0gwn17vq.jpg
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
}

func (s GetTempDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTempDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetTempDownloadUrlRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetTempDownloadUrlRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *GetTempDownloadUrlRequest) SetBizId(v string) *GetTempDownloadUrlRequest {
	s.BizId = &v
	return s
}

func (s *GetTempDownloadUrlRequest) SetOssKey(v string) *GetTempDownloadUrlRequest {
	s.OssKey = &v
	return s
}

func (s *GetTempDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
