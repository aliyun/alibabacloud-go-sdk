// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetOperationRecordsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetOperationRecordsShrinkHeaders
	GetAccountContextShrink() *string
}

type GetOperationRecordsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetOperationRecordsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetOperationRecordsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetOperationRecordsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetOperationRecordsShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetOperationRecordsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOperationRecordsShrinkHeaders) SetAccountContextShrink(v string) *GetOperationRecordsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetOperationRecordsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
