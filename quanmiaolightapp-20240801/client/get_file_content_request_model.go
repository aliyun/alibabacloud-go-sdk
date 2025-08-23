// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileKey(v string) *GetFileContentRequest
	GetFileKey() *string
}

type GetFileContentRequest struct {
	// example:
	//
	// oss://default/aimiaobi-service-prod/aimiaobi/temp/1154600634854327_10045847/300469535473178749_300469535473178749_ee11508152b74137ac5747a6f632256e.docx
	FileKey *string `json:"fileKey,omitempty" xml:"fileKey,omitempty"`
}

func (s GetFileContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileContentRequest) GoString() string {
	return s.String()
}

func (s *GetFileContentRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *GetFileContentRequest) SetFileKey(v string) *GetFileContentRequest {
	s.FileKey = &v
	return s
}

func (s *GetFileContentRequest) Validate() error {
	return dara.Validate(s)
}
