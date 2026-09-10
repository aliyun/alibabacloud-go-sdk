// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCronExecTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronRule(v string) *GetCronExecTimeRequest
	GetCronRule() *string
}

type GetCronExecTimeRequest struct {
	// The Cron expression. Replace spaces with plus signs `+` when passing the expression as a query parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0+0+0+*+*+?
	CronRule *string `json:"cronRule,omitempty" xml:"cronRule,omitempty"`
}

func (s GetCronExecTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCronExecTimeRequest) GoString() string {
	return s.String()
}

func (s *GetCronExecTimeRequest) GetCronRule() *string {
	return s.CronRule
}

func (s *GetCronExecTimeRequest) SetCronRule(v string) *GetCronExecTimeRequest {
	s.CronRule = &v
	return s
}

func (s *GetCronExecTimeRequest) Validate() error {
	return dara.Validate(s)
}
