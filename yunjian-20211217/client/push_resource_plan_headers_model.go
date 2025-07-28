// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushResourcePlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PushResourcePlanHeaders
	GetCommonHeaders() map[string]*string
	SetYunUserId(v string) *PushResourcePlanHeaders
	GetYunUserId() *string
}

type PushResourcePlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s PushResourcePlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s PushResourcePlanHeaders) GoString() string {
	return s.String()
}

func (s *PushResourcePlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PushResourcePlanHeaders) GetYunUserId() *string {
	return s.YunUserId
}

func (s *PushResourcePlanHeaders) SetCommonHeaders(v map[string]*string) *PushResourcePlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushResourcePlanHeaders) SetYunUserId(v string) *PushResourcePlanHeaders {
	s.YunUserId = &v
	return s
}

func (s *PushResourcePlanHeaders) Validate() error {
	return dara.Validate(s)
}
