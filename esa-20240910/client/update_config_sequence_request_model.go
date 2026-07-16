// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigSequenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateConfigSequenceRequest
	GetConfigId() *int64
	SetSequence(v int32) *UpdateConfigSequenceRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateConfigSequenceRequest
	GetSiteId() *int64
}

type UpdateConfigSequenceRequest struct {
	// The configuration ID. You can call the [ListSiteFunctions](~~ListSiteFunctions~~) operation to obtain the configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27490172680****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The target priority of the configuration. The value must be greater than 0. If the value is greater than the highest priority among all rule configurations under this feature, the priority of the configuration to be modified is set to the current highest priority. For example, if the CacheRules feature has three rule configurations with priorities 1, 2, and 3, and you change the priority of the rule with priority 1 to 5, the priority of that rule is set to 3, and the rules that originally had priorities 2 and 3 are changed to 1 and 2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 611133661****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateConfigSequenceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigSequenceRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigSequenceRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateConfigSequenceRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateConfigSequenceRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateConfigSequenceRequest) SetConfigId(v int64) *UpdateConfigSequenceRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateConfigSequenceRequest) SetSequence(v int32) *UpdateConfigSequenceRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateConfigSequenceRequest) SetSiteId(v int64) *UpdateConfigSequenceRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateConfigSequenceRequest) Validate() error {
	return dara.Validate(s)
}
