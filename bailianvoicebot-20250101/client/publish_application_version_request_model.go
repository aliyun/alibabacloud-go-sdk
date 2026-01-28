// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishApplicationVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *PublishApplicationVersionRequest
	GetApplicationId() *string
	SetBusinessUnitId(v string) *PublishApplicationVersionRequest
	GetBusinessUnitId() *string
	SetVersionId(v string) *PublishApplicationVersionRequest
	GetVersionId() *string
}

type PublishApplicationVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 82ea16d1-425c-4c03-9be5-cc91de9779ed
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s PublishApplicationVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishApplicationVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishApplicationVersionRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *PublishApplicationVersionRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *PublishApplicationVersionRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *PublishApplicationVersionRequest) SetApplicationId(v string) *PublishApplicationVersionRequest {
	s.ApplicationId = &v
	return s
}

func (s *PublishApplicationVersionRequest) SetBusinessUnitId(v string) *PublishApplicationVersionRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *PublishApplicationVersionRequest) SetVersionId(v string) *PublishApplicationVersionRequest {
	s.VersionId = &v
	return s
}

func (s *PublishApplicationVersionRequest) Validate() error {
	return dara.Validate(s)
}
