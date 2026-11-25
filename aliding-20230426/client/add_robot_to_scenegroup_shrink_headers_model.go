// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRobotToScenegroupShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddRobotToScenegroupShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AddRobotToScenegroupShrinkHeaders
	GetAccountContextShrink() *string
}

type AddRobotToScenegroupShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddRobotToScenegroupShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddRobotToScenegroupShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddRobotToScenegroupShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddRobotToScenegroupShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AddRobotToScenegroupShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddRobotToScenegroupShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddRobotToScenegroupShrinkHeaders) SetAccountContextShrink(v string) *AddRobotToScenegroupShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AddRobotToScenegroupShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
