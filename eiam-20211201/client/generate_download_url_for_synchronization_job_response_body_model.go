// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDownloadUrlForSynchronizationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileDownloadUrl(v string) *GenerateDownloadUrlForSynchronizationJobResponseBody
	GetFileDownloadUrl() *string
	SetRequestId(v string) *GenerateDownloadUrlForSynchronizationJobResponseBody
	GetRequestId() *string
}

type GenerateDownloadUrlForSynchronizationJobResponseBody struct {
	// example:
	//
	// https://test.oss.aliyuncs.com/idaas_ly77wa2oexrciw5v672vxxxx/tmp/eiam_user_export_1766469463365.csv
	FileDownloadUrl *string `json:"FileDownloadUrl,omitempty" xml:"FileDownloadUrl,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateDownloadUrlForSynchronizationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateDownloadUrlForSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDownloadUrlForSynchronizationJobResponseBody) GetFileDownloadUrl() *string {
	return s.FileDownloadUrl
}

func (s *GenerateDownloadUrlForSynchronizationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateDownloadUrlForSynchronizationJobResponseBody) SetFileDownloadUrl(v string) *GenerateDownloadUrlForSynchronizationJobResponseBody {
	s.FileDownloadUrl = &v
	return s
}

func (s *GenerateDownloadUrlForSynchronizationJobResponseBody) SetRequestId(v string) *GenerateDownloadUrlForSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDownloadUrlForSynchronizationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
