// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableDataByFormInstanceIdTableIdShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListTableDataByFormInstanceIdTableIdShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListTableDataByFormInstanceIdTableIdShrinkHeaders
	GetAccountContextShrink() *string
}

type ListTableDataByFormInstanceIdTableIdShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListTableDataByFormInstanceIdTableIdShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListTableDataByFormInstanceIdTableIdShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListTableDataByFormInstanceIdTableIdShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListTableDataByFormInstanceIdTableIdShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListTableDataByFormInstanceIdTableIdShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListTableDataByFormInstanceIdTableIdShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdShrinkHeaders) SetAccountContextShrink(v string) *ListTableDataByFormInstanceIdTableIdShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
