// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRCSSignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSignName(v string) *GetRCSSignatureRequest
	GetSignName() *string
}

type GetRCSSignatureRequest struct {
	// 签名名称
	//
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s GetRCSSignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRCSSignatureRequest) GoString() string {
	return s.String()
}

func (s *GetRCSSignatureRequest) GetSignName() *string {
	return s.SignName
}

func (s *GetRCSSignatureRequest) SetSignName(v string) *GetRCSSignatureRequest {
	s.SignName = &v
	return s
}

func (s *GetRCSSignatureRequest) Validate() error {
	return dara.Validate(s)
}
