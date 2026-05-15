// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteAllHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MuteAllHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *MuteAllHeadersAccountContext) *MuteAllHeaders
	GetAccountContext() *MuteAllHeadersAccountContext
}

type MuteAllHeaders struct {
	CommonHeaders  map[string]*string            `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *MuteAllHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s MuteAllHeaders) String() string {
	return dara.Prettify(s)
}

func (s MuteAllHeaders) GoString() string {
	return s.String()
}

func (s *MuteAllHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MuteAllHeaders) GetAccountContext() *MuteAllHeadersAccountContext {
	return s.AccountContext
}

func (s *MuteAllHeaders) SetCommonHeaders(v map[string]*string) *MuteAllHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MuteAllHeaders) SetAccountContext(v *MuteAllHeadersAccountContext) *MuteAllHeaders {
	s.AccountContext = v
	return s
}

func (s *MuteAllHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MuteAllHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s MuteAllHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s MuteAllHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *MuteAllHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *MuteAllHeadersAccountContext) SetAccountId(v string) *MuteAllHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *MuteAllHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
