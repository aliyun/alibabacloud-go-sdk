// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTrafficControlTaskCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateTrafficControlTaskCodeResponseBody
	GetCode() *string
	SetPreNeedConfig(v bool) *GenerateTrafficControlTaskCodeResponseBody
	GetPreNeedConfig() *bool
	SetRequestId(v string) *GenerateTrafficControlTaskCodeResponseBody
	GetRequestId() *string
}

type GenerateTrafficControlTaskCodeResponseBody struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	PreNeedConfig *bool   `json:"PreNeedConfig,omitempty" xml:"PreNeedConfig,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateTrafficControlTaskCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateTrafficControlTaskCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTrafficControlTaskCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateTrafficControlTaskCodeResponseBody) GetPreNeedConfig() *bool {
	return s.PreNeedConfig
}

func (s *GenerateTrafficControlTaskCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateTrafficControlTaskCodeResponseBody) SetCode(v string) *GenerateTrafficControlTaskCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateTrafficControlTaskCodeResponseBody) SetPreNeedConfig(v bool) *GenerateTrafficControlTaskCodeResponseBody {
	s.PreNeedConfig = &v
	return s
}

func (s *GenerateTrafficControlTaskCodeResponseBody) SetRequestId(v string) *GenerateTrafficControlTaskCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateTrafficControlTaskCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
