// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoDisposeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDisposeConfig(v *GetAutoDisposeConfigResponseBodyAutoDisposeConfig) *GetAutoDisposeConfigResponseBody
	GetAutoDisposeConfig() *GetAutoDisposeConfigResponseBodyAutoDisposeConfig
	SetRequestId(v string) *GetAutoDisposeConfigResponseBody
	GetRequestId() *string
}

type GetAutoDisposeConfigResponseBody struct {
	// The automatic response configuration.
	AutoDisposeConfig *GetAutoDisposeConfigResponseBodyAutoDisposeConfig `json:"AutoDisposeConfig,omitempty" xml:"AutoDisposeConfig,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAutoDisposeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutoDisposeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoDisposeConfigResponseBody) GetAutoDisposeConfig() *GetAutoDisposeConfigResponseBodyAutoDisposeConfig {
	return s.AutoDisposeConfig
}

func (s *GetAutoDisposeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutoDisposeConfigResponseBody) SetAutoDisposeConfig(v *GetAutoDisposeConfigResponseBodyAutoDisposeConfig) *GetAutoDisposeConfigResponseBody {
	s.AutoDisposeConfig = v
	return s
}

func (s *GetAutoDisposeConfigResponseBody) SetRequestId(v string) *GetAutoDisposeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoDisposeConfigResponseBody) Validate() error {
	if s.AutoDisposeConfig != nil {
		if err := s.AutoDisposeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoDisposeConfigResponseBodyAutoDisposeConfig struct {
	// Indicates whether automatic response is enabled. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// enabled
	AutoDecisionStatus *string `json:"AutoDecisionStatus,omitempty" xml:"AutoDecisionStatus,omitempty"`
	// The code for the cloud product.
	//
	// example:
	//
	// alibaba_cloud_sas
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s GetAutoDisposeConfigResponseBodyAutoDisposeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetAutoDisposeConfigResponseBodyAutoDisposeConfig) GoString() string {
	return s.String()
}

func (s *GetAutoDisposeConfigResponseBodyAutoDisposeConfig) GetAutoDecisionStatus() *string {
	return s.AutoDecisionStatus
}

func (s *GetAutoDisposeConfigResponseBodyAutoDisposeConfig) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetAutoDisposeConfigResponseBodyAutoDisposeConfig) SetAutoDecisionStatus(v string) *GetAutoDisposeConfigResponseBodyAutoDisposeConfig {
	s.AutoDecisionStatus = &v
	return s
}

func (s *GetAutoDisposeConfigResponseBodyAutoDisposeConfig) SetProductCode(v string) *GetAutoDisposeConfigResponseBodyAutoDisposeConfig {
	s.ProductCode = &v
	return s
}

func (s *GetAutoDisposeConfigResponseBodyAutoDisposeConfig) Validate() error {
	return dara.Validate(s)
}
