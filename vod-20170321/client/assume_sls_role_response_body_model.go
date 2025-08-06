// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeSlsRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssumeSlsRoleResponseBody
	GetRequestId() *string
	SetSlsAuthResult(v string) *AssumeSlsRoleResponseBody
	GetSlsAuthResult() *string
}

type AssumeSlsRoleResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsAuthResult *string `json:"SlsAuthResult,omitempty" xml:"SlsAuthResult,omitempty"`
}

func (s AssumeSlsRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssumeSlsRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeSlsRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssumeSlsRoleResponseBody) GetSlsAuthResult() *string {
	return s.SlsAuthResult
}

func (s *AssumeSlsRoleResponseBody) SetRequestId(v string) *AssumeSlsRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssumeSlsRoleResponseBody) SetSlsAuthResult(v string) *AssumeSlsRoleResponseBody {
	s.SlsAuthResult = &v
	return s
}

func (s *AssumeSlsRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
