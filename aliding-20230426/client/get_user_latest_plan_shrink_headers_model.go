// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserLatestPlanShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserLatestPlanShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetUserLatestPlanShrinkHeaders
	GetAccountContextShrink() *string
}

type GetUserLatestPlanShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetUserLatestPlanShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserLatestPlanShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetUserLatestPlanShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserLatestPlanShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetUserLatestPlanShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetUserLatestPlanShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserLatestPlanShrinkHeaders) SetAccountContextShrink(v string) *GetUserLatestPlanShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetUserLatestPlanShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
