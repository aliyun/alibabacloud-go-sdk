// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchEventHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PatchEventHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *PatchEventHeadersAccountContext) *PatchEventHeaders
	GetAccountContext() *PatchEventHeadersAccountContext
}

type PatchEventHeaders struct {
	CommonHeaders  map[string]*string               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *PatchEventHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s PatchEventHeaders) String() string {
	return dara.Prettify(s)
}

func (s PatchEventHeaders) GoString() string {
	return s.String()
}

func (s *PatchEventHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PatchEventHeaders) GetAccountContext() *PatchEventHeadersAccountContext {
	return s.AccountContext
}

func (s *PatchEventHeaders) SetCommonHeaders(v map[string]*string) *PatchEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PatchEventHeaders) SetAccountContext(v *PatchEventHeadersAccountContext) *PatchEventHeaders {
	s.AccountContext = v
	return s
}

func (s *PatchEventHeaders) Validate() error {
	return dara.Validate(s)
}

type PatchEventHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s PatchEventHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s PatchEventHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *PatchEventHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *PatchEventHeadersAccountContext) SetAccountId(v string) *PatchEventHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *PatchEventHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
