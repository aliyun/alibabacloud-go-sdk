// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSemanticFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UploadSemanticFileResponseBodyData) *UploadSemanticFileResponseBody
	GetData() *UploadSemanticFileResponseBodyData
	SetRequestId(v string) *UploadSemanticFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadSemanticFileResponseBody
	GetSuccess() *bool
}

type UploadSemanticFileResponseBody struct {
	// The attachment upload slot information. PUT the file to Data.UploadUrl before Data.ExpiresAt, then use Data.FileId to create a single-file semantic task.
	Data *UploadSemanticFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. Used for locating logs and troubleshooting issues.
	//
	// example:
	//
	// 676271D6-53B4-57BE-89FA-72F7AE1418DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadSemanticFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadSemanticFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadSemanticFileResponseBody) GetData() *UploadSemanticFileResponseBodyData {
	return s.Data
}

func (s *UploadSemanticFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadSemanticFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadSemanticFileResponseBody) SetData(v *UploadSemanticFileResponseBodyData) *UploadSemanticFileResponseBody {
	s.Data = v
	return s
}

func (s *UploadSemanticFileResponseBody) SetRequestId(v string) *UploadSemanticFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadSemanticFileResponseBody) SetSuccess(v bool) *UploadSemanticFileResponseBody {
	s.Success = &v
	return s
}

func (s *UploadSemanticFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadSemanticFileResponseBodyData struct {
	// The expiration time of UploadUrl, expressed as a Unix timestamp in milliseconds. After this time, call UploadSemanticFile again to request a new URL.
	//
	// example:
	//
	// 1700001800000
	ExpiresAt *int64 `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// The unique identifier of the attachment. After completing the PUT upload to UploadUrl, pass this value to the ReferenceFileIds parameter of CreateSemanticJob.
	//
	// example:
	//
	// FID1
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The temporary OSS PUT upload URL. Valid for 30 minutes. Only the specified object can be uploaded. Use the ContentType from the request when performing the PUT request. Do not log or distribute the full URL.
	//
	// example:
	//
	// https://example.com/temporary-upload-url
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
}

func (s UploadSemanticFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UploadSemanticFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadSemanticFileResponseBodyData) GetExpiresAt() *int64 {
	return s.ExpiresAt
}

func (s *UploadSemanticFileResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *UploadSemanticFileResponseBodyData) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *UploadSemanticFileResponseBodyData) SetExpiresAt(v int64) *UploadSemanticFileResponseBodyData {
	s.ExpiresAt = &v
	return s
}

func (s *UploadSemanticFileResponseBodyData) SetFileId(v string) *UploadSemanticFileResponseBodyData {
	s.FileId = &v
	return s
}

func (s *UploadSemanticFileResponseBodyData) SetUploadUrl(v string) *UploadSemanticFileResponseBodyData {
	s.UploadUrl = &v
	return s
}

func (s *UploadSemanticFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
