// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeywordLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibId(v string) *GetKeywordLibRequest
	GetLibId() *string
	SetRegionId(v string) *GetKeywordLibRequest
	GetRegionId() *string
}

type GetKeywordLibRequest struct {
	// example:
	//
	// customxx_xxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetKeywordLibRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *GetKeywordLibRequest) GetLibId() *string {
	return s.LibId
}

func (s *GetKeywordLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetKeywordLibRequest) SetLibId(v string) *GetKeywordLibRequest {
	s.LibId = &v
	return s
}

func (s *GetKeywordLibRequest) SetRegionId(v string) *GetKeywordLibRequest {
	s.RegionId = &v
	return s
}

func (s *GetKeywordLibRequest) Validate() error {
	return dara.Validate(s)
}
