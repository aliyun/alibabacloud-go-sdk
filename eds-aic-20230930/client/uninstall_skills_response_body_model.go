// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UninstallSkillsResponseBody
	GetCode() *string
	SetMessage(v string) *UninstallSkillsResponseBody
	GetMessage() *string
	SetRequestId(v string) *UninstallSkillsResponseBody
	GetRequestId() *string
}

type UninstallSkillsResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F07A1DA1-E1EB-5CCA-8EED-12F85D32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UninstallSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallSkillsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UninstallSkillsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UninstallSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallSkillsResponseBody) SetCode(v string) *UninstallSkillsResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallSkillsResponseBody) SetMessage(v string) *UninstallSkillsResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallSkillsResponseBody) SetRequestId(v string) *UninstallSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallSkillsResponseBody) Validate() error {
	return dara.Validate(s)
}
