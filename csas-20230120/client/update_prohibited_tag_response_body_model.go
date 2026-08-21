// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProhibitedTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateProhibitedTagResponseBody
	GetRequestId() *string
	SetTag(v *UpdateProhibitedTagResponseBodyTag) *UpdateProhibitedTagResponseBody
	GetTag() *UpdateProhibitedTagResponseBodyTag
}

type UpdateProhibitedTagResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1D3BCF94-7F83-559E-82D9-C891BBB32FC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The prohibited software tag.
	Tag *UpdateProhibitedTagResponseBodyTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
}

func (s UpdateProhibitedTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedTagResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProhibitedTagResponseBody) GetTag() *UpdateProhibitedTagResponseBodyTag {
	return s.Tag
}

func (s *UpdateProhibitedTagResponseBody) SetRequestId(v string) *UpdateProhibitedTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProhibitedTagResponseBody) SetTag(v *UpdateProhibitedTagResponseBodyTag) *UpdateProhibitedTagResponseBody {
	s.Tag = v
	return s
}

func (s *UpdateProhibitedTagResponseBody) Validate() error {
	if s.Tag != nil {
		if err := s.Tag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProhibitedTagResponseBodyTag struct {
	// The creation time of the prohibited software tag, in the yyyy-MM-dd HH:mm:ss format. The time is displayed in UTC+8.
	//
	// example:
	//
	// 2025-09-05 10:20:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the prohibited software tag.
	//
	// example:
	//
	// test constraints
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the prohibited software tag.
	//
	// example:
	//
	// PolicyC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the prohibited software tag.
	//
	// example:
	//
	// tag-d730092d87ec****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s UpdateProhibitedTagResponseBodyTag) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedTagResponseBodyTag) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedTagResponseBodyTag) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateProhibitedTagResponseBodyTag) GetDescription() *string {
	return s.Description
}

func (s *UpdateProhibitedTagResponseBodyTag) GetName() *string {
	return s.Name
}

func (s *UpdateProhibitedTagResponseBodyTag) GetTagId() *string {
	return s.TagId
}

func (s *UpdateProhibitedTagResponseBodyTag) SetCreateTime(v string) *UpdateProhibitedTagResponseBodyTag {
	s.CreateTime = &v
	return s
}

func (s *UpdateProhibitedTagResponseBodyTag) SetDescription(v string) *UpdateProhibitedTagResponseBodyTag {
	s.Description = &v
	return s
}

func (s *UpdateProhibitedTagResponseBodyTag) SetName(v string) *UpdateProhibitedTagResponseBodyTag {
	s.Name = &v
	return s
}

func (s *UpdateProhibitedTagResponseBodyTag) SetTagId(v string) *UpdateProhibitedTagResponseBodyTag {
	s.TagId = &v
	return s
}

func (s *UpdateProhibitedTagResponseBodyTag) Validate() error {
	return dara.Validate(s)
}
