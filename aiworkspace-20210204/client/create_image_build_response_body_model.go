// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageBuildResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageBuildId(v string) *CreateImageBuildResponseBody
	GetImageBuildId() *string
	SetImageBuildJobId(v string) *CreateImageBuildResponseBody
	GetImageBuildJobId() *string
}

type CreateImageBuildResponseBody struct {
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// build-****ks92
	ImageBuildId *string `json:"ImageBuildId,omitempty" xml:"ImageBuildId,omitempty"`
	// example:
	//
	// dlc-****ks92
	ImageBuildJobId *string `json:"ImageBuildJobId,omitempty" xml:"ImageBuildJobId,omitempty"`
}

func (s CreateImageBuildResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageBuildResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageBuildResponseBody) GetImageBuildId() *string {
	return s.ImageBuildId
}

func (s *CreateImageBuildResponseBody) GetImageBuildJobId() *string {
	return s.ImageBuildJobId
}

func (s *CreateImageBuildResponseBody) SetImageBuildId(v string) *CreateImageBuildResponseBody {
	s.ImageBuildId = &v
	return s
}

func (s *CreateImageBuildResponseBody) SetImageBuildJobId(v string) *CreateImageBuildResponseBody {
	s.ImageBuildJobId = &v
	return s
}

func (s *CreateImageBuildResponseBody) Validate() error {
	return dara.Validate(s)
}
