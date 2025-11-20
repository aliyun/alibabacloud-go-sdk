// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordVideoPlayInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryCloudRecordVideoPlayInfoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryCloudRecordVideoPlayInfoHeadersAccountContext) *QueryCloudRecordVideoPlayInfoHeaders
	GetAccountContext() *QueryCloudRecordVideoPlayInfoHeadersAccountContext
}

type QueryCloudRecordVideoPlayInfoHeaders struct {
	CommonHeaders  map[string]*string                                  `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryCloudRecordVideoPlayInfoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryCloudRecordVideoPlayInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryCloudRecordVideoPlayInfoHeaders) GetAccountContext() *QueryCloudRecordVideoPlayInfoHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryCloudRecordVideoPlayInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryCloudRecordVideoPlayInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoHeaders) SetAccountContext(v *QueryCloudRecordVideoPlayInfoHeadersAccountContext) *QueryCloudRecordVideoPlayInfoHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCloudRecordVideoPlayInfoHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryCloudRecordVideoPlayInfoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryCloudRecordVideoPlayInfoHeadersAccountContext) SetAccountId(v string) *QueryCloudRecordVideoPlayInfoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
