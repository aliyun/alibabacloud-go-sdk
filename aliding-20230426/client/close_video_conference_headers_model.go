// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseVideoConferenceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CloseVideoConferenceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CloseVideoConferenceHeadersAccountContext) *CloseVideoConferenceHeaders
	GetAccountContext() *CloseVideoConferenceHeadersAccountContext
}

type CloseVideoConferenceHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CloseVideoConferenceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CloseVideoConferenceHeaders) String() string {
	return dara.Prettify(s)
}

func (s CloseVideoConferenceHeaders) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CloseVideoConferenceHeaders) GetAccountContext() *CloseVideoConferenceHeadersAccountContext {
	return s.AccountContext
}

func (s *CloseVideoConferenceHeaders) SetCommonHeaders(v map[string]*string) *CloseVideoConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CloseVideoConferenceHeaders) SetAccountContext(v *CloseVideoConferenceHeadersAccountContext) *CloseVideoConferenceHeaders {
	s.AccountContext = v
	return s
}

func (s *CloseVideoConferenceHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloseVideoConferenceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CloseVideoConferenceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CloseVideoConferenceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CloseVideoConferenceHeadersAccountContext) SetAccountId(v string) *CloseVideoConferenceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CloseVideoConferenceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
