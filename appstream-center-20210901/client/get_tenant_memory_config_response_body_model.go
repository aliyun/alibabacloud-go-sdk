// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTenantMemoryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetTenantMemoryConfigResponseBodyData) *GetTenantMemoryConfigResponseBody
	GetData() *GetTenantMemoryConfigResponseBodyData
	SetRequestId(v string) *GetTenantMemoryConfigResponseBody
	GetRequestId() *string
}

type GetTenantMemoryConfigResponseBody struct {
	// The tenant memory configuration that matches the business combination.
	Data *GetTenantMemoryConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique identifier of the request, which can be used for troubleshooting.
	//
	// example:
	//
	// 01A1CEEB-CDFF-18F9-B2BF-9997600D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTenantMemoryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTenantMemoryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetTenantMemoryConfigResponseBody) GetData() *GetTenantMemoryConfigResponseBodyData {
	return s.Data
}

func (s *GetTenantMemoryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTenantMemoryConfigResponseBody) SetData(v *GetTenantMemoryConfigResponseBodyData) *GetTenantMemoryConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetTenantMemoryConfigResponseBody) SetRequestId(v string) *GetTenantMemoryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTenantMemoryConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTenantMemoryConfigResponseBodyData struct {
	// This field is not enabled. The current query result does not provide the disable reason. The example is used to illustrate the field status.
	//
	// example:
	//
	// N/A (not enabled)
	DisabledReason *string `json:"DisabledReason,omitempty" xml:"DisabledReason,omitempty"`
	// This field is not enabled. The current query result does not provide the effective switch of the instance.
	EffectiveEnabled *bool `json:"EffectiveEnabled,omitempty" xml:"EffectiveEnabled,omitempty"`
	// The tenant-level memory switch for the corresponding business combination. A value of `true` indicates that the switch is enabled, and a value of `false` indicates that the switch is disabled. This field may be omitted if no switch is configured for the combination.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This field is not enabled. The current query result does not provide the instance desired switch.
	InstanceDesiredEnabled *bool `json:"InstanceDesiredEnabled,omitempty" xml:"InstanceDesiredEnabled,omitempty"`
	// This field is not enabled. To obtain the tenant-level configuration switch returned by the current API operation, read the Enabled field.
	TenantEnabled *bool `json:"TenantEnabled,omitempty" xml:"TenantEnabled,omitempty"`
}

func (s GetTenantMemoryConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTenantMemoryConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTenantMemoryConfigResponseBodyData) GetDisabledReason() *string {
	return s.DisabledReason
}

func (s *GetTenantMemoryConfigResponseBodyData) GetEffectiveEnabled() *bool {
	return s.EffectiveEnabled
}

func (s *GetTenantMemoryConfigResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetTenantMemoryConfigResponseBodyData) GetInstanceDesiredEnabled() *bool {
	return s.InstanceDesiredEnabled
}

func (s *GetTenantMemoryConfigResponseBodyData) GetTenantEnabled() *bool {
	return s.TenantEnabled
}

func (s *GetTenantMemoryConfigResponseBodyData) SetDisabledReason(v string) *GetTenantMemoryConfigResponseBodyData {
	s.DisabledReason = &v
	return s
}

func (s *GetTenantMemoryConfigResponseBodyData) SetEffectiveEnabled(v bool) *GetTenantMemoryConfigResponseBodyData {
	s.EffectiveEnabled = &v
	return s
}

func (s *GetTenantMemoryConfigResponseBodyData) SetEnabled(v bool) *GetTenantMemoryConfigResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetTenantMemoryConfigResponseBodyData) SetInstanceDesiredEnabled(v bool) *GetTenantMemoryConfigResponseBodyData {
	s.InstanceDesiredEnabled = &v
	return s
}

func (s *GetTenantMemoryConfigResponseBodyData) SetTenantEnabled(v bool) *GetTenantMemoryConfigResponseBodyData {
	s.TenantEnabled = &v
	return s
}

func (s *GetTenantMemoryConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
