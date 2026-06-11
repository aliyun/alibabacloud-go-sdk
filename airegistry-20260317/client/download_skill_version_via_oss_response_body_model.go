// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSkillVersionViaOssResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DownloadSkillVersionViaOssResponseBody
	GetData() *string
	SetRequestId(v string) *DownloadSkillVersionViaOssResponseBody
	GetRequestId() *string
}

type DownloadSkillVersionViaOssResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadSkillVersionViaOssResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadSkillVersionViaOssResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadSkillVersionViaOssResponseBody) GetData() *string {
	return s.Data
}

func (s *DownloadSkillVersionViaOssResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadSkillVersionViaOssResponseBody) SetData(v string) *DownloadSkillVersionViaOssResponseBody {
	s.Data = &v
	return s
}

func (s *DownloadSkillVersionViaOssResponseBody) SetRequestId(v string) *DownloadSkillVersionViaOssResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadSkillVersionViaOssResponseBody) Validate() error {
	return dara.Validate(s)
}
