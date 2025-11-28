// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllPrivacyAlgorithmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListAllPrivacyAlgorithmRequest
	GetRegionId() *string
}

type ListAllPrivacyAlgorithmRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAllPrivacyAlgorithmRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllPrivacyAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *ListAllPrivacyAlgorithmRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAllPrivacyAlgorithmRequest) SetRegionId(v string) *ListAllPrivacyAlgorithmRequest {
	s.RegionId = &v
	return s
}

func (s *ListAllPrivacyAlgorithmRequest) Validate() error {
	return dara.Validate(s)
}
