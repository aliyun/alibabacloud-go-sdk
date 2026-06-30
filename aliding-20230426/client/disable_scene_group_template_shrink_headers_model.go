// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSceneGroupTemplateShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DisableSceneGroupTemplateShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DisableSceneGroupTemplateShrinkHeaders
	GetAccountContextShrink() *string
}

type DisableSceneGroupTemplateShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DisableSceneGroupTemplateShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DisableSceneGroupTemplateShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DisableSceneGroupTemplateShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DisableSceneGroupTemplateShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DisableSceneGroupTemplateShrinkHeaders) SetCommonHeaders(v map[string]*string) *DisableSceneGroupTemplateShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DisableSceneGroupTemplateShrinkHeaders) SetAccountContextShrink(v string) *DisableSceneGroupTemplateShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DisableSceneGroupTemplateShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
