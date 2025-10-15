// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationTokens(v *CreateApplicationTokenResponseBodyApplicationTokens) *CreateApplicationTokenResponseBody
	GetApplicationTokens() *CreateApplicationTokenResponseBodyApplicationTokens
	SetRequestId(v string) *CreateApplicationTokenResponseBody
	GetRequestId() *string
}

type CreateApplicationTokenResponseBody struct {
	ApplicationTokens *CreateApplicationTokenResponseBodyApplicationTokens `json:"ApplicationTokens,omitempty" xml:"ApplicationTokens,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationTokenResponseBody) GetApplicationTokens() *CreateApplicationTokenResponseBodyApplicationTokens {
	return s.ApplicationTokens
}

func (s *CreateApplicationTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationTokenResponseBody) SetApplicationTokens(v *CreateApplicationTokenResponseBodyApplicationTokens) *CreateApplicationTokenResponseBody {
	s.ApplicationTokens = v
	return s
}

func (s *CreateApplicationTokenResponseBody) SetRequestId(v string) *CreateApplicationTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationTokenResponseBody) Validate() error {
	if s.ApplicationTokens != nil {
		if err := s.ApplicationTokens.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApplicationTokenResponseBodyApplicationTokens struct {
	// 应用token
	//
	// example:
	//
	// SATFwqX8zxGf83pJcJw78KFGjmrft4erWeZYBGS8oE7NN6qoE217yaJpUdMb1UuuGqhDiF43sCA4CF91CTL5iGntqwyLuaAcS9FJ9HfGadE5a7TjiwVafwrBYkt3XXX
	ApplicationToken *string `json:"ApplicationToken,omitempty" xml:"ApplicationToken,omitempty"`
	// 应用token ID
	//
	// example:
	//
	// token_ndfxxigahelfne2y2hodehrxxxx
	ApplicationTokenId *string `json:"ApplicationTokenId,omitempty" xml:"ApplicationTokenId,omitempty"`
	// 应用token类型
	//
	// example:
	//
	// bearer_token
	ApplicationTokenType *string `json:"ApplicationTokenType,omitempty" xml:"ApplicationTokenType,omitempty"`
}

func (s CreateApplicationTokenResponseBodyApplicationTokens) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationTokenResponseBodyApplicationTokens) GoString() string {
	return s.String()
}

func (s *CreateApplicationTokenResponseBodyApplicationTokens) GetApplicationToken() *string {
	return s.ApplicationToken
}

func (s *CreateApplicationTokenResponseBodyApplicationTokens) GetApplicationTokenId() *string {
	return s.ApplicationTokenId
}

func (s *CreateApplicationTokenResponseBodyApplicationTokens) GetApplicationTokenType() *string {
	return s.ApplicationTokenType
}

func (s *CreateApplicationTokenResponseBodyApplicationTokens) SetApplicationToken(v string) *CreateApplicationTokenResponseBodyApplicationTokens {
	s.ApplicationToken = &v
	return s
}

func (s *CreateApplicationTokenResponseBodyApplicationTokens) SetApplicationTokenId(v string) *CreateApplicationTokenResponseBodyApplicationTokens {
	s.ApplicationTokenId = &v
	return s
}

func (s *CreateApplicationTokenResponseBodyApplicationTokens) SetApplicationTokenType(v string) *CreateApplicationTokenResponseBodyApplicationTokens {
	s.ApplicationTokenType = &v
	return s
}

func (s *CreateApplicationTokenResponseBodyApplicationTokens) Validate() error {
	return dara.Validate(s)
}
