// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSkillViaOssResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UploadSkillViaOssResponseBody
	GetData() *string
	SetRequestId(v string) *UploadSkillViaOssResponseBody
	GetRequestId() *string
}

type UploadSkillViaOssResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadSkillViaOssResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadSkillViaOssResponseBody) GoString() string {
	return s.String()
}

func (s *UploadSkillViaOssResponseBody) GetData() *string {
	return s.Data
}

func (s *UploadSkillViaOssResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadSkillViaOssResponseBody) SetData(v string) *UploadSkillViaOssResponseBody {
	s.Data = &v
	return s
}

func (s *UploadSkillViaOssResponseBody) SetRequestId(v string) *UploadSkillViaOssResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadSkillViaOssResponseBody) Validate() error {
	return dara.Validate(s)
}
