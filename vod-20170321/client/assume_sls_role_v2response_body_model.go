// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeSlsRoleV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssumeSlsRoleV2ResponseBody
	GetRequestId() *string
	SetSlsAuthResult(v string) *AssumeSlsRoleV2ResponseBody
	GetSlsAuthResult() *string
}

type AssumeSlsRoleV2ResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsAuthResult *string `json:"SlsAuthResult,omitempty" xml:"SlsAuthResult,omitempty"`
}

func (s AssumeSlsRoleV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssumeSlsRoleV2ResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeSlsRoleV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssumeSlsRoleV2ResponseBody) GetSlsAuthResult() *string {
	return s.SlsAuthResult
}

func (s *AssumeSlsRoleV2ResponseBody) SetRequestId(v string) *AssumeSlsRoleV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssumeSlsRoleV2ResponseBody) SetSlsAuthResult(v string) *AssumeSlsRoleV2ResponseBody {
	s.SlsAuthResult = &v
	return s
}

func (s *AssumeSlsRoleV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
