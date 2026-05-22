// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByOpenDingtalkIdShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserIdByOpenDingtalkIdShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetUserIdByOpenDingtalkIdShrinkHeaders
	GetAccountContextShrink() *string
}

type GetUserIdByOpenDingtalkIdShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetUserIdByOpenDingtalkIdShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOpenDingtalkIdShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByOpenDingtalkIdShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserIdByOpenDingtalkIdShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetUserIdByOpenDingtalkIdShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByOpenDingtalkIdShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByOpenDingtalkIdShrinkHeaders) SetAccountContextShrink(v string) *GetUserIdByOpenDingtalkIdShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetUserIdByOpenDingtalkIdShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
