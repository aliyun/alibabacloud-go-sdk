// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetUserIdByOpenDingtalkIdShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchGetUserIdByOpenDingtalkIdShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *BatchGetUserIdByOpenDingtalkIdShrinkHeaders
	GetAccountContextShrink() *string
}

type BatchGetUserIdByOpenDingtalkIdShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s BatchGetUserIdByOpenDingtalkIdShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchGetUserIdByOpenDingtalkIdShrinkHeaders) GoString() string {
	return s.String()
}

func (s *BatchGetUserIdByOpenDingtalkIdShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchGetUserIdByOpenDingtalkIdShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *BatchGetUserIdByOpenDingtalkIdShrinkHeaders) SetCommonHeaders(v map[string]*string) *BatchGetUserIdByOpenDingtalkIdShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdShrinkHeaders) SetAccountContextShrink(v string) *BatchGetUserIdByOpenDingtalkIdShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
