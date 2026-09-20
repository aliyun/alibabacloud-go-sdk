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
	// The OSS URL for downloading the skill.
	//
	// example:
	//
	// https://sample-bucket.oss-region.aliyuncs.com/xxxxxx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D9E87E66-9EF0-5C10-A5E6-924020A0C9B7
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
