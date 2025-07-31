// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAnswerSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *DeleteAnswerSampleRequest
	GetIds() *string
	SetLibId(v string) *DeleteAnswerSampleRequest
	GetLibId() *string
	SetRegionId(v string) *DeleteAnswerSampleRequest
	GetRegionId() *string
}

type DeleteAnswerSampleRequest struct {
	// example:
	//
	// [15463605]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// example:
	//
	// alxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAnswerSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAnswerSampleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAnswerSampleRequest) GetIds() *string {
	return s.Ids
}

func (s *DeleteAnswerSampleRequest) GetLibId() *string {
	return s.LibId
}

func (s *DeleteAnswerSampleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAnswerSampleRequest) SetIds(v string) *DeleteAnswerSampleRequest {
	s.Ids = &v
	return s
}

func (s *DeleteAnswerSampleRequest) SetLibId(v string) *DeleteAnswerSampleRequest {
	s.LibId = &v
	return s
}

func (s *DeleteAnswerSampleRequest) SetRegionId(v string) *DeleteAnswerSampleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAnswerSampleRequest) Validate() error {
	return dara.Validate(s)
}
