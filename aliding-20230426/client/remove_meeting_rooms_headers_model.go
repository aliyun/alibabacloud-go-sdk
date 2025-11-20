// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveMeetingRoomsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RemoveMeetingRoomsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *RemoveMeetingRoomsHeadersAccountContext) *RemoveMeetingRoomsHeaders
	GetAccountContext() *RemoveMeetingRoomsHeadersAccountContext
}

type RemoveMeetingRoomsHeaders struct {
	CommonHeaders  map[string]*string                       `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *RemoveMeetingRoomsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s RemoveMeetingRoomsHeaders) String() string {
	return dara.Prettify(s)
}

func (s RemoveMeetingRoomsHeaders) GoString() string {
	return s.String()
}

func (s *RemoveMeetingRoomsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RemoveMeetingRoomsHeaders) GetAccountContext() *RemoveMeetingRoomsHeadersAccountContext {
	return s.AccountContext
}

func (s *RemoveMeetingRoomsHeaders) SetCommonHeaders(v map[string]*string) *RemoveMeetingRoomsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveMeetingRoomsHeaders) SetAccountContext(v *RemoveMeetingRoomsHeadersAccountContext) *RemoveMeetingRoomsHeaders {
	s.AccountContext = v
	return s
}

func (s *RemoveMeetingRoomsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveMeetingRoomsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s RemoveMeetingRoomsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s RemoveMeetingRoomsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *RemoveMeetingRoomsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *RemoveMeetingRoomsHeadersAccountContext) SetAccountId(v string) *RemoveMeetingRoomsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *RemoveMeetingRoomsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
