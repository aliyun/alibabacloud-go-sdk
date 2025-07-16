// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceInfoByRoomCodeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryConferenceInfoByRoomCodeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryConferenceInfoByRoomCodeHeadersAccountContext) *QueryConferenceInfoByRoomCodeHeaders
	GetAccountContext() *QueryConferenceInfoByRoomCodeHeadersAccountContext
}

type QueryConferenceInfoByRoomCodeHeaders struct {
	CommonHeaders  map[string]*string                                  `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryConferenceInfoByRoomCodeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryConferenceInfoByRoomCodeHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeHeaders) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryConferenceInfoByRoomCodeHeaders) GetAccountContext() *QueryConferenceInfoByRoomCodeHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryConferenceInfoByRoomCodeHeaders) SetCommonHeaders(v map[string]*string) *QueryConferenceInfoByRoomCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConferenceInfoByRoomCodeHeaders) SetAccountContext(v *QueryConferenceInfoByRoomCodeHeadersAccountContext) *QueryConferenceInfoByRoomCodeHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryConferenceInfoByRoomCodeHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryConferenceInfoByRoomCodeHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryConferenceInfoByRoomCodeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryConferenceInfoByRoomCodeHeadersAccountContext) SetAccountId(v string) *QueryConferenceInfoByRoomCodeHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
