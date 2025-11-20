// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteLiveHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteLiveHeadersAccountContext) *DeleteLiveHeaders
	GetAccountContext() *DeleteLiveHeadersAccountContext
}

type DeleteLiveHeaders struct {
	CommonHeaders  map[string]*string               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteLiveHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteLiveHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveHeaders) GoString() string {
	return s.String()
}

func (s *DeleteLiveHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteLiveHeaders) GetAccountContext() *DeleteLiveHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteLiveHeaders) SetCommonHeaders(v map[string]*string) *DeleteLiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteLiveHeaders) SetAccountContext(v *DeleteLiveHeadersAccountContext) *DeleteLiveHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteLiveHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteLiveHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteLiveHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteLiveHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteLiveHeadersAccountContext) SetAccountId(v string) *DeleteLiveHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteLiveHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
