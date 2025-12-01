// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreOssImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestoreOssImageResponseBody
	GetRequestId() *string
	SetRestoredImageKey(v string) *RestoreOssImageResponseBody
	GetRestoredImageKey() *string
}

type RestoreOssImageResponseBody struct {
	// example:
	//
	// 208B016D-4CB9-4A85-96A5-0B8ED1EBF271
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// aliyun_dsc_original /dir1/test.png
	RestoredImageKey *string `json:"RestoredImageKey,omitempty" xml:"RestoredImageKey,omitempty"`
}

func (s RestoreOssImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreOssImageResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreOssImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreOssImageResponseBody) GetRestoredImageKey() *string {
	return s.RestoredImageKey
}

func (s *RestoreOssImageResponseBody) SetRequestId(v string) *RestoreOssImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreOssImageResponseBody) SetRestoredImageKey(v string) *RestoreOssImageResponseBody {
	s.RestoredImageKey = &v
	return s
}

func (s *RestoreOssImageResponseBody) Validate() error {
	return dara.Validate(s)
}
