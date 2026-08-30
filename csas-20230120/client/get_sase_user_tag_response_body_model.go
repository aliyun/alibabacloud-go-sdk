// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSaseUserTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSaseUserTagResponseBody
	GetRequestId() *string
	SetSaseUserTag(v *GetSaseUserTagResponseBodySaseUserTag) *GetSaseUserTagResponseBody
	GetSaseUserTag() *GetSaseUserTagResponseBodySaseUserTag
}

type GetSaseUserTagResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2CABFEBB-0CE7-575E-833A-266F75D46713
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user tag response body.
	SaseUserTag *GetSaseUserTagResponseBodySaseUserTag `json:"SaseUserTag,omitempty" xml:"SaseUserTag,omitempty" type:"Struct"`
}

func (s GetSaseUserTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSaseUserTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetSaseUserTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSaseUserTagResponseBody) GetSaseUserTag() *GetSaseUserTagResponseBodySaseUserTag {
	return s.SaseUserTag
}

func (s *GetSaseUserTagResponseBody) SetRequestId(v string) *GetSaseUserTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSaseUserTagResponseBody) SetSaseUserTag(v *GetSaseUserTagResponseBodySaseUserTag) *GetSaseUserTagResponseBody {
	s.SaseUserTag = v
	return s
}

func (s *GetSaseUserTagResponseBody) Validate() error {
	if s.SaseUserTag != nil {
		if err := s.SaseUserTag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSaseUserTagResponseBodySaseUserTag struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 141681795035****
	Aliuid *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The user tag description.
	//
	// example:
	//
	// These are the company\\"s employees
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user tag name.
	//
	// example:
	//
	// boss
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user tag ID.
	//
	// example:
	//
	// su-tag-1ae52f66039fa0d4****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s GetSaseUserTagResponseBodySaseUserTag) String() string {
	return dara.Prettify(s)
}

func (s GetSaseUserTagResponseBodySaseUserTag) GoString() string {
	return s.String()
}

func (s *GetSaseUserTagResponseBodySaseUserTag) GetAliuid() *string {
	return s.Aliuid
}

func (s *GetSaseUserTagResponseBodySaseUserTag) GetDescription() *string {
	return s.Description
}

func (s *GetSaseUserTagResponseBodySaseUserTag) GetName() *string {
	return s.Name
}

func (s *GetSaseUserTagResponseBodySaseUserTag) GetTagId() *string {
	return s.TagId
}

func (s *GetSaseUserTagResponseBodySaseUserTag) SetAliuid(v string) *GetSaseUserTagResponseBodySaseUserTag {
	s.Aliuid = &v
	return s
}

func (s *GetSaseUserTagResponseBodySaseUserTag) SetDescription(v string) *GetSaseUserTagResponseBodySaseUserTag {
	s.Description = &v
	return s
}

func (s *GetSaseUserTagResponseBodySaseUserTag) SetName(v string) *GetSaseUserTagResponseBodySaseUserTag {
	s.Name = &v
	return s
}

func (s *GetSaseUserTagResponseBodySaseUserTag) SetTagId(v string) *GetSaseUserTagResponseBodySaseUserTag {
	s.TagId = &v
	return s
}

func (s *GetSaseUserTagResponseBodySaseUserTag) Validate() error {
	return dara.Validate(s)
}
