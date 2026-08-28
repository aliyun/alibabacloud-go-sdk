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
	// The response data.
	//
	// example:
	//
	// skill-1234567890abcdef
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
