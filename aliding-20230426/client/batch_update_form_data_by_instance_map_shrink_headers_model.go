// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFormDataByInstanceMapShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceMapShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *BatchUpdateFormDataByInstanceMapShrinkHeaders
	GetAccountContextShrink() *string
}

type BatchUpdateFormDataByInstanceMapShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s BatchUpdateFormDataByInstanceMapShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceMapShrinkHeaders) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceMapShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchUpdateFormDataByInstanceMapShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *BatchUpdateFormDataByInstanceMapShrinkHeaders) SetCommonHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceMapShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapShrinkHeaders) SetAccountContextShrink(v string) *BatchUpdateFormDataByInstanceMapShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
