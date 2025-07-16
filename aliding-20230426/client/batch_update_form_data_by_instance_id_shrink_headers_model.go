// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFormDataByInstanceIdShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceIdShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *BatchUpdateFormDataByInstanceIdShrinkHeaders
	GetAccountContextShrink() *string
}

type BatchUpdateFormDataByInstanceIdShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s BatchUpdateFormDataByInstanceIdShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceIdShrinkHeaders) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceIdShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchUpdateFormDataByInstanceIdShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *BatchUpdateFormDataByInstanceIdShrinkHeaders) SetCommonHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceIdShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdShrinkHeaders) SetAccountContextShrink(v string) *BatchUpdateFormDataByInstanceIdShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
