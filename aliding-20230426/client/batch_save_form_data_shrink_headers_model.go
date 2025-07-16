// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSaveFormDataShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchSaveFormDataShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *BatchSaveFormDataShrinkHeaders
	GetAccountContextShrink() *string
}

type BatchSaveFormDataShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s BatchSaveFormDataShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchSaveFormDataShrinkHeaders) GoString() string {
	return s.String()
}

func (s *BatchSaveFormDataShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchSaveFormDataShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *BatchSaveFormDataShrinkHeaders) SetCommonHeaders(v map[string]*string) *BatchSaveFormDataShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchSaveFormDataShrinkHeaders) SetAccountContextShrink(v string) *BatchSaveFormDataShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *BatchSaveFormDataShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
