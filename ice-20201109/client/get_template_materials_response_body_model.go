// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateMaterialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaterialUrls(v string) *GetTemplateMaterialsResponseBody
	GetMaterialUrls() *string
	SetRequestId(v string) *GetTemplateMaterialsResponseBody
	GetRequestId() *string
}

type GetTemplateMaterialsResponseBody struct {
	// The URLs of the associated materials.
	//
	// example:
	//
	// {"music.mp3":"https://bucket.oss-cn-shanghai.aliyuncs.com/music.mp3?sign=xxx","config.json":"https://bucket.oss-cn-shanghai.aliyuncs.com/config.json?sign=xxx","assets/1.jpg":"https://bucket.oss-cn-shanghai.aliyuncs.com/assets/1.jpg?sign=xxx"}
	MaterialUrls *string `json:"MaterialUrls,omitempty" xml:"MaterialUrls,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTemplateMaterialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateMaterialsResponseBody) GetMaterialUrls() *string {
	return s.MaterialUrls
}

func (s *GetTemplateMaterialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateMaterialsResponseBody) SetMaterialUrls(v string) *GetTemplateMaterialsResponseBody {
	s.MaterialUrls = &v
	return s
}

func (s *GetTemplateMaterialsResponseBody) SetRequestId(v string) *GetTemplateMaterialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateMaterialsResponseBody) Validate() error {
	return dara.Validate(s)
}
