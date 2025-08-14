// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocParserResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetDocParserResultRequest
	GetId() *string
	SetLayoutNum(v int32) *GetDocParserResultRequest
	GetLayoutNum() *int32
	SetLayoutStepSize(v int32) *GetDocParserResultRequest
	GetLayoutStepSize() *int32
}

type GetDocParserResultRequest struct {
	// example:
	//
	// docmind-20220816-1e89d65c
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	LayoutNum      *int32  `json:"LayoutNum,omitempty" xml:"LayoutNum,omitempty"`
	LayoutStepSize *int32  `json:"LayoutStepSize,omitempty" xml:"LayoutStepSize,omitempty"`
}

func (s GetDocParserResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocParserResultRequest) GoString() string {
	return s.String()
}

func (s *GetDocParserResultRequest) GetId() *string {
	return s.Id
}

func (s *GetDocParserResultRequest) GetLayoutNum() *int32 {
	return s.LayoutNum
}

func (s *GetDocParserResultRequest) GetLayoutStepSize() *int32 {
	return s.LayoutStepSize
}

func (s *GetDocParserResultRequest) SetId(v string) *GetDocParserResultRequest {
	s.Id = &v
	return s
}

func (s *GetDocParserResultRequest) SetLayoutNum(v int32) *GetDocParserResultRequest {
	s.LayoutNum = &v
	return s
}

func (s *GetDocParserResultRequest) SetLayoutStepSize(v int32) *GetDocParserResultRequest {
	s.LayoutStepSize = &v
	return s
}

func (s *GetDocParserResultRequest) Validate() error {
	return dara.Validate(s)
}
