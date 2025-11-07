// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplication(v *CreateApplicationResponseBodyApplication) *CreateApplicationResponseBody
	GetApplication() *CreateApplicationResponseBodyApplication
	SetRequestId(v string) *CreateApplicationResponseBody
	GetRequestId() *string
}

type CreateApplicationResponseBody struct {
	// The information about the application.
	Application *CreateApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 274917E8-8E74-5928-A82F-4940F52F7ACB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) GetApplication() *CreateApplicationResponseBodyApplication {
	return s.Application
}

func (s *CreateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationResponseBody) SetApplication(v *CreateApplicationResponseBodyApplication) *CreateApplicationResponseBody {
	s.Application = v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationResponseBody) Validate() error {
	if s.Application != nil {
		if err := s.Application.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApplicationResponseBodyApplication struct {
	// The time when the application was created.
	//
	// example:
	//
	// 2021-09-07T09:17:46Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// example:
	//
	// Myapplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time when the application was updated.
	//
	// example:
	//
	// 2021-09-07T09:17:46Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateApplicationResponseBodyApplication) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplication) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateApplicationResponseBodyApplication) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationResponseBodyApplication) GetName() *string {
	return s.Name
}

func (s *CreateApplicationResponseBodyApplication) GetTags() map[string]*string {
	return s.Tags
}

func (s *CreateApplicationResponseBodyApplication) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *CreateApplicationResponseBodyApplication) SetCreateDate(v string) *CreateApplicationResponseBodyApplication {
	s.CreateDate = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetDescription(v string) *CreateApplicationResponseBodyApplication {
	s.Description = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetName(v string) *CreateApplicationResponseBodyApplication {
	s.Name = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetTags(v map[string]*string) *CreateApplicationResponseBodyApplication {
	s.Tags = v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetUpdateDate(v string) *CreateApplicationResponseBodyApplication {
	s.UpdateDate = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) Validate() error {
	return dara.Validate(s)
}
