// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetSubTaskResultRequest
	GetLang() *string
	SetRegId(v string) *GetSubTaskResultRequest
	GetRegId() *string
	SetSubTaskId(v int32) *GetSubTaskResultRequest
	GetSubTaskId() *int32
}

type GetSubTaskResultRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// example:
	//
	// 2
	SubTaskId *int32 `json:"SubTaskId,omitempty" xml:"SubTaskId,omitempty"`
}

func (s GetSubTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetSubTaskResultRequest) GetLang() *string {
	return s.Lang
}

func (s *GetSubTaskResultRequest) GetRegId() *string {
	return s.RegId
}

func (s *GetSubTaskResultRequest) GetSubTaskId() *int32 {
	return s.SubTaskId
}

func (s *GetSubTaskResultRequest) SetLang(v string) *GetSubTaskResultRequest {
	s.Lang = &v
	return s
}

func (s *GetSubTaskResultRequest) SetRegId(v string) *GetSubTaskResultRequest {
	s.RegId = &v
	return s
}

func (s *GetSubTaskResultRequest) SetSubTaskId(v int32) *GetSubTaskResultRequest {
	s.SubTaskId = &v
	return s
}

func (s *GetSubTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
