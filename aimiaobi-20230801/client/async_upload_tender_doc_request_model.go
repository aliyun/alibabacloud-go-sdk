// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncUploadTenderDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileKey(v string) *AsyncUploadTenderDocRequest
	GetFileKey() *string
	SetTenderDocName(v string) *AsyncUploadTenderDocRequest
	GetTenderDocName() *string
	SetWorkspaceId(v string) *AsyncUploadTenderDocRequest
	GetWorkspaceId() *string
}

type AsyncUploadTenderDocRequest struct {
	// example:
	//
	// oss://default/aimiaobi-service-prod/aimiaobi/materialDocument/1601892701595700_10169811/208757545922605632_yst-test_9eb7d7e1deb543d88e2d6f1c9df456ef.docx
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// tender.pdf
	TenderDocName *string `json:"TenderDocName,omitempty" xml:"TenderDocName,omitempty"`
	// example:
	//
	// llm-az2gglkjauwnnhpq
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AsyncUploadTenderDocRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadTenderDocRequest) GoString() string {
	return s.String()
}

func (s *AsyncUploadTenderDocRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *AsyncUploadTenderDocRequest) GetTenderDocName() *string {
	return s.TenderDocName
}

func (s *AsyncUploadTenderDocRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncUploadTenderDocRequest) SetFileKey(v string) *AsyncUploadTenderDocRequest {
	s.FileKey = &v
	return s
}

func (s *AsyncUploadTenderDocRequest) SetTenderDocName(v string) *AsyncUploadTenderDocRequest {
	s.TenderDocName = &v
	return s
}

func (s *AsyncUploadTenderDocRequest) SetWorkspaceId(v string) *AsyncUploadTenderDocRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncUploadTenderDocRequest) Validate() error {
	return dara.Validate(s)
}
