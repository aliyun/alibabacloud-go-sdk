// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearDataShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ClearDataShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ClearDataShrinkHeaders
	GetAccountContextShrink() *string
}

type ClearDataShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ClearDataShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ClearDataShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ClearDataShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ClearDataShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ClearDataShrinkHeaders) SetCommonHeaders(v map[string]*string) *ClearDataShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ClearDataShrinkHeaders) SetAccountContextShrink(v string) *ClearDataShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ClearDataShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
