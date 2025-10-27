// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocParserResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeFieldsShrink(v string) *GetDocParserResultShrinkRequest
	GetExcludeFieldsShrink() *string
	SetId(v string) *GetDocParserResultShrinkRequest
	GetId() *string
	SetLayoutNum(v int32) *GetDocParserResultShrinkRequest
	GetLayoutNum() *int32
	SetLayoutStepSize(v int32) *GetDocParserResultShrinkRequest
	GetLayoutStepSize() *int32
}

type GetDocParserResultShrinkRequest struct {
	ExcludeFieldsShrink *string `json:"ExcludeFields,omitempty" xml:"ExcludeFields,omitempty"`
	// example:
	//
	// docmind-20220816-1e89d65c
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	LayoutNum      *int32  `json:"LayoutNum,omitempty" xml:"LayoutNum,omitempty"`
	LayoutStepSize *int32  `json:"LayoutStepSize,omitempty" xml:"LayoutStepSize,omitempty"`
}

func (s GetDocParserResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocParserResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDocParserResultShrinkRequest) GetExcludeFieldsShrink() *string {
	return s.ExcludeFieldsShrink
}

func (s *GetDocParserResultShrinkRequest) GetId() *string {
	return s.Id
}

func (s *GetDocParserResultShrinkRequest) GetLayoutNum() *int32 {
	return s.LayoutNum
}

func (s *GetDocParserResultShrinkRequest) GetLayoutStepSize() *int32 {
	return s.LayoutStepSize
}

func (s *GetDocParserResultShrinkRequest) SetExcludeFieldsShrink(v string) *GetDocParserResultShrinkRequest {
	s.ExcludeFieldsShrink = &v
	return s
}

func (s *GetDocParserResultShrinkRequest) SetId(v string) *GetDocParserResultShrinkRequest {
	s.Id = &v
	return s
}

func (s *GetDocParserResultShrinkRequest) SetLayoutNum(v int32) *GetDocParserResultShrinkRequest {
	s.LayoutNum = &v
	return s
}

func (s *GetDocParserResultShrinkRequest) SetLayoutStepSize(v int32) *GetDocParserResultShrinkRequest {
	s.LayoutStepSize = &v
	return s
}

func (s *GetDocParserResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
