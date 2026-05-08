// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUploadTokenResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetOssUploadTokenResult
	GetRequestId() *string
	SetUploadInfo(v *UploadInfo) *GetOssUploadTokenResult
	GetUploadInfo() *UploadInfo
}

type GetOssUploadTokenResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	UploadInfo *UploadInfo `json:"uploadInfo,omitempty" xml:"uploadInfo,omitempty"`
}

func (s GetOssUploadTokenResult) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadTokenResult) GoString() string {
	return s.String()
}

func (s *GetOssUploadTokenResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssUploadTokenResult) GetUploadInfo() *UploadInfo {
	return s.UploadInfo
}

func (s *GetOssUploadTokenResult) SetRequestId(v string) *GetOssUploadTokenResult {
	s.RequestId = &v
	return s
}

func (s *GetOssUploadTokenResult) SetUploadInfo(v *UploadInfo) *GetOssUploadTokenResult {
	s.UploadInfo = v
	return s
}

func (s *GetOssUploadTokenResult) Validate() error {
	if s.UploadInfo != nil {
		if err := s.UploadInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
