// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveReplayUrlHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetLiveReplayUrlHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetLiveReplayUrlHeadersAccountContext) *GetLiveReplayUrlHeaders
	GetAccountContext() *GetLiveReplayUrlHeadersAccountContext
}

type GetLiveReplayUrlHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetLiveReplayUrlHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetLiveReplayUrlHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetLiveReplayUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetLiveReplayUrlHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetLiveReplayUrlHeaders) GetAccountContext() *GetLiveReplayUrlHeadersAccountContext {
	return s.AccountContext
}

func (s *GetLiveReplayUrlHeaders) SetCommonHeaders(v map[string]*string) *GetLiveReplayUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLiveReplayUrlHeaders) SetAccountContext(v *GetLiveReplayUrlHeadersAccountContext) *GetLiveReplayUrlHeaders {
	s.AccountContext = v
	return s
}

func (s *GetLiveReplayUrlHeaders) Validate() error {
	return dara.Validate(s)
}

type GetLiveReplayUrlHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetLiveReplayUrlHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetLiveReplayUrlHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetLiveReplayUrlHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetLiveReplayUrlHeadersAccountContext) SetAccountId(v string) *GetLiveReplayUrlHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetLiveReplayUrlHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
