// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddApListToApgroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApGroupId(v string) *AddApListToApgroupShrinkRequest
	GetApGroupId() *string
	SetApMacListShrink(v string) *AddApListToApgroupShrinkRequest
	GetApMacListShrink() *string
	SetAppCode(v string) *AddApListToApgroupShrinkRequest
	GetAppCode() *string
	SetAppName(v string) *AddApListToApgroupShrinkRequest
	GetAppName() *string
}

type AddApListToApgroupShrinkRequest struct {
	// This parameter is required.
	ApGroupId *string `json:"ApGroupId,omitempty" xml:"ApGroupId,omitempty"`
	// This parameter is required.
	ApMacListShrink *string `json:"ApMacList,omitempty" xml:"ApMacList,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s AddApListToApgroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddApListToApgroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddApListToApgroupShrinkRequest) GetApGroupId() *string {
	return s.ApGroupId
}

func (s *AddApListToApgroupShrinkRequest) GetApMacListShrink() *string {
	return s.ApMacListShrink
}

func (s *AddApListToApgroupShrinkRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *AddApListToApgroupShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddApListToApgroupShrinkRequest) SetApGroupId(v string) *AddApListToApgroupShrinkRequest {
	s.ApGroupId = &v
	return s
}

func (s *AddApListToApgroupShrinkRequest) SetApMacListShrink(v string) *AddApListToApgroupShrinkRequest {
	s.ApMacListShrink = &v
	return s
}

func (s *AddApListToApgroupShrinkRequest) SetAppCode(v string) *AddApListToApgroupShrinkRequest {
	s.AppCode = &v
	return s
}

func (s *AddApListToApgroupShrinkRequest) SetAppName(v string) *AddApListToApgroupShrinkRequest {
	s.AppName = &v
	return s
}

func (s *AddApListToApgroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
