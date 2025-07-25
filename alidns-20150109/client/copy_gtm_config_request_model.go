// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyGtmConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCopyType(v string) *CopyGtmConfigRequest
	GetCopyType() *string
	SetLang(v string) *CopyGtmConfigRequest
	GetLang() *string
	SetSourceId(v string) *CopyGtmConfigRequest
	GetSourceId() *string
	SetTargetId(v string) *CopyGtmConfigRequest
	GetTargetId() *string
}

type CopyGtmConfigRequest struct {
	// The type of the object that is copied. Only the INSTANCE type is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	CopyType *string `json:"CopyType,omitempty" xml:"CopyType,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the source object. Only instance IDs are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// gtm-cn-0pp1j84v60d
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The ID of the target object. Only instance IDs are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// gtm-cn-v0h1gaujg06
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s CopyGtmConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyGtmConfigRequest) GoString() string {
	return s.String()
}

func (s *CopyGtmConfigRequest) GetCopyType() *string {
	return s.CopyType
}

func (s *CopyGtmConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *CopyGtmConfigRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *CopyGtmConfigRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *CopyGtmConfigRequest) SetCopyType(v string) *CopyGtmConfigRequest {
	s.CopyType = &v
	return s
}

func (s *CopyGtmConfigRequest) SetLang(v string) *CopyGtmConfigRequest {
	s.Lang = &v
	return s
}

func (s *CopyGtmConfigRequest) SetSourceId(v string) *CopyGtmConfigRequest {
	s.SourceId = &v
	return s
}

func (s *CopyGtmConfigRequest) SetTargetId(v string) *CopyGtmConfigRequest {
	s.TargetId = &v
	return s
}

func (s *CopyGtmConfigRequest) Validate() error {
	return dara.Validate(s)
}
