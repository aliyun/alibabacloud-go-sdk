// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserTagValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagId(v string) *UpdateUserTagValueRequest
	GetTagId() *string
	SetTagValue(v string) *UpdateUserTagValueRequest
	GetTagValue() *string
	SetUserId(v string) *UpdateUserTagValueRequest
	GetUserId() *string
}

type UpdateUserTagValueRequest struct {
	// The ID of the tag to be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// pop_001
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The tag value to be modified.
	//
	// - To clear this tag, set the tag value to ($NULL$).
	//
	// - For multiple values, use English commas to separate them.
	//
	// - Format validation, maximum length: 3000 characters
	//
	// This parameter is required.
	//
	// example:
	//
	// Product Director
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// The user ID for which the tag value is to be modified. This user ID refers to the Quick BI UserID, not the Alibaba Cloud UID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateUserTagValueRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserTagValueRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserTagValueRequest) GetTagId() *string {
	return s.TagId
}

func (s *UpdateUserTagValueRequest) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateUserTagValueRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserTagValueRequest) SetTagId(v string) *UpdateUserTagValueRequest {
	s.TagId = &v
	return s
}

func (s *UpdateUserTagValueRequest) SetTagValue(v string) *UpdateUserTagValueRequest {
	s.TagValue = &v
	return s
}

func (s *UpdateUserTagValueRequest) SetUserId(v string) *UpdateUserTagValueRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserTagValueRequest) Validate() error {
	return dara.Validate(s)
}
