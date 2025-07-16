// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordVideoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryCloudRecordVideoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryCloudRecordVideoHeadersAccountContext) *QueryCloudRecordVideoHeaders
	GetAccountContext() *QueryCloudRecordVideoHeadersAccountContext
}

type QueryCloudRecordVideoHeaders struct {
	CommonHeaders  map[string]*string                          `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryCloudRecordVideoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryCloudRecordVideoHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoHeaders) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryCloudRecordVideoHeaders) GetAccountContext() *QueryCloudRecordVideoHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryCloudRecordVideoHeaders) SetCommonHeaders(v map[string]*string) *QueryCloudRecordVideoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCloudRecordVideoHeaders) SetAccountContext(v *QueryCloudRecordVideoHeadersAccountContext) *QueryCloudRecordVideoHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryCloudRecordVideoHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryCloudRecordVideoHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryCloudRecordVideoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryCloudRecordVideoHeadersAccountContext) SetAccountId(v string) *QueryCloudRecordVideoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryCloudRecordVideoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
