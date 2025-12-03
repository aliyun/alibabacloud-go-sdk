// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorAccountId(v string) *UpdateApplicationResponseBody
	GetCreatorAccountId() *string
	SetDescription(v string) *UpdateApplicationResponseBody
	GetDescription() *string
	SetGmtCreate(v string) *UpdateApplicationResponseBody
	GetGmtCreate() *string
	SetName(v string) *UpdateApplicationResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateApplicationResponseBody
	GetRequestId() *string
}

type UpdateApplicationResponseBody struct {
	// example:
	//
	// 1332695887xxxxxx
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// example:
	//
	// 应用描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00.000+00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// testApp
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// FC93CE1A-8D7A-13A9-8306-7465DE2E5C0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBody) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *UpdateApplicationResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdateApplicationResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationResponseBody) SetCreatorAccountId(v string) *UpdateApplicationResponseBody {
	s.CreatorAccountId = &v
	return s
}

func (s *UpdateApplicationResponseBody) SetDescription(v string) *UpdateApplicationResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateApplicationResponseBody) SetGmtCreate(v string) *UpdateApplicationResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *UpdateApplicationResponseBody) SetName(v string) *UpdateApplicationResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateApplicationResponseBody) SetRequestId(v string) *UpdateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
