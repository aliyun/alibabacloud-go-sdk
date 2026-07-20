// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProjectDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetThirdPartId(v string) *ProjectDeleteRequest
	GetThirdPartId() *string
}

type ProjectDeleteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s ProjectDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s ProjectDeleteRequest) GoString() string {
	return s.String()
}

func (s *ProjectDeleteRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *ProjectDeleteRequest) SetThirdPartId(v string) *ProjectDeleteRequest {
	s.ThirdPartId = &v
	return s
}

func (s *ProjectDeleteRequest) Validate() error {
	return dara.Validate(s)
}
