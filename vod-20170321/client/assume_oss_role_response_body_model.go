// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeOssRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOssAuthResult(v string) *AssumeOssRoleResponseBody
	GetOssAuthResult() *string
	SetRequestId(v string) *AssumeOssRoleResponseBody
	GetRequestId() *string
}

type AssumeOssRoleResponseBody struct {
	OssAuthResult *string `json:"OssAuthResult,omitempty" xml:"OssAuthResult,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssumeOssRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssumeOssRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeOssRoleResponseBody) GetOssAuthResult() *string {
	return s.OssAuthResult
}

func (s *AssumeOssRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssumeOssRoleResponseBody) SetOssAuthResult(v string) *AssumeOssRoleResponseBody {
	s.OssAuthResult = &v
	return s
}

func (s *AssumeOssRoleResponseBody) SetRequestId(v string) *AssumeOssRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssumeOssRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
