// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRobotToScenegroupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddRobotToScenegroupHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AddRobotToScenegroupHeadersAccountContext) *AddRobotToScenegroupHeaders
	GetAccountContext() *AddRobotToScenegroupHeadersAccountContext
}

type AddRobotToScenegroupHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddRobotToScenegroupHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddRobotToScenegroupHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddRobotToScenegroupHeaders) GoString() string {
	return s.String()
}

func (s *AddRobotToScenegroupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddRobotToScenegroupHeaders) GetAccountContext() *AddRobotToScenegroupHeadersAccountContext {
	return s.AccountContext
}

func (s *AddRobotToScenegroupHeaders) SetCommonHeaders(v map[string]*string) *AddRobotToScenegroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddRobotToScenegroupHeaders) SetAccountContext(v *AddRobotToScenegroupHeadersAccountContext) *AddRobotToScenegroupHeaders {
	s.AccountContext = v
	return s
}

func (s *AddRobotToScenegroupHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddRobotToScenegroupHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddRobotToScenegroupHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AddRobotToScenegroupHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddRobotToScenegroupHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AddRobotToScenegroupHeadersAccountContext) SetAccountId(v string) *AddRobotToScenegroupHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AddRobotToScenegroupHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
