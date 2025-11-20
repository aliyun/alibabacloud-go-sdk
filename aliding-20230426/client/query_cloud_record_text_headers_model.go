// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordTextHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryCloudRecordTextHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryCloudRecordTextHeadersAccountContext) *QueryCloudRecordTextHeaders
	GetAccountContext() *QueryCloudRecordTextHeadersAccountContext
}

type QueryCloudRecordTextHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryCloudRecordTextHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryCloudRecordTextHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordTextHeaders) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryCloudRecordTextHeaders) GetAccountContext() *QueryCloudRecordTextHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryCloudRecordTextHeaders) SetCommonHeaders(v map[string]*string) *QueryCloudRecordTextHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCloudRecordTextHeaders) SetAccountContext(v *QueryCloudRecordTextHeadersAccountContext) *QueryCloudRecordTextHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryCloudRecordTextHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCloudRecordTextHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryCloudRecordTextHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordTextHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryCloudRecordTextHeadersAccountContext) SetAccountId(v string) *QueryCloudRecordTextHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryCloudRecordTextHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
