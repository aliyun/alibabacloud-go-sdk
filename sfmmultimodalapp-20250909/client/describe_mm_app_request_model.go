// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMmAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeMmAppRequest
	GetAppId() *string
	SetWorkspaceId(v string) *DescribeMmAppRequest
	GetWorkspaceId() *string
}

type DescribeMmAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_xxxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// llm-xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeMmAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMmAppRequest) GoString() string {
	return s.String()
}

func (s *DescribeMmAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeMmAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeMmAppRequest) SetAppId(v string) *DescribeMmAppRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMmAppRequest) SetWorkspaceId(v string) *DescribeMmAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeMmAppRequest) Validate() error {
	return dara.Validate(s)
}
