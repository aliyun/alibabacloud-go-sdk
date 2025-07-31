// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnswerLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListAnswerLibRequest
	GetRegionId() *string
}

type ListAnswerLibRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAnswerLibRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAnswerLibRequest) GoString() string {
	return s.String()
}

func (s *ListAnswerLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAnswerLibRequest) SetRegionId(v string) *ListAnswerLibRequest {
	s.RegionId = &v
	return s
}

func (s *ListAnswerLibRequest) Validate() error {
	return dara.Validate(s)
}
