// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeywordLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibId(v string) *DeleteKeywordLibRequest
	GetLibId() *string
	SetRegionId(v string) *DeleteKeywordLibRequest
	GetRegionId() *string
}

type DeleteKeywordLibRequest struct {
	// Keyword library ID.
	//
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteKeywordLibRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeywordLibRequest) GetLibId() *string {
	return s.LibId
}

func (s *DeleteKeywordLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteKeywordLibRequest) SetLibId(v string) *DeleteKeywordLibRequest {
	s.LibId = &v
	return s
}

func (s *DeleteKeywordLibRequest) SetRegionId(v string) *DeleteKeywordLibRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteKeywordLibRequest) Validate() error {
	return dara.Validate(s)
}
