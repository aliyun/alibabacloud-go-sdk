// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppTemplateDisplayName(v string) *GetApplicationResponseBody
	GetAppTemplateDisplayName() *string
	SetAppTemplateName(v string) *GetApplicationResponseBody
	GetAppTemplateName() *string
	SetCreatorAccountId(v string) *GetApplicationResponseBody
	GetCreatorAccountId() *string
	SetDescription(v string) *GetApplicationResponseBody
	GetDescription() *string
	SetGmtCreate(v string) *GetApplicationResponseBody
	GetGmtCreate() *string
	SetName(v string) *GetApplicationResponseBody
	GetName() *string
	SetRequestId(v string) *GetApplicationResponseBody
	GetRequestId() *string
}

type GetApplicationResponseBody struct {
	// example:
	//
	// 应用模版展示名称A
	AppTemplateDisplayName *string `json:"appTemplateDisplayName,omitempty" xml:"appTemplateDisplayName,omitempty"`
	// example:
	//
	// 应用模版名称A
	AppTemplateName *string `json:"appTemplateName,omitempty" xml:"appTemplateName,omitempty"`
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

func (s GetApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) GetAppTemplateDisplayName() *string {
	return s.AppTemplateDisplayName
}

func (s *GetApplicationResponseBody) GetAppTemplateName() *string {
	return s.AppTemplateName
}

func (s *GetApplicationResponseBody) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *GetApplicationResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetApplicationResponseBody) GetName() *string {
	return s.Name
}

func (s *GetApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationResponseBody) SetAppTemplateDisplayName(v string) *GetApplicationResponseBody {
	s.AppTemplateDisplayName = &v
	return s
}

func (s *GetApplicationResponseBody) SetAppTemplateName(v string) *GetApplicationResponseBody {
	s.AppTemplateName = &v
	return s
}

func (s *GetApplicationResponseBody) SetCreatorAccountId(v string) *GetApplicationResponseBody {
	s.CreatorAccountId = &v
	return s
}

func (s *GetApplicationResponseBody) SetDescription(v string) *GetApplicationResponseBody {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBody) SetGmtCreate(v string) *GetApplicationResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetApplicationResponseBody) SetName(v string) *GetApplicationResponseBody {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
