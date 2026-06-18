// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferCallToSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TransferCallToSkillGroupResponseBody
	GetCode() *string
	SetMessage(v string) *TransferCallToSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *TransferCallToSkillGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TransferCallToSkillGroupResponseBody
	GetSuccess() *bool
}

type TransferCallToSkillGroupResponseBody struct {
	// Error encoding.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferCallToSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferCallToSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *TransferCallToSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *TransferCallToSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TransferCallToSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferCallToSkillGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TransferCallToSkillGroupResponseBody) SetCode(v string) *TransferCallToSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *TransferCallToSkillGroupResponseBody) SetMessage(v string) *TransferCallToSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *TransferCallToSkillGroupResponseBody) SetRequestId(v string) *TransferCallToSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferCallToSkillGroupResponseBody) SetSuccess(v bool) *TransferCallToSkillGroupResponseBody {
	s.Success = &v
	return s
}

func (s *TransferCallToSkillGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
