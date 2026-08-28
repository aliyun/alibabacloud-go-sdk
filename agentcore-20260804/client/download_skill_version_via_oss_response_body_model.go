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
	// The response data.
	//
	// example:
	//
	// https://example.com/artifacts/example.zip
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
