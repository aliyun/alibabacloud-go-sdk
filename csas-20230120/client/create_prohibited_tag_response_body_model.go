// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProhibitedTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateProhibitedTagResponseBody
	GetRequestId() *string
	SetTag(v *CreateProhibitedTagResponseBodyTag) *CreateProhibitedTagResponseBody
	GetTag() *CreateProhibitedTagResponseBodyTag
}

type CreateProhibitedTagResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// C7F30ABA-67BD-537D-A516-8DA20DA1F28C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The disabled software tag.
	Tag *CreateProhibitedTagResponseBodyTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
}

func (s CreateProhibitedTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProhibitedTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProhibitedTagResponseBody) GetTag() *CreateProhibitedTagResponseBodyTag {
	return s.Tag
}

func (s *CreateProhibitedTagResponseBody) SetRequestId(v string) *CreateProhibitedTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProhibitedTagResponseBody) SetTag(v *CreateProhibitedTagResponseBodyTag) *CreateProhibitedTagResponseBody {
	s.Tag = v
	return s
}

func (s *CreateProhibitedTagResponseBody) Validate() error {
	if s.Tag != nil {
		if err := s.Tag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProhibitedTagResponseBodyTag struct {
	// The time when the disabled software tag was created, in the yyyy-MM-dd HH:mm:ss format. The time is displayed in UTC+8.
	//
	// example:
	//
	// 2026-08-19 10:24:31
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the disabled software tag.
	//
	// example:
	//
	// test template create get delete
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the disabled software tag.
	//
	// example:
	//
	// autotest_37bf6a18
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the disabled software tag.
	//
	// example:
	//
	// tag-4a4046838f77****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s CreateProhibitedTagResponseBodyTag) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedTagResponseBodyTag) GoString() string {
	return s.String()
}

func (s *CreateProhibitedTagResponseBodyTag) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateProhibitedTagResponseBodyTag) GetDescription() *string {
	return s.Description
}

func (s *CreateProhibitedTagResponseBodyTag) GetName() *string {
	return s.Name
}

func (s *CreateProhibitedTagResponseBodyTag) GetTagId() *string {
	return s.TagId
}

func (s *CreateProhibitedTagResponseBodyTag) SetCreateTime(v string) *CreateProhibitedTagResponseBodyTag {
	s.CreateTime = &v
	return s
}

func (s *CreateProhibitedTagResponseBodyTag) SetDescription(v string) *CreateProhibitedTagResponseBodyTag {
	s.Description = &v
	return s
}

func (s *CreateProhibitedTagResponseBodyTag) SetName(v string) *CreateProhibitedTagResponseBodyTag {
	s.Name = &v
	return s
}

func (s *CreateProhibitedTagResponseBodyTag) SetTagId(v string) *CreateProhibitedTagResponseBodyTag {
	s.TagId = &v
	return s
}

func (s *CreateProhibitedTagResponseBodyTag) Validate() error {
	return dara.Validate(s)
}
