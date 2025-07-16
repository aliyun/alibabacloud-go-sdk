// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMediaHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UploadMediaHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UploadMediaHeadersAccountContext) *UploadMediaHeaders
	GetAccountContext() *UploadMediaHeadersAccountContext
}

type UploadMediaHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UploadMediaHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UploadMediaHeaders) String() string {
	return dara.Prettify(s)
}

func (s UploadMediaHeaders) GoString() string {
	return s.String()
}

func (s *UploadMediaHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UploadMediaHeaders) GetAccountContext() *UploadMediaHeadersAccountContext {
	return s.AccountContext
}

func (s *UploadMediaHeaders) SetCommonHeaders(v map[string]*string) *UploadMediaHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UploadMediaHeaders) SetAccountContext(v *UploadMediaHeadersAccountContext) *UploadMediaHeaders {
	s.AccountContext = v
	return s
}

func (s *UploadMediaHeaders) Validate() error {
	return dara.Validate(s)
}

type UploadMediaHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UploadMediaHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UploadMediaHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UploadMediaHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UploadMediaHeadersAccountContext) SetAccountId(v string) *UploadMediaHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UploadMediaHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
