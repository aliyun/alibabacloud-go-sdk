// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAnswerLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibId(v string) *DeleteAnswerLibRequest
	GetLibId() *string
	SetRegionId(v string) *DeleteAnswerLibRequest
	GetRegionId() *string
}

type DeleteAnswerLibRequest struct {
	// example:
	//
	// alxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAnswerLibRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAnswerLibRequest) GoString() string {
	return s.String()
}

func (s *DeleteAnswerLibRequest) GetLibId() *string {
	return s.LibId
}

func (s *DeleteAnswerLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAnswerLibRequest) SetLibId(v string) *DeleteAnswerLibRequest {
	s.LibId = &v
	return s
}

func (s *DeleteAnswerLibRequest) SetRegionId(v string) *DeleteAnswerLibRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAnswerLibRequest) Validate() error {
	return dara.Validate(s)
}
