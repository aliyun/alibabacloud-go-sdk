// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentName(v string) *CreateEnvironmentRequest
	GetEnvironmentName() *string
	SetNextEnvironmentName(v string) *CreateEnvironmentRequest
	GetNextEnvironmentName() *string
	SetRule(v string) *CreateEnvironmentRequest
	GetRule() *string
	SetSiteId(v int64) *CreateEnvironmentRequest
	GetSiteId() *int64
}

type CreateEnvironmentRequest struct {
	// The environment name.
	//
	// This parameter is required.
	//
	// example:
	//
	// NPDcP1
	EnvironmentName *string `json:"EnvironmentName,omitempty" xml:"EnvironmentName,omitempty"`
	// The name of the environment with the next priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 默认环境
	NextEnvironmentName *string `json:"NextEnvironmentName,omitempty" xml:"NextEnvironmentName,omitempty"`
	// The environment rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// ("ip" eq "1.1.1.1")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123**
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentRequest) GetEnvironmentName() *string {
	return s.EnvironmentName
}

func (s *CreateEnvironmentRequest) GetNextEnvironmentName() *string {
	return s.NextEnvironmentName
}

func (s *CreateEnvironmentRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateEnvironmentRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateEnvironmentRequest) SetEnvironmentName(v string) *CreateEnvironmentRequest {
	s.EnvironmentName = &v
	return s
}

func (s *CreateEnvironmentRequest) SetNextEnvironmentName(v string) *CreateEnvironmentRequest {
	s.NextEnvironmentName = &v
	return s
}

func (s *CreateEnvironmentRequest) SetRule(v string) *CreateEnvironmentRequest {
	s.Rule = &v
	return s
}

func (s *CreateEnvironmentRequest) SetSiteId(v int64) *CreateEnvironmentRequest {
	s.SiteId = &v
	return s
}

func (s *CreateEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}
