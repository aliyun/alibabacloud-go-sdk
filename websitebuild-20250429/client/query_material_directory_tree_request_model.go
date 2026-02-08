// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialDirectoryTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QueryMaterialDirectoryTreeRequest
	GetBizId() *string
	SetHiddenPublic(v bool) *QueryMaterialDirectoryTreeRequest
	GetHiddenPublic() *bool
	SetRoot(v bool) *QueryMaterialDirectoryTreeRequest
	GetRoot() *bool
}

type QueryMaterialDirectoryTreeRequest struct {
	// example:
	//
	// WS20250731233102000001
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	HiddenPublic *bool   `json:"HiddenPublic,omitempty" xml:"HiddenPublic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Root *bool `json:"Root,omitempty" xml:"Root,omitempty"`
}

func (s QueryMaterialDirectoryTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialDirectoryTreeRequest) GoString() string {
	return s.String()
}

func (s *QueryMaterialDirectoryTreeRequest) GetBizId() *string {
	return s.BizId
}

func (s *QueryMaterialDirectoryTreeRequest) GetHiddenPublic() *bool {
	return s.HiddenPublic
}

func (s *QueryMaterialDirectoryTreeRequest) GetRoot() *bool {
	return s.Root
}

func (s *QueryMaterialDirectoryTreeRequest) SetBizId(v string) *QueryMaterialDirectoryTreeRequest {
	s.BizId = &v
	return s
}

func (s *QueryMaterialDirectoryTreeRequest) SetHiddenPublic(v bool) *QueryMaterialDirectoryTreeRequest {
	s.HiddenPublic = &v
	return s
}

func (s *QueryMaterialDirectoryTreeRequest) SetRoot(v bool) *QueryMaterialDirectoryTreeRequest {
	s.Root = &v
	return s
}

func (s *QueryMaterialDirectoryTreeRequest) Validate() error {
	return dara.Validate(s)
}
