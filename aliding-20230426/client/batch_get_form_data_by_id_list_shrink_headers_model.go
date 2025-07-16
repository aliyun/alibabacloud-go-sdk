// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFormDataByIdListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchGetFormDataByIdListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *BatchGetFormDataByIdListShrinkHeaders
	GetAccountContextShrink() *string
}

type BatchGetFormDataByIdListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s BatchGetFormDataByIdListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFormDataByIdListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchGetFormDataByIdListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *BatchGetFormDataByIdListShrinkHeaders) SetCommonHeaders(v map[string]*string) *BatchGetFormDataByIdListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchGetFormDataByIdListShrinkHeaders) SetAccountContextShrink(v string) *BatchGetFormDataByIdListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *BatchGetFormDataByIdListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
