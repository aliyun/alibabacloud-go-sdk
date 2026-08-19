// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSiteProjectNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *CheckSiteProjectNameRequest
	GetProjectName() *string
	SetSiteId(v int64) *CheckSiteProjectNameRequest
	GetSiteId() *int64
}

type CheckSiteProjectNameRequest struct {
	// The real-time log project name.
	//
	// > Allowed character set (hyphens only, no underscores), length range, and naming rule examples (such as \\"ali-dcdn-log-56\\")
	//
	// This parameter is required.
	//
	// example:
	//
	// ali-dcdn-log-56
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The site ID. You can call [ListSites](https://help.aliyun.com/document_detail/2850189.html) to obtain the site ID.
	//
	// example:
	//
	// 12312312213212
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CheckSiteProjectNameRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckSiteProjectNameRequest) GoString() string {
	return s.String()
}

func (s *CheckSiteProjectNameRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CheckSiteProjectNameRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CheckSiteProjectNameRequest) SetProjectName(v string) *CheckSiteProjectNameRequest {
	s.ProjectName = &v
	return s
}

func (s *CheckSiteProjectNameRequest) SetSiteId(v int64) *CheckSiteProjectNameRequest {
	s.SiteId = &v
	return s
}

func (s *CheckSiteProjectNameRequest) Validate() error {
	return dara.Validate(s)
}
